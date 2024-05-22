# tatt - template all the things

**tatt** is a versatile command-line tool that makes it easy to render Go text and HTML templates with dynamic data from
YAML, JSON, or TOML files.

I often use Go templates for small, one-off tasks and got tired of writing the glue code needed to parse and render them 
with data. **tatt** eliminates this hassle by providing a simple command-line interface.

```
NAME:
   tatt - template all the things

USAGE:
   tatt [options] TEMPLATE_FILE [TEMPLATE_FILE_2 ... TEMPLATE_FILE_N]

VERSION:
   0.1.0

DESCRIPTION:
   Render a Go text or html template with data loaded from a YAML, JSON, or TOML file.

AUTHOR:
   Michael Henriksen <mchnrksn@gmail.com>

COMMANDS:
   cheatsheet, cheat  view Go template cheat sheet

GLOBAL OPTIONS:
   --data FILE, -d FILE  load data from FILE
   --html                use html/template package (default: text/template)
   --version, -v         print the version
```

## Template helpers

tatt ships with over 70 useful template helper functions provided by [Masterminds/sprig](https://github.com/Masterminds/sprig).
See the [sprig documentation](https://masterminds.github.io/sprig/) for an overview of available functions.

## Installing

Install the [latest pre-compiled binary](https://github.com/michenriksen/tatt/releases/latest) for your operating 
system, or with go:

```console
go install github.com/michenriksen/tatt/cmd/tatt@latest
```
