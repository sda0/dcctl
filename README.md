# Installation instruction
### Build
```
go build -o /usr/bin/dcctl main.go
```
### Create your local config file
```
cp .dcctl.yaml.example $HOME/.dcctl.yaml
```
### Setup config variables
```
# cat $HOME/.dcctl.yaml
project_dock: /path/to/docker/compose/files
project: /path/to/project/src/
```
# How to use
## How to add new service containers
You should create new docker-<servicename>.yml file in your project docker dir.
Thats enough.

## Commands
- Show service containers launched
```
dcctl
```
- Up general services (services listed in docker_compose.yaml file)
```
dcctl up
```

- Start general services (services listed in docker_compose.yaml file)
```
dcctl start
```

- Stop all launched service containers
```
dcctl stop
```

- Down all launched service containers
```
dcctl down
```

- Start some service (it may launch general services if any of these stopped)
```
dcctl start inapp
```

- Stop general services + some additional service
```
dcctl stop inapp
```
