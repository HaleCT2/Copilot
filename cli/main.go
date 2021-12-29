package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cli-example",
		Usage: "Example CLI application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "www.google.com",
				Usage: "FQDN of the host to check",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "ip",
				Usage: "Looks up the IP of a host",
				Action: func(c *cli.Context) error {
					fmt.Println(c.String("host") + ":")
					ip, err := net.LookupIP(c.String("host"))
					if err != nil {
						return err
					}
					for i := 0; i < len(ip); i++ {
						fmt.Println(ip[i])
					}
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
