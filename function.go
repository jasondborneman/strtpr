package strtpr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	app "github.com/jasondborneman/strtpr/StrtprApp"
)

func Strtpr(w http.ResponseWriter, r *http.Request) {
	var d struct {
		StupidAuth string `json:"stupidAuth"`
	}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stupidAuthLocal := os.Getenv("STUPID_AUTH")
	if d.StupidAuth == stupidAuthLocal {
		doTweet := os.Getenv("DO_TWEET") == "true"
		app.Run(doTweet)
		fmt.Fprint(w, "Success")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
	}
}