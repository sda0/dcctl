# PWCTL

## Install to PW Dock directory (recommended way)

build executable file to PW Dock directory and symlink by /usr/lib/pwctl:
```
go build -o /path_to_pw_dock_dir/pwctl main.go
ln -s /path_to_pw_dock_dir/pwctl /usr/bin/pwctl
```
## Install to any other directory (not recommended way)
You should build executable file to any path:
```
go build -o /any_dir/any_name main.go
```
Then create your local config file  $HOME/.pwctl.yaml
```
cp .pwctl.yaml.example $HOME/.pwctl.yaml
```
Then setup config variables in config file
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
