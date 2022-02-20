package main

import (
	"strings"

	"github.com/bxcodec/faker/v3"
	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func run() {
	fakeType := wf.Args()[0]
	addItem(fakeType)
	wf.SendFeedback()
}

func addItem(fakeType string) {
	var fakeData string
	switch strings.ToLower(fakeType) {
	case "email":
		fakeData = faker.Email()
	case "name":
		fakeData = faker.Name()
	case "paragraph":
		fakeData = faker.Paragraph()
	case "sentence":
		fakeData = faker.Sentence()
	}
	if fakeData != "" {
		wf.NewItem(fakeData).
			Copytext(fakeData).
			Subtitle(fakeType).
			Valid(true)
	} else {
		wf.Warn("unknown type", fakeType)
	}
}

func main() {
	wf = aw.New()
	wf.Run(run)
}

