//See https://github.com/alphacep/vosk-api/tree/master/go/example for Vosk installation, ensure all environment variables are set
//Vosk models available at https://alphacephei.com/vosk/models
//Also requires a local copy of ffmpeg

package main

import (
	"fmt"
	"log"

	"github.com/ristryder/gss/transcribers"
	"github.com/ristryder/gss/transcribers/vosk"
)

func main() {
	transcriptionOpt, transcriptionOptErr := vosk.NewTranscriptionOptions("/path/to/ffmpeg/binary", "/path/to/Vosk/models")
	if transcriptionOptErr != nil {
		log.Fatalf("Failed to create transcription options: %s\n", transcriptionOptErr)
	}

	transcriber, transcriberErr := vosk.NewTranscriber(transcriptionOpt)
	if transcriberErr != nil {
		log.Fatalf("Failed to create Vosk transcriber: %s\n", transcriberErr)
	}

	audioStreamNumber := uint64(0)

	transcription, transcriptionErr := transcriber.Transcribe(transcribers.NewTranscriptionOptions(audioStreamNumber, "/path/to/audio/or/video/file.mkv"))
	if transcriberErr != nil {
		log.Fatalf("Failed to transcribe file: %s\n", transcriptionErr)
	}

	for i, line := range transcription.Lines {
		fmt.Printf("[%d][%f - %f] --> %s\n", i, line.StartTime, line.EndTime, line.Text)
	}
}
