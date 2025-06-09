//go:build windows

package whisperf

import (
	"io/fs"
	"os"

	"github.com/cockroachdb/errors"
)

type TranscriptionOptions struct {
	ModelDirectory string

	model             string
	whisperFasterPath string
}

func DefaultTranscriptionOptions() (*TranscriptionOptions, error) {
	return NewTranscriptionOptions("", "", "faster-whisper-xxl.exe")
}

func NewTranscriptionOptions(model string, modelDirectory string, whisperFasterPath string) (*TranscriptionOptions, error) {
	if _, statErr := os.Stat(whisperFasterPath); errors.Is(statErr, fs.ErrNotExist) {
		return nil, errors.Wrap(statErr, "faster-whisper binary not found")
	}

	return &TranscriptionOptions{
		model:             model,
		ModelDirectory:    modelDirectory,
		whisperFasterPath: whisperFasterPath,
	}, nil
}
