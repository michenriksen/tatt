package cli_test

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/michenriksen/tatt/internal/cli"
)

func TestRenderAction(t *testing.T) {
	tt := []struct {
		name string
		args []string
	}{
		{
			"simple template with json data",
			[]string{"--data", filepath.Join("testdata", "data.json"), filepath.Join("testdata", "simple.tmpl")},
		},
		{
			"simple template with yaml data",
			[]string{"--data", filepath.Join("testdata", "data.yaml"), filepath.Join("testdata", "simple.tmpl")},
		},
		{
			"simple template with toml data",
			[]string{"--data", filepath.Join("testdata", "data.toml"), filepath.Join("testdata", "simple.tmpl")},
		},
		{
			"semgrep example",
			[]string{
				"--data", filepath.Join("..", "..", "examples", "semgrep_aws-sdk-js-clients.data.yaml"),
				filepath.Join("..", "..", "examples", "semgrep_aws-sdk-js-clients.yaml.tmpl"),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			outBuf := new(bytes.Buffer)
			errBuf := new(bytes.Buffer)

			app := cli.App()
			app.Writer = outBuf
			app.ErrWriter = errBuf

			err := app.Run(append([]string{"tatt"}, tc.args...))
			if err != nil {
				t.Errorf("expected no error, but got: %v", err)
			}

			if errBuf.Len() != 0 {
				t.Errorf("expected empty stderr, but is: %q", errBuf.String())
			}

			expectGolden(t, outBuf)
		})
	}
}

func expectGolden(tb testing.TB, actual *bytes.Buffer) {
	tb.Helper()

	name := filepath.Join("testdata", "golden", tb.Name()+".golden")

	if _, ok := os.LookupEnv("UPDATE_GOLDEN"); ok {
		if err := os.MkdirAll(filepath.Dir(name), 0o744); err != nil {
			tb.Fatalf("error creating directory for golden file: %v", err)
		}

		if err := os.WriteFile(name, actual.Bytes(), 0o644); err != nil {
			tb.Fatalf("error writing data to golden file: %v", err)
		}

		tb.Logf("wrote %d bytes to golden file %s", actual.Len(), name)
	}

	golden, err := os.ReadFile(name)
	if err != nil {
		tb.Fatalf("error reading golden file: %v", err)
	}

	if bytes.Compare(golden, actual.Bytes()) != 0 {
		diff := getDiff(tb, actual)
		tb.Errorf("expected data to match golden data; run tests with UPDATE_GOLDEN=1 to update\n\nDIFF:\n\n%s", diff)
		return
	}

	tb.Logf("actual data matches golden data")
}

func getDiff(tb testing.TB, actual *bytes.Buffer) string {
	tb.Helper()

	actualFile := filepath.Join(tb.TempDir(), "data.actual")
	goldenFile := filepath.Join("testdata", "golden", tb.Name()+".golden")

	if err := os.WriteFile(actualFile, actual.Bytes(), 0o644); err != nil {
		tb.Fatalf("error writing actual data to temporary file: %v", err)
	}

	diff, _ := exec.Command("git", "diff", "--no-index", "--color=always", goldenFile, actualFile).CombinedOutput()

	return string(diff)
}
