package mutations
import (
  "github.com/graphql-go/graphql"
  fields "github.com/junkert/simple-backend/mutations/fields"
)
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootMutation",
  Fields: graphql.Fields{
    "createNotTodo": fields.CreateNotTodo,
  },
})