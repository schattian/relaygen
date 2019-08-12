package relay

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

type Node interface {
	IsNode()
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type NodeWriter struct{}

func NewNodeWriter() *NodeWriter {
	return &NodeWriter{}
}

func (w *NodeWriter) Name() string {
	return "relayNode"
}

func (w *NodeWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	return
}

func (w *NodeWriter) Write(wr io.Writer, t typewriter.Type) error {
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
	if err := typewriter.Register(NewNodeWriter()); err != nil {
		panic(err)
	}
}
