# strtpr
[![Go](https://github.com/jasondborneman/strtpr/actions/workflows/go.yml/badge.svg)](https://github.com/jasondborneman/strtpr/actions/workflows/go.yml)

Pulls random words out of an API, then uses the Merriam-Webster API to chck for part of speech and build stupid startup ideas and names.

## APIs Used
* Twitter: =https://github.com/dghubble/go-twitter
* Random Word API: https://random-word-api.herokuapp.com/home
* Merriam-Webster Dictionary API: https://dictionaryapi.com/

## Environment Variables Needed
* DICT_APIKEY : Merriam-Webster Dictionary API Key
* TWITTER_API_KEY : Twitter Key
* TWITTER_API_SECRET : Twitter Secret
* TWITTER_ACCESS_TOKEN : Twitter Access Token
* TWITTER_ACCESS_TOKEN_SECRET : Twitter Access Secret
* STUPID_AUTH : Make up your own GUID or any string really.

### Optional Environment Variables
* DO_TWEET : [true|false] to turn on and off tweeting. Defaults to false

## Auth
Set a local environment variable wherever your function is running called STUPID_AUTH with a value of some string key (such as a GUID). When calling the function endpoint, the body should be:

```
{
  "stupidAuth":"<your key here>"
}
```

As the name suggests, auth is really stupid. It just checks that the body key matches the Environment Variable.

## Tech Stack

* Function is written in Go.
* Runs as a GCP Cloud Function
* Triggered by a GCP Cloud Scheduler Job
* Stored and CI/CD in Github 
* Also stored in a linked GCP Cloud Source Repository

## Deployment

It's currently set up using a GitHub CI/CD Action. This will be fleshed out further to include builds of all PRs to `main` in addition to the Build/Deploy of all commits to `main`.

It's also set up to deploy from a local command line. Set up a linked GCP Cloud Source Repository for the GitHub repo (or just store all your code there and skip GitHub altogether). The following shell command will deploy from the Cloud Source Repository to the Cloud Function:

```
gcloud functions deploy <project_name> \
  --source https://source.developers.google.com/projects/<project id>repos/<cloud source repo name>/moveable-aliases/<branch>/paths/ \
  --runtime go116 \
  --trigger-http \
  --allow-unauthenticated
```

it is recommended, if you're going to use the GitHub Action to deploy, to set up the Function first time on the command line, as GitHub Action for deploying to GCP Cloud Functions doesn't allow unauthenticated functions to be deployed from what I can tell, but it won't mess with the setting if it's already there.

## Running locally
Set local ENVs first...

```
$ go build cmd/strtpr.go
$ ./strtpr
```

## TODO:
* Unit tests - still need to figure out mocking in Go to do this. This will probably mean some restructuring of the code to make it more testable. I was writing/learning as I went writing this.
