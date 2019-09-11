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
	Pkg     string
	Name    string
	Marshal string
	Type    string
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

func main() {
	var d data
	flag.StringVar(&d.Pkg, "pkg", "relay", "String. The name of the package where the generated entities will live. Default: relay.")
	// flag.StringVar(&d.Mode, "mode", "w", "The mode used. Could be: w (default) for write, a for append")
	flag.StringVar(&d.Marshal, "marshal", "bson", "String. The marshaling mode for the generated fields. Default: bson.")
	flag.StringVar(&d.Name, "name", "", "String. The name of the entity to generate its relay types. Required.")
	flag.StringVar(&d.Type, "type", "", "String. The entity type used in your GQL pipelines (usually the pointer to the entity). Required.")
	flag.BoolVar(&d.RenderCursorTemplate, "cursor", false, "Boolean. Generate cursor template. Default: false.")
	flag.Parse()

	var out bytes.Buffer

	err := nodeTemplate.Execute(&out, d)
	if err != nil {
		log.Fatal(err)
	}

	if d.RenderCursorTemplate {
		err := cursorTemplate.Execute(&out, d)
		if err != nil {
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
		//		exec.Command("goimports"),
	)

	go func(ch <-chan error) {
		select {
		case err := <-ch:
			if err != nil {
				panic(err)
			}
		}
		return
	}(errCh)
	panic(&out)

	_, err = io.Copy(wc, &out)
	if err != nil {
		log.Fatal(err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, rc)
	if err != nil {
		log.Fatal(err)
	}

}
