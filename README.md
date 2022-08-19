# Boilerplate
First of all, you will probably need to cut this header Boilerplate header part


Then you can find & replace these names:
```
project_module
project-name
```


## Development

### Start the application 

### Basic
```bash
go run cmd/main.go
```

### Live reload with [Air](https://github.com/cosmtrek/air)

```bash
go install github.com/cosmtrek/air@latest
air
```

Alternatively:
```bash
make air
```

For Windows you will need to use different configuration:
```bash
air -c windows.air.toml
```
You may need to add `$GOPATH/bin` to your `$PATH`
### Use local container

```
# Clean packages
make clean-packages

# Generate go.mod & go.sum files
make requirements

# Generate docker image
make build

# Generate docker image with no cache
make build-no-cache

# Run the projec in a local container
make up

# Run local container in background
make up-silent

# Stop container
make stop

# Start container
make start
```

## Production

```bash
docker build -t gofiber .
docker run -d -p 8000:8000 gofiber
```

Go to http://localhost:8000:


![Go Fiber Docker Boilerplate](./go_fiber_boilerplate.gif)
