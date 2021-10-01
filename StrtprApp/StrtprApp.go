package StrtprApp


import (
	"fmt"
	"strings"
	rw "github.com/jasondborneman/strtpr/RandomWordApi"
	tw "github.com/jasondborneman/strtpr/Twitter"
)

func Run(doTweet bool) {
	fmt.Println("Strtpr!-----------------------------")
	noun,_ := rw.GetWordByPartOfSpeech("noun")
	noun = strings.Split(noun,":")[0]
	startupName := convertToStartupyWord(noun)
	verb,_ := rw.GetWordByPartOfSpeech("verb")
	verb = strings.Split(verb,":")[0]
	item,_ := rw.GetWordByPartOfSpeech("noun")
	item = strings.Split(item,":")[0]
	message := fmt.Sprintf("%s: %s your %s in the cloud!\nNow go MONETIZE THEM STRATEGIES with %s!", startupName, strings.Title(verb), item, startupName)
	fmt.Println(fmt.Sprintf(message))
	if doTweet == true {
		tw.Tweet(message)
		return
	}
	fmt.Println("Did not tweet this time.")
}

func convertToStartupyWord(word string) (string){
	stripped := strings.Title(strings.Replace(stripSomeLetters(word), " ", "", -1))
	return fmt.Sprintf("%sr", stripped)
}

func stripSomeLetters(word string) (string) {
	newWord := strings.TrimRight(word, "a")
	newWord = strings.TrimRight(newWord, "e")
	newWord = strings.TrimRight(newWord, "i")
	newWord = strings.TrimRight(newWord, "o")
	newWord = strings.TrimRight(newWord, "u")
	newWord = strings.TrimRight(newWord, "er")
	newWord = strings.TrimRight(newWord, "ier")
	newWord = strings.TrimRight(newWord, "ing")
	if newWord == word {
		return newWord
	}
	return stripSomeLetters(newWord)
}