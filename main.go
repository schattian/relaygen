package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"text/template"

	"github.com/joncalhoun/pipe"
)

func main() {
	var wg sync.WaitGroup
	var err error
	var out bytes.Buffer
	var d data

	flag.StringVar(&d.Pkg, "pkg", "relay", "String. The name of the package where the generated entities will live. Default: relay.")
	flag.StringVar(&d.Marshal, "marshal", "bson", "String. The marshaling mode for the generated fields. Default: bson.")
	flag.StringVar(&d.Name, "name", "", "String. The name of the entity to generate its relay types. Required if not base.")
	flag.StringVar(&d.Type, "type", "", "String. The entity type used in your GQL pipelines (usually the pointer to the entity w/pkg name). Required if not base.")
	flag.BoolVar(&d.IsSDL, "sdl", false, "Boolean. Generate the SDL into a .graphql file for the desired template. Default: false.")
	flag.BoolVar(&d.RenderBaseTemplate, "base", false, "Boolean. Generate the base template with the common interfaces. Default: false.")
	flag.Parse()

	err = d.validate()

	if err != nil {
		log.Fatal(err)
	}

	err = d.selectTemplate().Execute(&out, d)

	if err != nil {
		log.Fatal(err)
	}

	cmds := d.commands()

	// BREAKPOINT: If there arent commands, ommit the processing
	if len(cmds) == 0 {
		fmt.Println(&out)
		return
	}

	rc, wc, errCh := pipe.Commands(cmds...)

	wg.Add(1)
	go func(ch <-chan error, wg *sync.WaitGroup) {
		defer wg.Done()
		select {
		case err := <-ch:
			if err != nil {
				panic(err)
			}
		}
		return
	}(errCh, &wg)

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

	wg.Wait()
}

type data struct {
	Pkg     string
	Name    string
	Marshal string
	Type    string

	RenderBaseTemplate bool
	IsSDL              bool
}

func (d *data) commands() []*exec.Cmd {
	if d.IsSDL {
		return []*exec.Cmd{}
	}
	return []*exec.Cmd{exec.Command("gofmt"), exec.Command("goimports")}
}

func (d *data) validate() (err error) {
	if d.RenderBaseTemplate {
		if d.Name == "" || d.Type == "" {
			return errors.New("Name or Type flags cannot be empty when calling a non-base generation")
		}
	}
	return nil
}

func (d *data) selectTemplate() (temp *template.Template) {
	if d.RenderBaseTemplate {
		if d.IsSDL {
			temp = sdlBaseTemplate
		} else {
			temp = baseTemplate
		}
	} else {
		if d.IsSDL {
			temp = sdlEntityTemplate
		} else {
			temp = entityTemplate
		}
	}
	return temp
}
