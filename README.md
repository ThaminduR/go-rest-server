# Go REST Server

A simple, configurable REST API server written in Go with Docker support. Perfect for mocking APIs, testing, and development.

## ✨ Features

- **🎯 Configurable API endpoints via JSON** - Define paths, methods, and responses without code changes
- **⚡ Fast & Lightweight** - Built with Go's standard library
- ** Request logging** - Monitor all incoming requests
- **🐳 Docker ready** - Full Docker and Docker Compose support
- **🔄 Live config updates** - Modify endpoints without rebuilding
- **📦 Minimal Docker image** - Uses Alpine Linux (~12-15 MB total)

## 🚀 Quick Start

### Option 1: Docker Compose (Recommended)

The easiest way to get started:

```bash
# Start the server
docker-compose up

# Or run in background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop the server
docker-compose down
```

Server will be available at: **http://localhost:8080**

### Option 2: Docker Build & Run

```bash
# Build the Docker image
docker build -t go-rest-server .

# Run the container
docker run -p 8080:8080 go-rest-server

# Run with custom port
docker run -p 3000:8080 go-rest-server

# Run with custom config (mount your config file)
docker run -p 8080:8080 -v $(pwd)/config.json:/root/config.json go-rest-server
```

### Option 3: Run Locally (Go Required)

First, install Go 1.21 or higher from [golang.org](https://golang.org/dl/)

```bash
# Run directly
go run main.go

# Or build and run
go build -o rest-server
./rest-server

# Run with custom port
PORT=3000 go run main.go

# Run with custom config file
CONFIG_FILE=config.example.json go run main.go
```

## 🧪 Test the Server

Once running, test the endpoints (all support GET and POST):

```bash
# Health check - GET
curl http://localhost:8080/healthcheck

# Health check - POST
curl -X POST http://localhost:8080/healthcheck

# Service A - GET
curl http://localhost:8080/servicea

# Service A - POST
curl -X POST http://localhost:8080/servicea

# Service B - GET
curl http://localhost:8080/serviceb

# Service B - POST with JSON body
curl -X POST http://localhost:8080/serviceb \
  -H "Content-Type: application/json" \
  -d '{"action":"test"}'

# Service C - GET
curl http://localhost:8080/servicec

# Service C - POST
curl -X POST http://localhost:8080/servicec
```

Or open in your browser (for GET requests):
- http://localhost:8080/healthcheck
- http://localhost:8080/servicea
- http://localhost:8080/serviceb
- http://localhost:8080/servicec

## ⚙️ Configuration

### Understanding config.json

The server reads all endpoint configurations from `config.json`:

```json
{
  "server": {
    "port": "8080"              // Default port (can be overridden by PORT env var)
  },
  "endpoints": [                 // List of all endpoints
    {
      "path": "/custom",         // URL path
      "method": "GET",           // HTTP method
      "response": {
        "status": 200,           // HTTP status code
        "body": {                // Response JSON body
          "message": "Your custom response"
        }
      }
    }
  ]
}
```

### Default Endpoints

The default `config.json` includes these endpoints (all support GET and POST):

| Methods | Path | Status | Description |
|---------|------|--------|-------------|
| GET, POST | `/healthcheck` | 200 | Health check endpoint |
| GET, POST | `/servicea` | 200 | Service A endpoint |
| GET, POST | `/serviceb` | 200 | Service B endpoint |
| GET, POST | `/servicec` | 200 | Service C endpoint |

### Adding Custom Endpoints

You can easily add your own endpoints by editing `config.json`:

**Step 1:** Open `config.json` and add a new endpoint to the `endpoints` array:

**Single method:**
```json
{
  "path": "/my-api",
  "method": "GET",
  "response": {
    "status": 200,
    "body": {
      "message": "My custom API",
      "data": [1, 2, 3],
      "custom_field": "any value"
    }
  }
}
```

**Multiple methods:**
```json
{
  "path": "/submit",
  "methods": ["POST", "PUT"],
  "response": {
    "status": 201,
    "body": {
      "message": "Resource created",
      "id": 123
    }
  }
}
```

**Step 2:** Restart the server:

```bash
# If using Docker Compose
docker-compose restart

# If using Docker
docker restart <container-id>

# If running locally
# Stop (Ctrl+C) and restart: go run main.go
```

**Step 3:** Test your new endpoint:

```bash
# Test GET request
curl http://localhost:8080/my-api

# Test POST request
curl -X POST http://localhost:8080/submit

# Test with JSON body
curl -X POST http://localhost:8080/submit \
  -H "Content-Type: application/json" \
  -d '{"name":"test"}'
```

**Note:** Responses will contain exactly what you define in the config - no additional fields are added. The server ignores request bodies and always returns the configured response.

### Configuration Examples

#### Success Response
```json
{
  "path": "/success",
  "method": "GET",
  "response": {
    "status": 200,
    "body": {
      "success": true,
      "message": "Operation completed"
    }
  }
}
```

#### Error Response
```json
{
  "path": "/not-found",
  "method": "GET",
  "response": {
    "status": 404,
    "body": {
      "error": "Resource not found",
      "code": "NOT_FOUND"
    }
  }
}
```

#### Complex Data Structure
```json
{
  "path": "/products",
  "method": "GET",
  "response": {
    "status": 200,
    "body": {
      "products": [
        {"id": 1, "name": "Product 1", "price": 29.99},
        {"id": 2, "name": "Product 2", "price": 49.99}
      ],
      "total": 2,
      "page": 1
    }
  }
}
```

### Supported HTTP Status Codes

You can use any HTTP status code in your configuration:

- **2xx Success**: 200 (OK), 201 (Created), 204 (No Content)
- **3xx Redirection**: 301 (Moved), 302 (Found), 304 (Not Modified)
- **4xx Client Errors**: 400 (Bad Request), 401 (Unauthorized), 403 (Forbidden), 404 (Not Found), 429 (Too Many Requests)
- **5xx Server Errors**: 500 (Internal Server Error), 502 (Bad Gateway), 503 (Service Unavailable)

### Supported HTTP Methods

You can configure endpoints to accept one or multiple HTTP methods:
- **GET** - Retrieve data
- **POST** - Create/submit data
- **PUT** - Update data
- **DELETE** - Delete data
- **PATCH** - Partial update

#### Single Method Example:
```json
{
  "path": "/users",
  "method": "GET",
  "response": { "status": 200, "body": {"users": []} }
}
```

#### Multiple Methods Example:
```json
{
  "path": "/data",
  "methods": ["GET", "POST"],
  "response": { "status": 200, "body": {"success": true} }
}
```

Both `GET /data` and `POST /data` will return the same configured response

## 🔧 Environment Variables

Configure the server using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` (or from config.json) |
| `CONFIG_FILE` | Path to configuration file | `config.json` |

**Examples:**

```bash
# Change port
PORT=3000 go run main.go

# Use different config file
CONFIG_FILE=config.example.json go run main.go

# Both together
PORT=9000 CONFIG_FILE=prod-config.json go run main.go
```

**With Docker:**

```bash
docker run -p 3000:3000 -e PORT=3000 go-rest-server
```

**With Docker Compose:** Edit the `environment` section in `docker-compose.yml`

## 🔄 Live Configuration Updates

When using Docker Compose, you can update endpoints without rebuilding:

1. **Edit** `config.json` on your host machine
2. **Restart** the container: `docker-compose restart`
3. **Test** your changes immediately!

The config file is mounted as a volume, so changes are reflected instantly after restart.

## 📂 Files Included

- **`config.json`** - Default endpoint configuration (4 endpoints)
- **`openapi.yaml`** - OpenAPI 3.0 specification defining all endpoints
- **`QUICKSTART.md`** - Quick reference guide

**Note:** Responses contain only the data defined in `config.json` - no automatic fields are added.

## 📁 Project Structure

```
.
├── main.go                 # Main application code (Go server)
├── go.mod                  # Go module dependencies
├── config.json             # Default endpoint configuration ⚙️
├── openapi.yaml            # OpenAPI 3.0 specification 📋
├── Dockerfile              # Docker image definition 🐳
├── docker-compose.yml      # Docker Compose setup 🐳
├── .dockerignore           # Docker build exclusions
├── .gitignore              # Git exclusions
├── README.md               # This documentation 📖
└── QUICKSTART.md           # Quick reference guide
```

## 🎯 Common Use Cases

### Mock API for Frontend Development
Use this server to mock backend APIs while developing your frontend:

1. Define your API endpoints in `config.json`
2. Start the server with `docker-compose up`
3. Point your frontend to `http://localhost:8080`

### API Testing & Prototyping
Quickly prototype API responses:

1. Add endpoints to `config.json` with desired responses
2. Test your client code against these endpoints
3. Iterate quickly without backend changes

### Simulate Error Scenarios
Test error handling in your applications:

```json
{
  "path": "/api/timeout",
  "method": "GET",
  "response": {
    "status": 504,
    "body": {"error": "Gateway timeout"}
  }
}
```

## 🐛 Troubleshooting

### Server won't start
- **Port already in use**: Change the port with `PORT=3000 docker-compose up`
- **Config file error**: Validate your JSON at [jsonlint.com](https://jsonlint.com)

### Endpoint not responding
- Check the `path` in `config.json` matches your request exactly
- Verify the HTTP method is `GET` (currently the only supported method)
- Look at server logs: `docker-compose logs`

### Config changes not applied
- Restart the container: `docker-compose restart`
- If still not working, rebuild: `docker-compose up --build`

### Docker issues
- **Docker not running**: Start Docker Desktop
- **Permission denied**: Try `sudo docker-compose up` (Linux)
- **Port conflict**: Change port in docker-compose.yml or use `-p 9000:8080`

## � API Documentation

### OpenAPI Specification

The API is fully documented using OpenAPI 3.0 specification in `openapi.yaml`.

**View the API documentation:**
- Copy `openapi.yaml` to [Swagger Editor](https://editor.swagger.io/)
- Or use any OpenAPI viewer/tool

**Key features documented:**
- All endpoint paths and methods
- Request/response schemas
- Status codes
- Example responses
- Error responses

### Interactive API Testing

You can use tools like:
- **Swagger UI** - Visual documentation and testing
- **Postman** - Import the OpenAPI spec for testing
- **Insomnia** - REST client with OpenAPI support
- **curl** - Command-line testing (examples in README)

##  Current Endpoints

The server includes 4 configurable endpoints:
- `/healthcheck` - Health check endpoint
- `/servicea` - Service A simulation
- `/serviceb` - Service B simulation  
- `/servicec` - Service C simulation

All endpoints support both GET and POST methods and return the configured JSON responses.

## 🤝 Contributing

Feel free to submit issues and enhancement requests!

## 📄 License

This project is open source and available for any use.

---

## 📋 Quick Commands Cheat Sheet

### Using Docker Compose
```bash
docker-compose up              # Start server
docker-compose up -d           # Start in background
docker-compose down            # Stop server
docker-compose restart         # Restart
docker-compose logs -f         # View logs
```

### Using Docker
```bash
docker build -t go-rest-server .           # Build image
docker run -p 8080:8080 go-rest-server     # Run container
docker ps                                   # List running containers
docker stop <container-id>                  # Stop container
```

### Testing Endpoints
```bash
curl http://localhost:8080/healthcheck
curl http://localhost:8080/servicea
curl http://localhost:8080/serviceb
curl http://localhost:8080/servicec
```

### Local Development (Go required)
```bash
go run main.go                             # Run directly
go build -o rest-server && ./rest-server   # Build and run
PORT=3000 go run main.go                   # Custom port
```

---

**Happy API Mocking! 🚀**
