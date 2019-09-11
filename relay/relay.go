package relay

import "fmt"

type ID interface {
	fmt.Stringer
}

type Node interface {
	IsNode()
	GetID() ID
}

type PageInfo struct {
	HasNextPage     bool    `bson:"hasNextPage"`
	HasPreviousPage bool    `bson:"hasPreviousPage"`
	StartCursor     *string `bson:"startCursor"`
	EndCursor       *string `bson:"endCursor"`
}

func NewString(s string) *string {
	return &s
}
