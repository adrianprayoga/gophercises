package story

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Story map[string]StoryStructure

type StoryStructure struct {
	Title   string     `json:"title"`
	Story   []string   `json:"story"`
	Options []StoryArc `json:"options"`
}

type StoryArc struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func GetJson(storyPtr *string) Story {

	jsonByte, err := ioutil.ReadFile(*storyPtr)
	if err != nil {
		log.Fatal("Unable to read input file " + *storyPtr)
		log.Fatal(err)
		os.Exit(1)
	}

	var story Story
	json.Unmarshal(jsonByte, &story)

	return story
}
