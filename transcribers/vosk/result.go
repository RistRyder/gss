package vosk

import (
	"encoding/json"

	"github.com/ristryder/gss/common"
)

type result struct {
	Text  string       `json:"text"`
	Words []resultWord `json:"result"`
}

func fromString(resultString string) result {
	var res result

	jsonErr := json.Unmarshal([]byte(resultString), &res)
	if jsonErr != nil {
		return result{}
	}

	return res
}

func toTextLine(res result) common.TextLine {
	return common.TextLine{
		EndTime:   res.Words[len(res.Words)-1].EndTime,
		StartTime: res.Words[0].StartTime,
		Text:      res.Text,
	}
}
