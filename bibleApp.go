package main

import (
	constants "bibleAppGo/constants"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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

	t := time.Now().YearDay
	todayAPIURL := fmt.Sprintf(constants.APIURL, t())

	req, err := http.NewRequest("GET", todayAPIURL, nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("x-youversion-developer-token", constants.BibleAPIKey) // You can replace this with your API Key

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
