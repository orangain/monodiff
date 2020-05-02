package monodiff

import (
	"bufio"
	"os"

	"github.com/urfave/cli/v2"
)

// App is a CLI app of monodiff using urfave/cli/v2.
var App = &cli.App{
	Name:  "monodiff",
	Usage: "Simple tool to detect modified part of monorepo",
	Action: func(c *cli.Context) error {
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
			println(changedProject.Name)
		}

		return nil
	},
}
