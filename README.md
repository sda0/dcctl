# Installation instruction
### Build
```
go build -o /usr/bin/pwctl main.go
```
### Create your local config file
```
cp .pwctl.yaml.example $HOME/.pwctl.yaml
```
### Setup config variables
```
# cat $HOME/.pwctl.yaml
PW_DOCK=<path to powodock dir>
PW_HOME=<path to pw home dir>
```
# How to use
## How to add new service containers
You should create new docker-<servicename>.yml file in your powodock dir.
Thats enough.

## Commands
- Show service containers launched
```
pwctl
```
- Up general services (nginx, web, api, cli, redis, rabbit, pg)
```
pwctl up
```

- Start general services (nginx, web, api, cli, redis, rabbit, pg)
```
pwctl start
```

- Stop all launched service containers
```
pwctl stop
```

- Down all launched service containers
```
pwctl down
```

- Start some service (it may launch general services if any of these stopped)
```
pwctl start inapp
```

- Stop general services + some additional service
```
pwctl stop inapp
```
