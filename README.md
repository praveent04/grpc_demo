# **gRPC Demo in Go**

This project demonstrates all four types of gRPC communication in Go:

1. Unary RPC  
2. Server Streaming RPC  
3. Client Streaming RPC  
4. Bidirectional Streaming RPC  

## **Project Structure**
```plaintext
grpc-demo/
├── client/
│   ├── main.go             # Entry point for the client
│   ├── unary.go            # Unary RPC client logic
│   ├── server_stream.go    # Server streaming RPC client logic
│   ├── client_stream.go    # Client streaming RPC client logic
│   └── bi_stream.go        # Bidirectional streaming RPC client logic
├── server/
│   ├── main.go             # Entry point for the server
│   ├── unary.go            # Unary RPC server logic
│   ├── server_stream.go    # Server streaming RPC server logic
│   ├── client_stream.go    # Client streaming RPC server logic
│   └── bi_stream.go        # Bidirectional streaming RPC server logic
├── proto/
│   ├── greet.proto         # gRPC service definition
│   ├── greet.pb.go         # Generated Go code for messages
│   └── greet_grpc.pb.go    # Generated Go code for gRPC service
├── go.mod                  # Go module file
└── go.sum                  # Go dependency checksums
```

## **Setup**

1. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

2. **Run the Server**:
   ```bash
   go run server/main.go
   ```

3. **Run the Client**:
   ```bash
   go run client/main.go
   ```

## **Communication Types**

1. **Unary RPC**:
   - Sends a single request and receives a single response.

2. **Server Streaming RPC**:
   - Sends a request, and the server streams multiple responses.

3. **Client Streaming RPC**:
   - Client sends multiple requests, and the server responds once.

4. **Bidirectional Streaming RPC**:
   - Client and server send streams concurrently.

## **Generated Files**
Use `protoc` to generate Go files from `greet.proto`:
```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

---
