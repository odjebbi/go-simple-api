# Go Simple API

Simple REST API built with Go for buildpack deployment.

## Endpoints

- `GET /` - Home page
- `GET /api/hello` - Hello message (JSON)
- `GET /health` - Health check
- `GET /api/info` - API information

## Local Development
```bash
# Run application
go run cmd/server/main.go

# Build binary
go build -o go-simple-api cmd/server/main.go
./go-simple-api
```

## Deploy with Buildpacks

This app is ready to deploy with Go buildpack.
The buildpack will automatically detect Go and build the binary.
