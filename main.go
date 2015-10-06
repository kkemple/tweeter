package main

import (
  "github.com/ChimeraCoder/anaconda"
  "github.com/codegangsta/cli"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "os"
)

type TwitterCreds struct {
  ConsumerKey string `yaml:"consumer_key"`
  ConsumerSecret string `yaml:"consumer_secret"`
  UserToken string `yaml:"user_token"`
  UserSecret string `yaml:"user_secret"`
}

func main() {
  app := cli.NewApp()
  app.Name = "tweeter"
  app.Usage = "Tweet from the command line!"
  app.Action = func(c *cli.Context) {
    var twitterCreds TwitterCreds

    if len(c.Args()) == 0 {
      println("You must provide a tweet")
      return
    }

    if len(c.Args()) > 1 {
      println("tweeter only takes one argument")
      return
    }

    source, err := ioutil.ReadFile(".twittercreds.yml")
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(source, &twitterCreds)
    if err != nil {
        panic(err)
    }

    anaconda.SetConsumerKey(twitterCreds.ConsumerKey)
    anaconda.SetConsumerSecret(twitterCreds.ConsumerSecret)
    api := anaconda.NewTwitterApi(
      twitterCreds.UserToken,
      twitterCreds.UserSecret,
    )

    api.PostTweet(c.Args()[0], nil)

    println("Jolly good tweet!", c.Args()[0])
  }

  app.Run(os.Args)
}
