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

#### User

##### Query

Get user

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ user(id:1) { firstname,roles { id,name } } }"}' localhost:9527/graphql
```

Get users

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ users { id,firstname,roles { id,name } } }"}' localhost:9527/graphql
```

##### Mutation

Add user

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "mutation { addUser(firstname: \"Mimi\", lastname: \"Lo\", roles: [\"Archeologist\"]) { id,firstname,lastname,roles { id,name } } }"}' localhost:9527/graphql
```

## Todo

- [ ] Add auth middleware ([jwt-go](https://github.com/dgrijalva/jwt-go))
