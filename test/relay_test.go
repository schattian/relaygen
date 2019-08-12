package main

import (
	"os"
	"strings"
	"testing"

	"github.com/clipperhouse/typewriter"
	"github.com/stretchr/testify/require"

	_ "github.com/hookactions/gqlgen-relay"
)

func TestNodeWriter_Write(t *testing.T) {
	filter := func(f os.FileInfo) bool {
		return !strings.HasSuffix(f.Name(), "_test.go") && !strings.HasSuffix(f.Name(), "_set.go")
	}

	app, err := typewriter.NewAppFiltered("+test", filter)
	require.NoError(t, err)

	_, err = app.WriteAll()
	require.NoError(t, err)
}
