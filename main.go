package main // import "loe.yt/drone-mailjet"

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Version is changed by the build process.
var Version = "dev"

func main() {
	app := cli.NewApp()
	app.Name = "drone-mailjet"
	app.Version = Version
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, "DRONE_") {
			fmt.Println(v)
		}
	}
	return nil
}
