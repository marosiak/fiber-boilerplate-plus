## Development

### Start the application 


```bash
go run cmd/main.go
```

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
