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

#### CPU

Get CPU

```
query {
    cpu(id:"cpu1") {
        id
        total
        user
        system
        idle
        iowait
    }
}
```

Get CPUs

```
query {
    cpus {
        id
        total
        user
        system
        idle
        iowait
    }
}
```

#### Memory

Get Memory

```
query {
    memory {
        total
        free
        used
        shared
        buffer
        cache
        swap
    }
}
```

#### Network Iface

Get Iface

```
query {
    iface(name: "eno1") {
        name
        mac
        addrv4 {
            ip
            mask
        }
        addrv6 {
            ip
            mask
        }
        mtu
        rx
        tx
    }
}
```

Get Ifaces

```
query {
    ifaces {
        name
        mac
        addrv4 {
            ip
            mask
        }
        addrv6 {
            ip
            mask
        }
        mtu
        rx
        tx
    }
}
```

#### Service

Get service

```
query {
    service (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

Start service

```
mutation {
	startService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
  }
}
```

Stop service

```
mutation {
    stopService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

Enable service

```
mutation {
    enableService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

Disable service

```
mutation {
    disableService (name: "smb") {
        name
        mainPID
        activeState
        unitFileState
    }
}
```

## Query using curl

```
curl -X POST -H 'Content-Type: application/json' -d '{"query": "{ user(uid:\"0\") { uid, gid, name, home, groups { gid, name } } }"}' localhost:9527/graphql
```

## Todo

- [ ] Add auth middleware ([jwt-go](https://github.com/dgrijalva/jwt-go))
