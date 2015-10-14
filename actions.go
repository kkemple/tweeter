package main

import (
  "github.com/ChimeraCoder/anaconda"
  "github.com/codegangsta/cli"
  "github.com/kardianos/osext"
  "github.com/peterh/liner"
  "github.com/ttacon/chalk"
  "gopkg.in/yaml.v2"

  "io/ioutil"
  "os"
  "path"
)

var folderName, err = osext.ExecutableFolder()
var INSTALL_PATH = path.Join(folderName, "../src/github.com/kkemple/tweeter")
var CREDS_FILE = path.Join(INSTALL_PATH, ".twittercreds.yml")

type TwitterCreds struct {
  ConsumerKey string `yaml:"consumer_key"`
  ConsumerSecret string `yaml:"consumer_secret"`
  UserToken string `yaml:"user_token"`
  UserSecret string `yaml:"user_secret"`
}

var TweetAction = func(c *cli.Context) {
  var twitterCreds TwitterCreds

  if len(c.Args()) == 0 {
    println(chalk.Magenta.Color("You must provide a tweet"))
    return
  }

  if len(c.Args()) > 1 {
    println(chalk.Magenta.Color("Tweeter only takes one argument"))
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

  println(chalk.Green.Color("Jolly good tweet!"), "\"" + c.Args().First() + "\"")
}

var LoginAction = func(c *cli.Context) {
  var twitterCreds TwitterCreds

  line := liner.NewLiner()
  defer line.Close()

  line.SetCtrlCAborts(true)

  if consumerKey, err := line.Prompt(chalk.Magenta.Color("What is your consumer key? ")); err == nil {
      twitterCreds.ConsumerKey = consumerKey
  } else if err == liner.ErrPromptAborted {
      println(chalk.Magenta.Color("Aborted"))
      return
  } else {
      println("Error reading line: ", err)
  }

  if consumerSecret, err := line.Prompt(chalk.Magenta.Color("What is your consumer secret? ")); err == nil {
      twitterCreds.ConsumerSecret = consumerSecret
  } else if err == liner.ErrPromptAborted {
      println(chalk.Magenta.Color("Aborted"))
      return
  } else {
      println("Error reading line: ", err)
  }

  if userToken, err := line.Prompt(chalk.Magenta.Color("What is your user token? ")); err == nil {
      twitterCreds.UserToken = userToken
  } else if err == liner.ErrPromptAborted {
      println(chalk.Magenta.Color("Aborted"))
      return
  } else {
      println("Error reading line: ", err)
  }

  if userSecret, err := line.Prompt(chalk.Magenta.Color("What is your user secret? ")); err == nil {
      twitterCreds.UserSecret = userSecret
  } else if err == liner.ErrPromptAborted {
      println(chalk.Magenta.Color("Aborted"))
      return
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

  println(chalk.Green.Color("Logged in successfully!"))
}

var LogoutAction = func(c *cli.Context) {
  err := os.Remove(CREDS_FILE)

  if err != nil {
    panic(err)
  }
}
