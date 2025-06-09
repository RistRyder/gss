package vosk

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strconv"

	vosk "github.com/alphacep/vosk-api/go"
	"github.com/cockroachdb/errors"
	"github.com/ristryder/gss/io"
	"github.com/ristryder/gss/transcribers"
)

const sampleRate = 16000

type Transcriber struct {
	options TranscriptionOptions
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

	model, modelErr := vosk.NewModel(t.options.modelPath)
	if modelErr != nil {
		return nil, errors.Wrap(modelErr, "failed to create Vosk model")
	}

	recognizer, recognizerErr := vosk.NewRecognizer(model, sampleRate)
	if recognizerErr != nil {
		defer model.Free()

		return nil, errors.Wrap(recognizerErr, "failed to create Vosk recognizer")
	}

	defer recognizer.Free()
	defer model.Free()

	recognizer.SetMaxAlternatives(0)
	recognizer.SetWords(1)

	progressCallback := options.ProgressCallback
	if progressCallback == nil {
		progressCallback = func(t transcribers.TextLine) {}
	}

	//-ar              | sampling rate
	//-af volume=1.75  | boost the volume
	//-ac              | # of audio channels
	//-map 0:a:N       | select Nth audio stream
	//-vn              | no video output
	ffmpegCmd := exec.Command(t.options.ffmpegPath, "-nostdin", "-vn", "-loglevel", "quiet", "-i", options.FullFileName, "-ar", strconv.Itoa(sampleRate), "-af", "volume=1.75", "-ac", "1", "-map", fmt.Sprintf("0:a:%d", options.AudioTrackNumber), "-f", "s16le", "-")
	stdoutPipe, pipeErr := ffmpegCmd.StdoutPipe()
	if pipeErr != nil {
		return nil, errors.Wrap(pipeErr, "failed to initialize pipe for ffmpeg stream")
	}

	defer stdoutPipe.Close()

	stdoutScanner := bufio.NewScanner(stdoutPipe)
	stdoutScanner.Split(io.GetSplitter(4000))

	startErr := ffmpegCmd.Start()
	if startErr != nil {
		return nil, errors.Wrap(startErr, "failed to start ffmpeg")
	}

	results := []transcribers.TextLine{}

	for stdoutScanner.Scan() {
		if recognizer.AcceptWaveform(stdoutScanner.Bytes()) != 0 {
			currentResult := fromString(recognizer.Result())

			if len(currentResult.Words) > 0 {
				textLine := toTextLine(currentResult)

				results = append(results, textLine)

				progressCallback(textLine)
			}
		}
	}

	finalResult := fromString(recognizer.FinalResult())

	if len(finalResult.Words) > 0 {
		results = append(results, toTextLine(finalResult))
	}

	defer ffmpegCmd.Wait()

	return &transcribers.TranscriptionResults{Lines: results}, nil
}
