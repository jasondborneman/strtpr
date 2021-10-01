package Twitter

import (
	"fmt"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/go-twitter/twitter"
)

func Tweet(message string) error {
	config := oauth1.NewConfig(os.Getenv("TWITTER_API_KEY"), os.Getenv("TWITTER_API_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	tweetParams := &twitter.StatusUpdateParams{}

	_, _, err3 := client.Statuses.Update(message, tweetParams)
	if err3 != nil {
		fmt.Println(err3)
		return err3
	}
	return nil
}