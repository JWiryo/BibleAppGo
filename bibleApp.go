package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// VerseOfTheDay Struct
type VerseOfTheDay struct {
	Day   int   `json:"day"`
	Verse Verse `json:"verse"`
}

// Verse Struct
type Verse struct {
	Ref  string `json:"human_reference"`
	Text string `json:"text"`
}

func callBibleVerseOfDay() VerseOfTheDay {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://developers.youversionapi.com/1.0/verse_of_the_day/1?version_id=1", nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("x-youversion-developer-token", "Your-API-Key-Here")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resp, err := client.Do(req)
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseVerse VerseOfTheDay
	json.Unmarshal(responseData, &responseVerse)

	return responseVerse
}

func main() {
	verseOfTheDay := callBibleVerseOfDay()

	verse, text := verseOfTheDay.Verse.Ref, verseOfTheDay.Verse.Text
	fmt.Println(verse + " - " + text)
}
