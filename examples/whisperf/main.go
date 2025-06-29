//See https://github.com/Purfview/whisper-standalone-win for more information, no extra steps are required after downloading

package main

import (
	"fmt"
	"log"

	"github.com/ristryder/gss/transcribers"
	"github.com/ristryder/gss/transcribers/whisperf"
)

func main() {
	transcriptionOpt, transcriptionOptErr := whisperf.NewTranscriptionOptions("medium", "", "C:\\Path\\To\\Faster-Whisper-XXL_r245.4_windows\\Faster-Whisper-XXL\\faster-whisper-xxl.exe")
	if transcriptionOptErr != nil {
		log.Fatalf("Failed to create transcription options: %s\n", transcriptionOptErr)
	}

	transcriber, transcriberErr := whisperf.NewTranscriber(transcriptionOpt)
	if transcriberErr != nil {
		log.Fatalf("Failed to create Whisperf transcriber: %s\n", transcriberErr)
	}

	audioStreamNumber := uint64(0)

	transcription, transcriptionErr := transcriber.Transcribe(transcribers.NewTranscriptionOptions(audioStreamNumber, "C:\\Path\\To\\Audio\\Or\\Video\\File.mkv"))
	if transcriberErr != nil {
		log.Fatalf("Failed to transcribe file: %s\n", transcriptionErr)
	}

	for i, line := range transcription.Lines {
		fmt.Printf("[%d][%f - %f] --> %s\n", i, line.StartTime, line.EndTime, line.Text)
	}
}
