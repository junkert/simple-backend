package queries
import (
  "github.com/graphql-go/graphql"
	fields "github.com/junkert/simple-backend/queries/fields"
)
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootQuery",
  Fields: graphql.Fields{
    "getNotTodos": fields.GetNotTodos,
  },
})