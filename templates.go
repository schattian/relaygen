package relay

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	node,
	cursor,
}

var node = &typewriter.Template{
	Name: "RelayNode",
	Text: `
// {{.Name}}Edge is the edge representation of {{.Name}}
type {{.Name}}Edge struct {
	Cursor string ` + "`json:\"cursor\"`" + `
	Node {{.Pointer}}{{.Name}} ` + "`json:\"node\"`" + `
}

// {{.Name}}Connection is the connection containing edges of {{.Name}}
type {{.Name}}Connection struct {
	Edges []{{.Name}}Edge ` + "`json:\"edges\"`" + `
	PageInfo relay.PageInfo ` + "`json:\"pageInfo\"`" + `
	TotalCount *int ` + "`json:\"totalCount\"`" + `
}
`,
}

var cursor = &typewriter.Template{
	Name: "RelayCursor",
	Text: `
// {{.Name}}Cursor is the edge representation of {{.Name}}
type {{.Name}}Cursor struct {
	Offset int ` + "`json:\"offset\"`" + `
	ID string ` + "`json:\"id\"`" + `
}

func New{{.Name}}Cursor(offset int, id fmt.Stringer) *{{.Name}}Cursor {
	return &{{.Name}}Cursor{Offset: offset, ID: id.String()}
}

func Encode{{.Name}}Cursor(cursor {{.Name}}Cursor) string {
	b, err := msgpack.Marshal(cursor)
	if err != nil {
		panic("unable to marshal cursor: " + err.Error())
	}
	return base64.StdEncoding.EncodeToString(b)
}

func Decode{{.Name}}Cursor(cursor string) (*{{.Name}}Cursor, error) {
	b, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode cursor")
	}
	
	var out {{.Name}}Cursor
	if err := errors.Wrap(msgpack.Unmarshal(b, &out), "unable to unmarshal {{.Name}} cursor"); err != nil {
		return nil, err
	}
	return &out, nil
}

func MustDecode{{.Name}}Cursor(cursor string) *{{.Name}}Cursor {
	decoded, err := Decode{{.Name}}Cursor(cursor)
	if err != nil {
		panic("unable to decode cursor: " + err.Error())
	}
	return decoded
}
`,
}
