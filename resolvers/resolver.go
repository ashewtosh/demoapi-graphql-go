package resolvers

import (
	"demoApi-graphql-go/data"

	"github.com/graphql-go/graphql"
)

func TutorialList(params graphql.ResolveParams) (interface{}, error) {
	tutorial := data.Populate()
	return tutorial, nil
}
