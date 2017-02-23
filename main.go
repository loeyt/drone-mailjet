package main // import "loe.yt/drone-mailjet"

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/urfave/cli"
)

// Version is changed by the build process.
var Version = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "drone-mailjet"
	app.Usage = "Notification plugin for Drone CI using the Mailjet Send API."
	app.Version = Version
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username",
			Usage:  "Mailjet API username",
			EnvVar: "MJ_APIKEY_PUBLIC,PLUGIN_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "Mailjet API password",
			EnvVar: "MJ_APIKEY_PRIVATE,PLUGIN_PASSWORD",
		},
		cli.StringFlag{
			Name:   "fromname",
			Usage:  "email sender name",
			EnvVar: "PLUGIN_FROMNAME",
		},
		cli.StringFlag{
			Name:   "fromemail",
			Usage:  "email sender address",
			EnvVar: "PLUGIN_FROMEMAIL",
		},
		cli.StringFlag{
			Name:   "recipientname",
			Usage:  "email recipient name",
			EnvVar: "PLUGIN_RECIPIENTNAME",
		},
		cli.StringFlag{
			Name:   "recipientemail",
			Usage:  "email recipient address",
			EnvVar: "PLUGIN_RECIPIENTEMAIL",
		},
		cli.StringFlag{
			Name:   "template",
			Usage:  "email template",
			EnvVar: "PLUGIN_TEMPLATE",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("username") == "" {
		return fmt.Errorf("missing username")
	}
	if c.String("password") == "" {
		return fmt.Errorf("missing password")
	}
	vars := make(map[string]string)
	for _, v := range os.Environ() {
		i := strings.Index(v, "=")
		if i < 0 {
			continue
		}
		if strings.HasPrefix(v, "DRONE_") {
			vars[v[:i]] = v[i+1:]
		}
	}
	mj := mailjet.NewMailjetClient(c.String("username"), c.String("password"))
	email := &mailjet.InfoSendMail{
		FromEmail:          c.String("fromemail"),
		FromName:           c.String("fromname"),
		Subject:            fmt.Sprintf("[%s] build %s", vars["DRONE_REPO_NAME"], vars["DRONE_BUILD_STATUS"]),
		MjTemplateID:       c.String("template"),
		MjTemplateLanguage: "true",
		Recipients: []mailjet.Recipient{
			mailjet.Recipient{
				Name:  c.String("recipientname"),
				Email: c.String("recipientemail"),
			},
		},
		Vars: vars,
	}
	res, err := mj.SendMail(email)
	if err == nil {
		for _, sent := range res.Sent {
			fmt.Printf("email message sent to %s\n", sent.Email)
		}
	}
	return err
}
