package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"go.lafronz.net/web_scrap/internal"
)

var app *cli.App

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name: "version", Aliases: []string{"V", "v"},
		Usage: "print only the version",
	}

	app = &cli.App{
		Name:    "web",
		Version: "v0.1.0",
		Usage:   "a web scrap tool",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Tyler La Fronz",
				Email: "tyler@lafronz.com",
			},
		},
		Action: func(c *cli.Context) error {
			cli.ShowAppHelpAndExit(c, 0)
			return nil
		},
		Commands: []*cli.Command{
			&cli.Command{
				Name:        "scrap",
				Aliases:     []string{"s"},
				Usage:       "return the output of a web crawl",
				Description: "Use this command to retrieve the contents of a website",
				ArgsUsage:   "[url]",
				Action: func(c *cli.Context) error {
					if c.Args().Len() <= 0 {
						cli.ShowCommandHelp(c, "scrap")
						return nil
					}
					internal.GetURL(c.Args().First())
					return nil
				},
			},
			&cli.Command{
				Name:        "markdown",
				Aliases:     []string{"m"},
				Usage:       "write markdown file of the manual",
				Description: "This command will write a file call Manual.md to the local directory",
				Action: func(c *cli.Context) error {
					markdownCreate()
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func markdownCreate() error {
	res, err := app.ToMarkdown()
	if err != nil {
		return err
	}
	err1 := ioutil.WriteFile("Manual.md", []byte(res), 0644)
	if err1 != nil {
		return err1
	}

	return nil
}
