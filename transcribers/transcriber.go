package transcribers

type Transcriber interface {
	Transcribe(options TranscriptionOptions) (*TranscriptionResults, error)
}
