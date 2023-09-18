 ```
# Demo API

This is a simple demo API built using Go and GraphQL. It demonstrates how to create a GraphQL schema, resolvers, and a server.

## Prerequisites

To run this API, you will need:

* Go 1.17 or later
* A GraphQL client (such as GraphiQL)

## Getting Started

1. Clone this repository.
2. Run `go mod download`.
3. Run `go run server.go`.
4. Open your GraphQL client and connect to `http://localhost:9091/graphql`.

## Code Overview

The code in this repository is organized into the following files:

* `data/data.go`: This file contains the data model for the API.
* `go.mod`: This file specifies the Go dependencies for the API.
* `go.sum`: This file lists the specific versions of the Go dependencies that are used by the API.
* `query/query.go`: This file contains the GraphQL queries for the API.
* `resolvers/resolver.go`: This file contains the GraphQL resolvers for the API.
* `schema/schema.go`: This file contains the GraphQL schema for the API.
* `server.go`: This file contains the main function for the API.

## Data Model

The data model for the API is defined in the `data/data.go` file. It consists of two structs: `Tutorial` and `Comment`.

```go
type Tutorial struct {
	ID       int
	Title    string
	Author   string
	Comments []Comment
}

type Comment struct {
	Body string
}
```

## GraphQL Queries

The GraphQL queries for the API are defined in the `query/query.go` file. The `RootQuery` object defines the entry point for the API. It contains a single field, `getTutorialList`, which returns a list of all tutorials.

```go
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
```
