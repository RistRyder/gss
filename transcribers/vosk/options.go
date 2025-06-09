package vosk

import (
	"io/fs"
	"os"

	"github.com/cockroachdb/errors"
)

type TranscriptionOptions struct {
	ffmpegPath string
	modelPath  string
}

func DefaultTranscriptionOptions() (*TranscriptionOptions, error) {
	return NewTranscriptionOptions("ffmpeg", "model")
}

func NewTranscriptionOptions(ffmpegPath, modelPath string) (*TranscriptionOptions, error) {
	if _, statErr := os.Stat(ffmpegPath); errors.Is(statErr, fs.ErrNotExist) {
		return nil, errors.Wrap(statErr, "ffmpeg binary not found")
	}
	if _, statErr := os.Stat(modelPath); errors.Is(statErr, fs.ErrNotExist) {
		return nil, errors.Wrap(statErr, "model path not found")
	}

	return &TranscriptionOptions{
		ffmpegPath: ffmpegPath,
		modelPath:  modelPath,
	}, nil
}
