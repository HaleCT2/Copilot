package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cli-example"
	app.Usage = "Example CLI application"

	// flags
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.google.com",
		},
	}

	// commands
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Looks up the IP of a host",
			Flags: myFlags,
			// execute `ip` command
			Action: func(c *cli.Context) error {
				// a simple lookup function
				ip, err := net.LookupIP(c.String("url"))
				if err != nil {
					return err
				}
				// we log the results to our console
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
	}

	// start application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
