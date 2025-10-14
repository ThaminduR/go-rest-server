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

Once running, test the endpoints:

```bash
# Health check
curl http://localhost:8080/health

# Hello endpoint
curl http://localhost:8080/hello

# Sample data
curl http://localhost:8080/data

# User list
curl http://localhost:8080/users

# API status
curl http://localhost:8080/api/status

# Error simulation
curl http://localhost:8080/error
```

Or open in your browser:
- http://localhost:8080/health
- http://localhost:8080/users

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

The default `config.json` includes these endpoints:

| Method | Path | Status | Description |
|--------|------|--------|-------------|
| GET | `/health` | 200 | Health check endpoint |
| GET | `/hello` | 200 | Hello world message |
| GET | `/data` | 200 | Sample data array |
| GET | `/users` | 200 | Sample user list |
| GET | `/api/status` | 200 | API status information |
| GET | `/error` | 500 | Simulated error response |

### Adding Custom Endpoints

You can easily add your own endpoints by editing `config.json`:

**Step 1:** Open `config.json` and add a new endpoint to the `endpoints` array:

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
curl http://localhost:8080/my-api
```

**Note:** Responses will contain exactly what you define in the config - no additional fields are added.

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

Currently supported:
- **GET** - Retrieve data

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

- **`config.json`** - Default endpoint configuration (6 endpoints)
- **`config.example.json`** - Example with more advanced configurations
- **`QUICKSTART.md`** - Quick reference guide

**Note:** Responses contain only the data defined in `config.json` - no automatic fields are added.

## 📁 Project Structure

```
.
├── main.go                 # Main application code (Go server)
├── go.mod                  # Go module dependencies
├── config.json             # Default endpoint configuration ⚙️
├── config.example.json     # Example configurations
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

## 📝 Examples

See `config.example.json` for more configuration examples including:
- Product catalog endpoints
- Weather API simulation
- Authentication error responses (401, 403)
- Various HTTP status codes

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
curl http://localhost:8080/health
curl http://localhost:8080/hello
curl http://localhost:8080/users
```

### Local Development (Go required)
```bash
go run main.go                             # Run directly
go build -o rest-server && ./rest-server   # Build and run
PORT=3000 go run main.go                   # Custom port
```

---

**Happy API Mocking! 🚀**
