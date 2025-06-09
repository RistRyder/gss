# gss - Go Subtitle Suite
This library aims to provide functionality for manipulating subtitles and is currently pre-release.

At the moment there is a transcription API to generate text from an input audio or video file using the following:
[Vosk](https://alphacephei.com/vosk/)
[Whisper on Windows](https://github.com/Purfview/whisper-standalone-win)

## Example - Offline Audio Transcription with [Vosk](https://alphacephei.com/vosk/)
```go
//See https://github.com/alphacep/vosk-api/tree/master/go/example for Vosk installation, ensure all environment variables are set
//Vosk models available at https://alphacephei.com/vosk/models
//Also requires a local copy of ffmpeg

package main

import (
	"fmt"

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
```

## Example - Offline Audio Transcription with [Whisper on Windows](https://github.com/Purfview/whisper-standalone-win)
```go
//See https://github.com/Purfview/whisper-standalone-win for more information, no extra steps are required after downloading

package main

import (
	"fmt"

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
```

## License
`gss` is licensed under the GNU LESSER GENERAL PUBLIC LICENSE Version 3, 
so it free to use for commercial software, as long as you don't modify the library itself. 
LGPL 3.0 allows linking to the library in a way that doesn't require you to open source your own code. 
This means that if you use `gss` in your project, you can keep your own code private, 
as long as you don't modify `gss` itself.