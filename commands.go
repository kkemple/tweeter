package main

import "github.com/codegangsta/cli"

var Commands = []cli.Command{
  {
    Name: "tweet",
    Aliases: []string{"t"},
    Usage: "Tweet up on your Twitter!",
    Action: TweetAction,
  },

  {
    Name: "login",
    Aliases: []string{"li"},
    Usage: "Auth with Twitter API",
    Action: LoginAction,
  },

  {
    Name: "logout",
    Aliases: []string{"lo"},
    Usage: "Unauth with Twitter API",
    Action: LogoutAction,
  },
}
