# go-graphql-template

Basic file hierarchy of graphql API server in golang

## Query

GET

```
curl -g 'http://localhost:9527/heartbeat?query={status}'
```

POST

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ status }"}' localhost:9527/heartbeat
```
