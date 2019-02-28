# GraphQL x Go x Postgres x Elasticsearch

Simple skeleton for a Go GraphQL API connected to a Postgres database and Elasticsearch cluster.

## Contributing

### Requirements
- Go v1.11
```sh
brew install go
```

- Postgres v11
```sh
brew install postgresql
```

- ElasticSearch v6.6
```sh
brew install elasticsearch
```

### Setup
Build Go module:
```sh
go build ./...
```

### Run
Start the Elasticsearch cluster:
```sh
elasticsearch
```

Start the GraphQL server:
```sh
go run server/server.go
```

Try some operations:
```graphql
query sayHello {
  hello(id:"1")
}

query searchFor {
  search(query:"luke") {
    id
    name
  }
}
```

## Architecture
- `db.go` - Database setup and connection
- `model.go` - Application type definitions
- `resolver.go` - Resolve API operations, with access to database and Elasticsearch cluster
- `schema.graphql` - GraphQL schema in SDL
- `server/server.go` - Serve GraphQL API
- `util/schema.go` - Transform SDL to parseable string
