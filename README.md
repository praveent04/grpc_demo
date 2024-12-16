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

Here's a detailed explanation of **gRPC project** workflow, code structure, and critical points.

---

## **Code Flow Explanation**

Your project demonstrates the usage of all four types of **gRPC communication**:
1. **Unary RPC**  
2. **Server Streaming RPC**  
3. **Client Streaming RPC**  
4. **Bidirectional Streaming RPC**  

It’s well-structured with clear separation between client-side, server-side, and the proto definitions.

---

### **1. Proto File (`greet.proto`)**
- This is the core definition for the gRPC service `GreetService` and its methods. It defines:
   - **Service Methods**:
     - `SayHello` → Unary RPC
     - `SayHelloServerStreaming` → Server Streaming
     - `SayHelloClientStreaming` → Client Streaming
     - `SayHelloBidirectionalStreaming` → Bidirectional Streaming
   - **Messages**:
     - `NoParam` → Empty request message.
     - `HelloRequest` → Takes a single string field, `name`.
     - `HelloResponse` → Server response containing `message`.
     - `NamesList` → Accepts a list of names (string array).
     - `MessagesList` → Response with a list of messages (string array).

The `greet.proto` file acts as a **contract** for the communication between the client and the server. The `protoc` compiler generates the Go files (`greet.pb.go` and `greet_grpc.pb.go`) from this definition, which both client and server import.

---

### **2. Client-Side Workflow**
The client calls different RPC methods defined in `greet.proto` through the `GreetServiceClient` interface. The `main.go` file orchestrates the client workflow.

#### **Files and Responsibilities**:
- **`main.go`**:  
   - Establishes a connection to the server (`grpc.Dial`) using an insecure transport.  
   - Creates a gRPC client (`GreetServiceClient`).  
   - Calls all four RPC methods in sequence, each defined in separate files:
      - `unary.go` → `callSayHello()`
      - `server_stream.go` → `callSayHelloServerStream()`
      - `client_stream.go` → `callSayHelloClientStream()`
      - `bi_stream.go` → `callSayHelloBidirectionalStream()`

- **`unary.go`** (Unary RPC):  
   - Sends a single request (`NoParam`) and receives a single response (`HelloResponse`).  
   - Timeout is set using `context.WithTimeout`.

- **`server_stream.go`** (Server Streaming RPC):  
   - Sends a request (`NamesList`) to the server.  
   - Receives a **stream of responses** (`HelloResponse`) from the server.  
   - Uses a `for` loop to **read messages** from the stream until `io.EOF`.

- **`client_stream.go`** (Client Streaming RPC):  
   - Opens a client-side stream.  
   - Sends multiple `HelloRequest` messages sequentially (with a `Sleep` of 2 seconds between sends).  
   - After sending all messages, it calls `CloseAndRecv()` to receive a final `MessagesList` response from the server.

- **`bi_stream.go`** (Bidirectional Streaming RPC):  
   - Opens a bidirectional stream where **client and server exchange messages concurrently**.  
   - Starts a goroutine to **listen** for server responses (`stream.Recv`).  
   - Sends `HelloRequest` messages in a loop.  
   - Closes the stream using `stream.CloseSend()` and waits for the goroutine to finish.

---

### **3. Server-Side Workflow**
The server implements the `GreetService` interface and provides concrete definitions for the RPC methods.

#### **Files and Responsibilities**:
- **`main.go`**:  
   - Creates a gRPC server instance.  
   - Registers the `helloServer` service.  
   - Listens for connections on port `8080`.

- **`unary.go`** (Unary RPC):  
   - Implements the `SayHello` method.  
   - Returns a `HelloResponse` message with a simple greeting.

- **`server_stream.go`** (Server Streaming RPC):  
   - Implements the `SayHelloServerStreaming` method.  
   - Loops through the names provided in `NamesList`.  
   - Sends a `HelloResponse` for each name with a 2-second delay.

- **`client_stream.go`** (Client Streaming RPC):  
   - Implements the `SayHelloClientStreaming` method.  
   - Reads multiple `HelloRequest` messages from the stream.  
   - Appends each name to a `messages` array.  
   - Sends a single `MessagesList` response once the client finishes sending data.

- **`bi_stream.go`** (Bidirectional Streaming RPC):  
   - Implements the `SayHelloBidirectionalStreaming` method.  
   - Reads a stream of `HelloRequest` messages.  
   - For each request, sends a `HelloResponse` message back to the client.

---

## **Major Important Things**
1. **Proto File**: Defines the structure and contract of communication.
2. **gRPC Types**:
   - **Unary RPC** → Single request-response.
   - **Server Streaming RPC** → Server sends multiple responses.
   - **Client Streaming RPC** → Client sends multiple requests, server responds once.
   - **Bidirectional Streaming RPC** → Both send streams concurrently.
3. **Stream Management**:
   - **Client Streaming** and **Bidirectional Streaming** use `stream.Send()` and `stream.Recv()`.
   - Always handle `io.EOF` to detect the end of a stream.
4. **Concurrency**:
   - `goroutines` are used for bidirectional streaming to receive messages concurrently while sending.

---
