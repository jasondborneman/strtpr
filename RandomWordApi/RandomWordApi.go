package RandomWordApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	da "github.com/jasondborneman/strtpr/DictionaryApi"
)

var wordsApiUrl = "https://random-word-api.herokuapp.com/word?number=100&swear=0"

func GetWordByPartOfSpeech(partOfSpeech string) (string, error) {
	url := wordsApiUrl
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	resp, getErr := netClient.Get(url)

	if getErr != nil {
		log.Fatal(fmt.Sprintf("Error getting calling Random Word API: %s", getErr))
		return "", getErr

	}
	if resp.StatusCode != 200 {
		message := fmt.Sprintf("Non-200 Status Code Returned: %d [%s]", resp.StatusCode, url)
		log.Fatal(message)
		return "", errors.New(message)
	}
	chosenWord := ""
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error converting Random Word API body to string. %s", err))
		return "", err
	}
	words := respToArr(string(body))
	for _, word := range words {
		wordIsType, checkedWord, _ := da.IsPartOfSpeech(word, partOfSpeech)
		if wordIsType == true {
			chosenWord = checkedWord
			break
		}
	}

	if chosenWord == "" {
		fmt.Printf("No word of type found, trying again")
		return GetWordByPartOfSpeech(partOfSpeech)
	}
	defer resp.Body.Close()
	return chosenWord, nil
}

func respToArr(resp string) ([]string) {
	var arr []string
	_ = json.Unmarshal([]byte(resp), &arr)
	return arr
}