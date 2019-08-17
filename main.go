package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/joncalhoun/pipe"
)

type data struct {
	Pkg  string
	Name string
	Type string

	RenderCursorTemplate bool
}

func pipeline(funcs ...func() error) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var d data
	flag.StringVar(&d.Pkg, "pkg", "main", "The name of the package.")
	flag.StringVar(&d.Name, "name", "", "The name of the type to generate for.")
	flag.StringVar(&d.Type, "type", "", "The type to generate for.")
	flag.BoolVar(&d.RenderCursorTemplate, "cursor", false, "Generate cursor template")
	flag.Parse()

	var out bytes.Buffer

	if err := nodeTemplate.Execute(&out, d); err != nil {
		log.Fatal(err)
	}
	if d.RenderCursorTemplate {
		if err := cursorTemplate.Execute(&out, d); err != nil {
			log.Fatal(err)
		}
	}

	rc, wc, _ := pipe.Commands(
		exec.Command("gofmt"),
		exec.Command("goimports"),
	)

	if err := pipeline(func() error {
		_, err := io.Copy(wc, &out)
		return err
	}, wc.Close, func() error {
		_, err := io.Copy(os.Stdout, rc)
		return err
	}); err != nil {
		log.Fatal(err)
	}
}
