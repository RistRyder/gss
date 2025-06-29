package fix

import (
	"github.com/ristryder/gss/common"
)

func Overlap(subtitles []common.TextLine) ([]common.TextLine, error) {
	fixedSubtitles := []common.TextLine{}

	for i := 0; i < len(subtitles); i++ {
		if i < len(subtitles)-1 && subtitles[i].EndTime > subtitles[i+1].StartTime {
			joinedLine := common.TextLine{StartTime: subtitles[i].StartTime, EndTime: subtitles[i+1].EndTime, Text: subtitles[i].Text + "\n" + subtitles[i+1].Text}

			fixedSubtitles = append(fixedSubtitles, joinedLine)

			i++
		} else {
			fixedSubtitles = append(fixedSubtitles, subtitles[i])
		}
	}

	return fixedSubtitles, nil
}
