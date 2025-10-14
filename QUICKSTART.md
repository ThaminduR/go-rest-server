# Quick Start Guide

## 1. Basic Usage

Start the server with default configuration:
```bash
docker-compose up
```

Test the endpoints:
```bash
curl http://localhost:8080/health
curl http://localhost:8080/hello
curl http://localhost:8080/users
```

## 2. Customize Your API

Edit `config.json` to add your own endpoints:

```json
{
  "path": "/my-api",
  "method": "GET",
  "response": {
    "status": 200,
    "body": {
      "message": "My custom API response"
    }
  }
}
```

Restart to apply changes:
```bash
docker-compose restart
```

## 3. Configuration Options

### Supported HTTP Status Codes
- `200` - OK
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `500` - Internal Server Error

### Supported HTTP Methods
Currently supported: `GET`

### Response Body
Can contain any valid JSON structure:
- Objects
- Arrays
- Nested structures
- Numbers, strings, booleans

## 4. Examples

### Simple Success Response
```json
{
  "path": "/success",
  "method": "GET",
  "response": {
    "status": 200,
    "body": {
      "success": true
    }
  }
}
```

### List of Items
```json
{
  "path": "/items",
  "method": "GET",
  "response": {
    "status": 200,
    "body": {
      "items": ["item1", "item2", "item3"],
      "count": 3
    }
  }
}
```

### Error Response
```json
{
  "path": "/not-found",
  "method": "GET",
  "response": {
    "status": 404,
    "body": {
      "error": "Resource not found"
    }
  }
}
```

## 5. Advanced Usage

### Use Custom Config File
```bash
CONFIG_FILE=config.example.json go run main.go
```

### Change Port
```bash
PORT=3000 docker-compose up
```

### View Logs
```bash
docker-compose logs -f
```

### Stop Server
```bash
docker-compose down
```

## 6. Troubleshooting

**Server won't start:**
- Check if port 8080 is already in use
- Verify `config.json` is valid JSON

**Endpoint not responding:**
- Check the path in `config.json` matches your request
- Verify the HTTP method matches

**Config changes not applied:**
- Restart the container: `docker-compose restart`
- Or rebuild: `docker-compose up --build`
