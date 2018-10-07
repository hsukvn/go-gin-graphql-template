package controller

import (
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/hsukvn/go-gin-graphql-template/graphql/resolver"
	"github.com/hsukvn/go-gin-graphql-template/lib/util"
)

type GraphqlController struct{}

func (ctr *GraphqlController) NewHandler() *relay.Handler {
	s, err := util.LoadSchema("graphql/schema/schema.graphql")

	if err != nil {
		log.Fatalf("Fail to get schema, error: %v", err)
		return nil
	}

	schema := graphql.MustParseSchema(s, &resolver.Resolver{})

	return &relay.Handler{Schema: schema}
}

func (ctr *GraphqlController) NewGraphiQLHandlerFunc() http.HandlerFunc {
	var page = []byte(`
	<!DOCTYPE html>
	<html>
		<head>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
			<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		</head>
		<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
			<div id="graphiql" style="height: 100vh;">Loading...</div>
			<script>
				function graphQLFetcher(graphQLParams) {
					return fetch("/graphql", {
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

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write(page) })
}
