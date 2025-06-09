//go:build windows

package whisperf

import (
	"bufio"
	"io/fs"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	"github.com/cockroachdb/errors"
	"github.com/ristryder/gss/transcribers"
)

type Transcriber struct {
	lineRegex *regexp.Regexp
	options   TranscriptionOptions
}

func (t *Transcriber) createTextLine(textLine string) transcribers.TextLine {
	if t.lineRegex == nil {
		lineRegex, lineRegexErr := regexp.Compile("\\[(?P<fromH>\\d{0,2}):?(?P<fromM>\\d{2}):(?P<fromS>\\d{2})\\.(?P<fromMS>\\d{3}) --> (?P<toH>\\d{0,2}):?(?P<toM>\\d{2}):(?P<toS>\\d{2})\\.(?P<toMS>\\d{3})\\]  (?P<text>.+)")
		if lineRegexErr != nil {
			return transcribers.TextLine{}
		}

		t.lineRegex = lineRegex
	}

	match := t.lineRegex.FindStringSubmatch(textLine)

	matchResults := map[string]string{}
	for i, name := range match {
		matchResults[t.lineRegex.SubexpNames()[i]] = name
	}

	startTime := getIntFromMatch("fromS", matchResults)
	startTime += getIntFromMatch("fromM", matchResults) * 60
	startTime += getIntFromMatch("fromH", matchResults) * 60 * 60
	startTime += getIntFromMatch("fromMS", matchResults) / 1000

	endTime := getIntFromMatch("toS", matchResults)
	endTime += getIntFromMatch("toM", matchResults) * 60
	endTime += getIntFromMatch("toH", matchResults) * 60 * 60
	endTime += getIntFromMatch("toMS", matchResults) / 1000

	return transcribers.TextLine{
		EndTime:   endTime,
		StartTime: startTime,
		Text:      matchResults["text"],
	}
}

func getIntFromMatch(groupName string, matchResults map[string]string) float64 {
	num, numErr := strconv.Atoi(matchResults[groupName])
	if numErr != nil {
		return 0
	}

	return float64(num)
}

func NewTranscriber(options *TranscriptionOptions) (*Transcriber, error) {
	if options == nil {
		transcriptionOpt, transcriptionOptErr := DefaultTranscriptionOptions()
		if transcriptionOptErr != nil {
			return nil, transcriptionOptErr
		}

		return &Transcriber{options: *transcriptionOpt}, nil
	}

	return &Transcriber{options: *options}, nil
}

func (t *Transcriber) Transcribe(options transcribers.TranscriptionOptions) (*transcribers.TranscriptionResults, error) {
	if _, statErr := os.Stat(options.FullFileName); errors.Is(statErr, fs.ErrNotExist) {
		return nil, errors.Wrap(statErr, "input file not found")
	}

	progressCallback := options.ProgressCallback
	if progressCallback == nil {
		progressCallback = func(t transcribers.TextLine) {}
	}

	whisperFasterCmd := exec.Command(t.options.whisperFasterPath, options.FullFileName, "--beep_off", "-o", "NUL")
	stdoutPipe, pipeErr := whisperFasterCmd.StdoutPipe()
	if pipeErr != nil {
		return nil, errors.Wrap(pipeErr, "failed to initialize pipe for faster-whisper stream")
	}

	defer stdoutPipe.Close()

	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stdoutScanner.Split(bufio.ScanLines)

	startErr := whisperFasterCmd.Start()
	if startErr != nil {
		return nil, errors.Wrap(startErr, "failed to start faster-whisper")
	}

	results := []transcribers.TextLine{}

	for stdoutScanner.Scan() {
		textLine := t.createTextLine(stdoutScanner.Text())

		if len(textLine.Text) > 0 {
			results = append(results, textLine)
		}
	}

	defer whisperFasterCmd.Wait()

	return &transcribers.TranscriptionResults{Lines: results}, nil
}
