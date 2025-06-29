package util

import (
	"sort"

	"github.com/ristryder/gss/common"
)

func Join(subtitles ...[]common.TextLine) ([]common.TextLine, error) {
	finalSubtitles := []common.TextLine{}

	for _, subs := range subtitles {
		finalSubtitles = append(finalSubtitles, subs...)
	}

	sort.Slice(finalSubtitles, func(i, j int) bool {
		return finalSubtitles[i].StartTime < finalSubtitles[j].StartTime
	})

	return finalSubtitles, nil
}
