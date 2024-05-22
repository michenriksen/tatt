package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const cheatsheet = `	
SYNTAX

  {{.}}                           - Root/current context element
  {{.Field}}                      - Field of the root/current context element
  {{if .}}...{{end}}              - Conditional block
  {{range .}}...{{end}}           - Loop over elements
  {{range $i, $v := .}}...{{end}} - Loop over elements with index
  {{range $k, $v := .}}...{{end}} - Loop over map keys and values
  {{with .}}...{{end}}            - Set context
  {{block "name" .}}...{{end}}    - Define a template block
  {{template "name" .}}           - Include another template

PREDEFINED VARIABLES

  .   - Current context
  $   - Root context

FUNCTIONS

  {{len .}}     - Length of . (slice, array, map, etc.)
  {{print .}}   - Print . as string
  {{index . i}} - Access i-th element (slice, array, map)

  See https://masterminds.github.io/sprig/ for available helper functions.

EXAMPLES

  CONDITIONAL:

  {{if .IsAdmin}}
    Welcome, Admin!
  {{else}}
    Welcome, User!
  {{end}}

  LOOP:
  
  {{range .Items}}
    Item: {{.}}
  {{end}}  

  NESTED:

  {{with .User}}
    Name: {{.Name}}
    Email: {{.Email}}
  {{end}}

  See https://pkg.go.dev/text/template for more documentation on Go templates.
`

func cheatsheetAction(cCtx *cli.Context) error {
	fmt.Fprint(cCtx.App.Writer, cheatsheet)
	return nil
}
