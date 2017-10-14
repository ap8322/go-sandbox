package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App{
	app := cli.NewApp()
	app.Name = "sample_client"
	app.Usage = "cli test"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "hello",
			Aliases: []string{"h"},
			Usage:   "hello world",
			Action:  helloAction,
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun,d",
			Usage: "dryrun",
		},
	}


	app.Before = func(c *cli.Context) error {
		fmt.Println("開始")
		return nil
	}

	app.After = func(c *cli.Context) error {
		fmt.Println("終了")
		return nil
	}

	return app
}

func helloAction(c *cli.Context) {

	var isDry = c.GlobalBool("dryrun")
	if isDry {
		fmt.Println("this is dry-run")
	}

	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	}

	fmt.Printf("hello world %s\n", paramFirst)

}
