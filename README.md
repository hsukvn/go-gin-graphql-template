# go-gin-graphql-template

Basic file hierarchy of graphql API server in golang

## Requirement

* golang installed
* add $GOPATH/bin to $PATH

```
export PATH=$PATH:$GOPATH/bin
```

## Usage

* install go-bindata

```
make setup
```

* run server

```
make run
```

* build static binary

```
make
```

## Controllers

### Ping

To check whether server is alive

```
curl http://localhost:9527/ping
```

### Graphql

#### Scalar

##### Int64

This scalar type is an example showing how to add a custom scalar type and use it.`

The default int type of graphQL is 32-bit

```
 The Int scalar type represents a signed 32‐bit numeric non‐fractional value. Response formats that support a 32‐bit integer or a number type should use that type to represent this scalar.
```

So becareful using this type.

#### Resolver

##### User

###### Query

* Get user

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ user(id:2047483648) { firstname,roles { id,name } } }"}' localhost:9527/graphql
```

* Get users

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ users { id,firstname,roles { id,name } } }"}' localhost:9527/graphql
```

###### Mutation

* Add user

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "mutation { addUser(firstname: \"Mimi\", lastname: \"Lo\", roles: [\"Archeologist\"]) { id,firstname,lastname,roles { id,name } } }"}' localhost:9527/graphql
```

## Todo

- [ ] Add auth middleware ([jwt-go](https://github.com/dgrijalva/jwt-go))
