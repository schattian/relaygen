package relay

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

type CursorWriter struct{}

func NewCursorWriter() *CursorWriter {
	return &CursorWriter{}
}

func (w *CursorWriter) Name() string {
	return "relayCursor"
}

func (w *CursorWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	result = append(result, typewriter.ImportSpec{
		Path: "github.com/vmihailenco/msgpack",
	})
	result = append(result, typewriter.ImportSpec{
		Path: "fmt",
	})
	result = append(result, typewriter.ImportSpec{
		Path: "github.com/pkg/errors",
	})
	return
}

func (w *CursorWriter) Write(wr io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(w)

	if !found {
		// nothing to be done
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(wr, t); err != nil {
		return err
	}

	return nil
}

func init() {
	if err := typewriter.Register(NewCursorWriter()); err != nil {
		panic(err)
	}
}
