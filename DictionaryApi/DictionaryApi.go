package DictionaryApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	dw "github.com/jasondborneman/strtpr/DictionaryWord"
)

var dictApiUrl = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/"
var dictApiKey = os.Getenv("DICT_APIKEY")

func IsPartOfSpeech(word string, partOfSpeech string) (bool, string, error) {
	url := fmt.Sprintf("%s%s?key=%s", dictApiUrl, word, dictApiKey)
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	resp, getErr := netClient.Get(url)

	if getErr != nil {
		log.Fatal(fmt.Sprintf("Error getting calling Dictionary API: %s", getErr))
		return false, word, getErr

	}
	if resp.StatusCode != 200 {
		message := fmt.Sprintf("Non-200 Status Code Returned: %d [%s]", resp.StatusCode, url)
		log.Fatal(message)
		return false, word, errors.New(message)
	}

	wordInfo := &dw.WordInfo{}
	body, _ := ioutil.ReadAll(resp.Body)
	if strings.HasPrefix(string(body), `["`) {
		message := fmt.Sprintf("Word not found in dictionary: %s", word)
		return false, word, errors.New(message)
	}

	decodeErr := json.Unmarshal([]byte(body), &wordInfo)
	if decodeErr != nil {
		log.Fatal(fmt.Sprintf("Error decoding Word Info response: %s", decodeErr))
		return false, word, decodeErr
	}
	foundPartOfSpeech := (*wordInfo)[0].Fl
	finalWord := (*wordInfo)[0].Meta.ID
	defer resp.Body.Close()
	return foundPartOfSpeech == partOfSpeech, finalWord, nil
}
