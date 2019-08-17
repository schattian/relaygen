package relay

import "fmt"

type ID interface{
	fmt.Stringer
}

type Node interface {
	IsNode()
	GetID() ID
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

func NewString(s string) *string {
	return &s
}
