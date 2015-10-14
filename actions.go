package main

import (
  "github.com/ChimeraCoder/anaconda"
  "github.com/codegangsta/cli"
  "github.com/peterh/liner"
  "gopkg.in/yaml.v2"

  "io/ioutil"
  "os"
)

const CREDS_FILE = ".twittercreds.yml"

type TwitterCreds struct {
  ConsumerKey string `yaml:"consumer_key"`
  ConsumerSecret string `yaml:"consumer_secret"`
  UserToken string `yaml:"user_token"`
  UserSecret string `yaml:"user_secret"`
}

var TweetAction = func(c *cli.Context) {
  var twitterCreds TwitterCreds

  if len(c.Args()) == 0 {
    println("You must provide a tweet")
    return
  }

  if len(c.Args()) > 1 {
    println("tweeter only takes one argument")
    return
  }

  source, err := ioutil.ReadFile(CREDS_FILE)
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

  api.PostTweet(c.Args().First(), nil)

  println("Jolly good tweet!", c.Args().First())
}

var LoginAction = func(c *cli.Context) {
  var twitterCreds TwitterCreds

  line := liner.NewLiner()
  defer line.Close()

  line.SetCtrlCAborts(true)

  if consumerKey, err := line.Prompt("What is consumer key? "); err == nil {
      twitterCreds.ConsumerKey = consumerKey
  } else if err == liner.ErrPromptAborted {
      println("Aborted")
  } else {
      println("Error reading line: ", err)
  }

  if consumerSecret, err := line.Prompt("What is your consumer secret? "); err == nil {
      twitterCreds.ConsumerSecret = consumerSecret
  } else if err == liner.ErrPromptAborted {
      println("Aborted")
  } else {
      println("Error reading line: ", err)
  }

  if userToken, err := line.Prompt("What is your user token? "); err == nil {
      twitterCreds.UserToken = userToken
  } else if err == liner.ErrPromptAborted {
      println("Aborted")
  } else {
      println("Error reading line: ", err)
  }

  if userSecret, err := line.Prompt("What is your user secret? "); err == nil {
      twitterCreds.UserSecret = userSecret
  } else if err == liner.ErrPromptAborted {
      println("Aborted")
  } else {
      println("Error reading line: ", err)
  }

  y, err := yaml.Marshal(&twitterCreds)
  if err != nil {
    panic(err)
  }

  err = ioutil.WriteFile(CREDS_FILE, y, 0644)
  if err != nil {
      panic(err)
  }

  println("Logged in successfully!")
}

var LogoutAction = func(c *cli.Context) {
  err := os.Remove(CREDS_FILE)

  if err != nil {
    panic(err)
  }
}
