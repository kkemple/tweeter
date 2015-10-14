package main

import (
  "github.com/codegangsta/cli"
  "os"
)

func main() {
  app := cli.NewApp()
  app.Name = "Tweeter"
  app.Usage = "Tweet from the command line!"
  app.Version = "0.2.0"
  app.Commands = Commands

  app.Run(os.Args)
}
