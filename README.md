# A2A Go SDK Demo Project

This is a complete demonstration project using the [A2A Go SDK](https://github.com/yeeaiclub/a2a-go), showcasing how to build a simple Agent-to-Agent (A2A) communication system.

## 🚀 Project Features

- **Complete Client-Server Architecture**: Includes independent client and server applications
- **Interactive Message Sending**: Supports user input for custom messages
- **Comprehensive Error Handling**: Provides detailed error information and logging
- **Elegant Configuration Management**: Supports environment variables and custom configuration
- **Concurrency-Safe Queue Management**: Uses mutex locks to protect shared resources
- **Beautiful Console Output**: Uses emojis and formatted output to enhance user experience
- **Complete Test Coverage**: Includes unit tests and integration tests
- **Convenient Build Tools**: Provides Makefile to simplify building and running
- **Detailed Documentation**: Includes complete usage instructions and API documentation

## 📁 Project Structure

```
demo/
├── basic/
│   ├── client/           # Client application
│   │   ├── main.go       # Client main program
│   │   ├── config.go     # Configuration file
│   │   ├── main_test.go  # Client tests
│   │   ├── go.mod        # Go module file
│   │   └── go.sum        # Dependency checksum file
│   └── server/           # Server application
│       ├── main.go       # Server main program
│       ├── config.go     # Configuration file
│       ├── executor.go   # Task executor
│       ├── queue.go      # Queue manager
│       ├── main_test.go  # Server tests
│       ├── go.mod        # Go module file
│       └── go.sum        # Dependency checksum file
├── scripts/
│   └── demo.sh           # Automatic demo script
├── Makefile              # Build script
├── config.example        # Configuration example
└── README.md             # Project documentation
```

## 🛠️ Installation and Running

### Prerequisites

- Go 1.19 or higher
- Git

### 1. Clone the project

```bash
git clone <repository-url>
cd demo
```

### 2. Start the server

```bash
# Using Makefile (recommended)
make run-server

# Or manually
cd basic/server
go mod tidy
go run .
```

The server will start at `http://localhost:8080`, and you'll see output similar to:

```
🎯 A2A Go SDK Server Demo
🚀 Starting A2A server...
📋 Agent Information:
   Name: Print Agent
   Description: A simple print agent for A2A Go SDK demonstration
   Version: v0.1.0
🌐 Server URL: http://localhost:8080
📄 Agent Card: http://localhost:8080/card
🔌 API Endpoint: http://localhost:8080/api
```

### 3. Run the client

In a new terminal window:

```bash
# Using Makefile (recommended)
make run-client

# Or manually
cd basic/client
go mod tidy
go run .
```

### 4. Automatic Demo (Recommended)

Use the automatic demo script to start the complete demonstration with one command:

```bash
# Using Makefile
make demo-auto

# Or run the script directly
chmod +x scripts/demo.sh
./scripts/demo.sh
```

The automatic demo script will:
- Check Go environment
- Install dependencies
- Start the server
- Run the client
- Automatically clean up resources

The client will connect to the server and start the demonstration:

```
🎯 A2A Go SDK Client Demo
🚀 Starting A2A client demonstration...
Getting agent card information...
Successfully got agent card: Print Agent (version: v0.1.0)
📋 Agent Information:
   Name: Print Agent
   Description: A simple print agent for A2A Go SDK demonstration
   Version: v0.1.0

📤 Sending default message...
Sending message to task 1: hello, world
Message sent successfully, task ID: 1
✅ Default message sent successfully, task ID: 1

💬 Now you can send custom messages:
Enter message to send: 
```

## 🔧 Configuration Options

### Environment Variable Configuration

The project supports configuration through environment variables. Copy `config.example` to `.env` and modify as needed:

```bash
cp config.example .env
```

### Client Configuration

Supported environment variables:

- `A2A_SERVER_URL`: Server URL (default: http://localhost:8080)
- `A2A_AGENT_CARD_PATH`: Agent card path (default: card)
- `A2A_API_PATH`: API path (default: /api)
- `A2A_TIMEOUT_SECONDS`: Request timeout in seconds (default: 30)

### Server Configuration

Supported environment variables:

- `A2A_SERVER_PORT`: Server port (default: 8080)
- `A2A_CARD_PATH`: Agent card path (default: /card)
- `A2A_API_PATH`: API path (default: /api)
- `A2A_AGENT_NAME`: Agent name (default: Print Agent)
- `A2A_AGENT_DESC`: Agent description
- `A2A_AGENT_VERSION`: Agent version (default: v0.1.0)

### Code Configuration

You can also modify configuration directly in the code:

```go
// Client configuration
type Config struct {
    ServerURL     string        // Server URL
    AgentCardPath string        // Agent card path
    APIPath       string        // API path
    Timeout       time.Duration // Request timeout
}

// Server configuration
type Config struct {
    Port          int    // Server port
    CardPath      string // Agent card path
    APIPath       string // API path
    AgentName     string // Agent name
    AgentDesc     string // Agent description
    AgentVersion  string // Agent version
}
```

## 📡 API Endpoints

### Agent Card Endpoint

- **URL**: `GET /card`
- **Description**: Get basic agent information
- **Response**: JSON format agent card information

### Message Sending Endpoint

- **URL**: `POST /api/messages`
- **Description**: Send message to agent
- **Request Body**: JSON containing task ID, user role, and message content

## 🎯 Usage Examples

### 1. Get Agent Information

```bash
curl http://localhost:8080/card
```

### 2. Send Message

```bash
curl -X POST http://localhost:8080/api/messages \
  -H "Content-Type: application/json" \
  -d '{
    "message": {
      "taskId": "1",
      "role": "user",
      "parts": [
        {
          "kind": "text",
          "text": "Hello, A2A!"
        }
      ]
    }
  }'
```

## 🔍 How It Works

1. **Server Startup**: Server creates an in-memory task store, print executor, and queue manager
2. **Client Connection**: Client connects to server and gets agent card information
3. **Message Sending**: Client sends message to server
4. **Task Execution**: Server receives message, creates task and executes it
5. **Result Return**: Execution results are returned to client through queue

## 🛠️ Using Makefile

The project provides a convenient Makefile to simplify common operations:

```bash
# View all available commands
make help

# Build all applications
make build

# Run server
make run-server

# Run client
make run-client

# Run tests
make test

# Clean build artifacts
make clean

# Install dependencies
make deps

# Format code
make fmt

# Code linting
make lint

# Run automatic demo
make demo-auto
```

## 🐛 Troubleshooting

### Common Issues

1. **Port Already in Use**
   - Error: `bind: address already in use`
   - Solution: Modify environment variable `A2A_SERVER_PORT` or use `make clean`

2. **Connection Refused**
   - Error: `connection refused`
   - Solution: Ensure server is running, check `A2A_SERVER_URL` configuration

3. **Module Dependency Issues**
   - Error: `go: module not found`
   - Solution: Run `make deps` or `go mod tidy` to update dependencies

4. **Permission Issues**
   - Error: `permission denied`
   - Solution: Ensure sufficient permissions to run the program

### Debug Mode

Enable detailed logging:

```go
log.SetFlags(log.LstdFlags | log.Lshortfile)
```

### Environment Variable Debugging

Check if environment variables are loaded correctly:

```bash
# Client
cd basic/client && go run . -debug

# Server
cd basic/server && go run . -debug
```

## 🤝 Contributing

Welcome to submit Issues and Pull Requests to improve this demo project!

## 📄 License

This project is licensed under the Apache2.0 License. See the [LICENSE](LICENSE) file for details.

## 🔗 Related Links

- [A2A Go SDK](https://github.com/yeeaiclub/a2a-go)
- [Go Official Documentation](https://golang.org/doc/)
- [HTTP Client Documentation](https://golang.org/pkg/net/http/) 