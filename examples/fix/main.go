package main

import (
	"fmt"
	"log"

	"github.com/ristryder/gss/common"
	"github.com/ristryder/gss/fix"
	"github.com/ristryder/gss/util"
)

func main() {
	defaultSubtitles := []common.TextLine{
		{StartTime: 4347790, EndTime: 4349220, Text: "Hi, Hi"},
		{StartTime: 4410175, EndTime: 4411271, Text: "Have you"},
	}
	forcedSubtitles := []common.TextLine{
		{StartTime: 4347085, EndTime: 4348170, Text: "G'day, Bobert."},
		{StartTime: 4349546, EndTime: 4351130, Text: "G'day, Bobert!"},
	}

	finalSubtitles, finalSubtitlesErr := util.Join(defaultSubtitles, forcedSubtitles)
	if finalSubtitlesErr != nil {
		log.Fatal(finalSubtitlesErr)
	}

	fixedSubtitles, fixedSubtitlesErr := fix.Overlap(finalSubtitles)
	if fixedSubtitlesErr != nil {
		log.Fatal(fixedSubtitlesErr)
	}

	for i, line := range fixedSubtitles {
		fmt.Printf("[%d][%f - %f] --> %s\n", i, line.StartTime, line.EndTime, line.Text)
	}
}
