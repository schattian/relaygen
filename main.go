package main

import (
	"bytes"
	"flag"
	"fmt"
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
	// Mode string

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
func cp(writer io.WriteCloser, reader bytes.Buffer) error {
	fmt.Println(&reader)
	_, err := io.Copy(writer, &reader)
	if err != nil {
		return err
	}
	return nil
}

func cpToStdout(rc io.ReadCloser) error {
	_, err := io.Copy(os.Stdout, rc)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var d data
	flag.StringVar(&d.Pkg, "pkg", "relay", "The name of the package.")
	// flag.StringVar(&d.Mode, "mode", "w", "The mode used. Could be: w (default) for write, a for append")
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

	//		switch d.Mode {
	//	case "w":
	////
	//	case "a":
	//
	//	}

	rc, wc, errCh := pipe.Commands(
		exec.Command("gofmt"),
		exec.Command("goimports"),
	)

	go func() {
		select {
		case err, ok := <-errCh:
			if ok && err != nil {
				panic(err)
			}
		}
	}()

	var err error
	if err = cp(wc, out); err != nil {
		log.Fatal(err)
	}
	if err = wc.Close(); err != nil {
		log.Fatal(err)
	}
	if err = cpToStdout(rc); err != nil {
		log.Fatal(err)
	}
}
