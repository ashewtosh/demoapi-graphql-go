package query

import (
	"demoApi-graphql-go/resolvers"
	"demoApi-graphql-go/schema"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getTutorialList": &graphql.Field{
			Type:        graphql.NewList(schema.TutorialType),
			Description: "Get Tutorial list",
			Resolve:     resolvers.TutorialList,
		},
	},
})
