# go-graphql-template

Basic file hierarchy of graphql API server in golang

## Query

GET

```
curl -g 'http://localhost:9527/?query={heartbeat{status}}'
```

POST

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ heartbeat { status } }"}' localhost:9527
```
