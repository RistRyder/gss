package transcribers

import "github.com/ristryder/gss/common"

type TranscriptionOptions struct {
	AudioTrackNumber uint64
	FullFileName     string
	ProgressCallback func(common.TextLine)
}

func NewTranscriptionOptions(audioTrackNumber uint64, fullFileName string) TranscriptionOptions {
	return TranscriptionOptions{
		AudioTrackNumber: audioTrackNumber,
		FullFileName:     fullFileName,
	}
}
