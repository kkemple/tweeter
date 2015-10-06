#tweeter

The tweeter app is powered by `codegangsta/cli` for the cli app and `ChimeraCoder/anaconda` for the twitter api.

## How it works

The app takes a single command line argument that is the tweet to be tweeted
Internally tweeter reads the Twitter API creds from a `.twittercreds.yml` file (git ignored) and loads up the Twitter API and posts the tweet.

## How to run it

### Locally
Run `go build`, then run `./tweeter "<your tweet>"`

### Globally
Run `go install`, then from anywhere run `tweeter "<your tweet>"`
