# gss - Go Subtitle Suite
This library aims to provide functionality for manipulating subtitles and is currently pre-release.

At the moment there is a transcription API to generate text from an input audio or video file. [Vosk](https://alphacephei.com/vosk/) is the only supported transcription method so far.

## Example - Offline Audio Transcription with [Vosk](https://alphacephei.com/vosk/)
```go
//See https://github.com/alphacep/vosk-api/tree/master/go/example for Vosk installation, ensure all environment variables are set
//Also requires a local copy of ffmpeg

package main

import (
	"fmt"

	"github.com/ristryder/gss/transcribers"
	"github.com/ristryder/gss/transcribers/vosk"
)

func main() {
	transcriber, transcriberErr := vosk.New("/path/to/ffmpeg/binary", "/path/to/Vosk/models")
	if transcriberErr != nil {
		fmt.Println("Failed to create Vosk transcriber: ", transcriberErr)

		return
	}

	audioStreamNumber := uint64(0)

	transcription, transcriptionErr := transcriber.Transcribe(transcribers.NewTranscriptionOptions(audioStreamNumber, "/path/to/audio/or/video/file.mkv"))
	if transcriberErr != nil {
		fmt.Println("Failed to transcribe file: ", transcriptionErr)

		return
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