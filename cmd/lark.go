package main

import (
	"os"
	"log"
	
	cli "gopkg.in/urfave/cli.v1"

	"github.com/klnchu/lark/pkg"
)
var token string 

func main(){
	log.Println("Starting")

	app := cli.NewApp()
	app.Name = "Lark"
	app.Usage = "Open Source Letter."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "token",
			Value: "xxxxx",
			Usage: "DingDing Robot token",
			Destination: &token,
			EnvVar: "DING_ROBOT_TOKEN",
		},
	}
	app.Action = func(c *cli.Context) error{
		err := pkg.Tweet(token)
		if err != nil{
			log.Fatal(err)
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	
	
}