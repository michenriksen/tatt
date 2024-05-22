package cli

import (
	"encoding/json"
	"fmt"
	htmlTemplate "html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	textTemplate "text/template"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/sprig/v3"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type templater interface {
	Execute(io.Writer, any) error
}

func renderAction(cCtx *cli.Context) error {
	if cCtx.NArg() == 0 {
		fmt.Fprintln(cCtx.App.ErrWriter, "Missing template file argument")
		cli.ShowAppHelpAndExit(cCtx, 1)
	}

	tmpl, err := loadTemplate(cCtx)
	if err != nil {
		return cli.Exit(fmt.Sprintf("loading template: %v", err), 1)
	}

	data, err := loadData(cCtx)
	if err != nil {
		return cli.Exit(fmt.Sprintf("loading data: %v", err), 1)
	}

	if err := tmpl.Execute(cCtx.App.Writer, data); err != nil {
		return cli.Exit(fmt.Sprintf("rendering template: %v", err), 1)
	}

	return nil
}

func loadTemplate(cCtx *cli.Context) (templater, error) {
	name := filepath.Base(cCtx.Args().Get(0))

	if cCtx.Bool("html") {
		tmpl, err := htmlTemplate.New(name).Funcs(sprig.HtmlFuncMap()).ParseFiles(cCtx.Args().Slice()...)
		if err != nil {
			return nil, fmt.Errorf("parsing html template: %w", err)
		}
		return tmpl, nil
	}

	tmpl, err := textTemplate.New(name).Funcs(sprig.FuncMap()).ParseFiles(cCtx.Args().Slice()...)
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}
	return tmpl, nil
}

func loadData(cCtx *cli.Context) (any, error) {
	dataMap := make(map[string]any)
	name := cCtx.String("data")
	if name == "" {
		return dataMap, nil
	}

	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()

	switch strings.ToLower(filepath.Ext(name)) {
	case ".json":
		if err := json.NewDecoder(f).Decode(&dataMap); err != nil {
			return nil, fmt.Errorf("decoding JSON: %w", err)
		}
	case ".toml":
		if _, err := toml.NewDecoder(f).Decode(&dataMap); err != nil {
			return nil, fmt.Errorf("decoding TOML: %w", err)
		}
	default:
		if err := yaml.NewDecoder(f).Decode(&dataMap); err != nil {
			return nil, fmt.Errorf("decoding YAML: %w", err)
		}
	}

	return dataMap, nil
}
