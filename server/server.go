package main

import (
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/lukehedger/graphql-go-postgres"
	"github.com/lukehedger/graphql-go-postgres/util"
)

var schema *graphql.Schema

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write(page)
}

func init() {
	db, err := api.OpenDB()
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}

	es, err := api.OpenES()
	if err != nil {
		log.Fatalf("Unable to connect to elasticsearch: %s", err)
	}

	s, err := util.GetSchema("./schema.graphql")
	if err != nil {
		log.Fatalf("Unable to read GraphQL schema: %s \n", err)
	}

	schema = graphql.MustParseSchema(s, &api.Resolver{DB: db, ES: es})
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))

	http.Handle("/query", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":4000", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.12.0/graphiql.min.css" rel="stylesheet" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.8.2/umd/react.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.8.2/umd/react-dom.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.12.0/graphiql.min.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
