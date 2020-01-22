package main
import (
  "log"
  "fmt"
  "net/http"
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "app/queries"
  "app/mutations"
	"os"
)
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
  Query: queries.RootQuery,
  Mutation: mutations.RootMutation,
})
func main() {
	port := 3000
  h := handler.New(&handler.Config{
    Schema: &schema,
    Pretty: true,
  })
  http.Handle("/graphql", disableCors(h))
  log.Println("Now server is running on port %d", port)
  if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
    log.Println("Error binding to port %d", port)
  	os.Exit(1)
  }
}
func disableCors(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
  if r.Method == "OPTIONS" {
    w.Header().Set("Access-Control-Max-Age", "86400")
    w.WriteHeader(http.StatusOK)
    return
  }
  h.ServeHTTP(w, r)
 })
}