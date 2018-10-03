# go-graphql-template

Basic file hierarchy of graphql API server in golang

## Query

GET

```
curl -g 'http://localhost:9527/?query={heartbeat{status}}'
curl -g 'http://localhost:9527/?query={user{id,firstname,roles{id,name}}}'
```

POST

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ heartbeat { status } }"}' localhost:9527
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ user { id,firstname,roles { id,name } } }"}' localhost:9527
```

## Mutation

POST

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "mutation { createUser(firstname: \"Mimi\", lastname: \"Lo\") { id,firstname,lastname } }"}' localhost:9527
```
