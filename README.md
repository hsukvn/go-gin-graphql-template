# go-gin-graphql-template

GraphQL API server in golang to get linux system info.

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
query {
    user(uid: "0") {
        uid
        gid
        name
        home
        groups {
            gid
            name
        }

    }
}
```

Get user by user name

```
query {
    userByName(name: "root") {
        uid
        gid
        name
        home
        groups {
            gid
            name
        }

    }
}
```

Get users

```
query {
    users {
        uid
        gid
        name
        home
        groups {
            gid
            name
        }

    }
}
```

## Query using curl

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ user(uid:\"0\") { uid, gid, name, home, groups { gid, name } } }"}' localhost:9527/graphql
```


## Todo

- [ ] Add auth middleware ([jwt-go](https://github.com/dgrijalva/jwt-go))
