package cli

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var versionTempl = `%s:
  Version: %s
  Commit:  %s
  Time:    %s
`

var supportedDataExts = map[string]struct{}{
	".json": {},
	".toml": {},
	".yaml": {},
	".yml":  {},
}

var app = &cli.App{
	Name:        "tatt",
	Usage:       "template all the things",
	UsageText:   "tatt [options] TEMPLATE_FILE [TEMPLATE_FILE_2 ... TEMPLATE_FILE_N]",
	Description: "Render a Go text or html template with data loaded from a YAML, JSON, or TOML file.",
	Version:     Version(),
	Compiled:    BuildTime(),
	Authors: []*cli.Author{
		{
			Name:  "Michael Henriksen",
			Email: "mchnrksn@gmail.com",
		},
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:      "data",
			Usage:     "load data from `FILE`",
			Aliases:   []string{"d"},
			TakesFile: true,
			Action: func(cCtx *cli.Context, s string) error {
				ext := strings.ToLower(filepath.Ext(s))
				if _, ok := supportedDataExts[ext]; !ok {
					return fmt.Errorf("unsupported data file extension %q", ext)
				}

				if _, err := os.Stat(s); err != nil {
					if errors.Is(err, fs.ErrNotExist) {
						return fmt.Errorf("data file %s does not exist", s)
					}
					return fmt.Errorf("getting info on data file: %w", err)
				}

				return nil
			},
		},
		&cli.BoolFlag{
			Name:        "html",
			Usage:       "use html/template package",
			Value:       false,
			DefaultText: "text/template",
		},
	},
	HideHelp: true,
	Args:     true,
	Action:   renderAction,
	Commands: []*cli.Command{
		{
			Name:    "cheatsheet",
			Aliases: []string{"cheat"},
			Usage:   "view Go template cheat sheet",
			Action:  cheatsheetAction,
		},
	},
}

func App() *cli.App {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Fprintf(cCtx.App.Writer, versionTempl,
			cCtx.App.Name, cCtx.App.Version, BuildRevision(), BuildTime().Format(time.RFC3339),
		)
	}

	return app
}
