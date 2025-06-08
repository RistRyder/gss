package transcribers

type TranscriptionOptions struct {
	AudioTrackNumber uint64
	FullFileName     string
}

func NewTranscriptionOptions(audioTrackNumber uint64, fullFileName string) TranscriptionOptions {
	return TranscriptionOptions{
		AudioTrackNumber: audioTrackNumber,
		FullFileName:     fullFileName,
	}
}
