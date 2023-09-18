package main

import (
	"demoApi-graphql-go/query"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: query.RootQuery,
	},
)

func main() {
	handle := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", handle)
	fmt.Println("Server is running on port 9091")
	listenerr := http.ListenAndServe(":9091", nil)
	if listenerr != nil {
		log.Println("ListenAndServe: ", listenerr)
	}
}
