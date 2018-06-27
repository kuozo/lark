package main

import (
	"os"
	"log"
	
	cli "gopkg.in/urfave/cli.v1"
	"github.com/robfig/cron"

	"github.com/klnchu/lark/pkg"
)
var token string
var schedule string

func start(){
	var c = cron.New()
	log.Println(schedule)
	c.AddFunc(schedule, func(){
		log.Println("Tweet...")
		err := pkg.Tweet(token)
		if err != nil{
			log.Fatal(err)
		}
	})
	c.Run()
}

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
		cli.StringFlag{
			Name: "schedule",
			Value: "0 25 * * * *",
			Usage: "Crontab Schedule Formart",
			Destination: &schedule,
			EnvVar: "CRONJOB_SCHEDULE",
		},
	}
	app.Action = func(c *cli.Context) error{
		start()
		return nil 
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("-EOF-")
}