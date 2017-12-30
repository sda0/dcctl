# Docker compose control utility

## Install to Dock directory (recommended way)

build executable file to Dock directory and symlink by /usr/lib/dcctl:
```
go build -o ~/path_to_dock_dir/dcctl main.go
ln -s ~/path_to_dock_dir/dcctl /usr/bin/dcctl
```
## Install to any other directory (not recommended way)
You should build executable file to any path:
```
go build -o /any_dir/any_name main.go
```
Then create your local config file  $HOME/.dcctl.yaml
```
cp .dcctl.yaml.example $HOME/.dcctl.yaml
```
Then setup config variables in config file
```
# cat $HOME/.dcctl.yaml
project_dock=<path to dock dir>
project=<path to project src dir>
```
# How to use
## How to add new service containers
You should create new docker-<servicename>.yml file in your dock dir.
That's enough.

## Commands
- Show service containers launched
```
dcctl
```
- Up general services (nginx and all links from it)
```
dcctl up
```

- Start general services (nginx and all links from it)
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
