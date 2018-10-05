# go-graphql-template

Basic file hierarchy of graphql API server in golang

## Controllers

### Ping

To check whether server is alive

```
curl http://localhost:9527/ping
```

### Graphql

#### User

##### Query

Get users

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ users { id,firstname,roles { id,name } } }"}' localhost:9527
```

##### Mutation

Create user

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "mutation { createUser(firstname: \"Mimi\", lastname: \"Lo\") { id,firstname,lastname } }"}' localhost:9527
```

## Todo

- [ ] Add auth middleware ([jwt-go](https://github.com/dgrijalva/jwt-go))
