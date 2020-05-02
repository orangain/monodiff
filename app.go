package monodiff

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// App is a CLI app of monodiff using urfave/cli/v2.
var App = &cli.App{
	Name:  "monodiff",
	Usage: "Simple tool to detect modified part of monorepo",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "prefix",
			Value: "",
			Usage: "Prefix of output line",
		},
		&cli.StringFlag{
			Name:  "suffix",
			Value: "",
			Usage: "Suffix of output line",
		},
		&cli.StringFlag{
			Name:  "separator",
			Value: "",
			Usage: "Path separator of output line",
		},
	},
	Action: func(c *cli.Context) error {
		prefix := c.String("prefix")
		suffix := c.String("suffix")
		separator := c.String("separator")

		spec, err := loadSpec()
		if err != nil {
			return err
		}

		changedFiles := []string{}
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			text := stdin.Text()
			if text != "" {
				changedFiles = append(changedFiles, text)
			}
		}

		changedProjects, err := detectChanges(spec, changedFiles)
		for _, changedProject := range changedProjects {
			name := changedProject.Name
			if separator != "" {
				name = strings.ReplaceAll(name, "/", separator)
			}
			fmt.Println(prefix + name + suffix)
		}

		return nil
	},
}
