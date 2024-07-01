# flake

flake is a Go-based project that generates 128-bit globally unique identifiers (GUIDs) using a system inspired by Twitter's Snowflake IDs. The project offers both an HTTP server and a client library for flexible integration into various applications.
Features

- 128-bit GUID Generation: Generates unique 128-bit IDs with time-based ordering.
- HTTP Server: Provides RESTful endpoints for GUID generation.
- Client Library: Simple Go library for embedding GUID generation in your own applications.
- Time-Ordered IDs: Ensures IDs are roughly ordered by generation time.
- Scalable: Suitable for distributed systems with high ID generation rates.


```
+----------------------+----------------------+----------------------+
|      Timestamp       |      WorkerID        |      Sequence        |
|      (64 bits)       |      (48 bits)       |      (16 bits)       |
+----------------------+----------------------+----------------------+
```


## 128-bit Flake ID Structure

| Field     | Description                         | Size     |
|-----------|-------------------------------------|----------|
| Timestamp | Time in milliseconds                | 64 bits  |
| WorkerID  | Worker, node, or process identifier | 48 bits  |
| Sequence  | Unique sequence number              | 16 bits  |

### Explanation:
- **Timestamp**: Represents the time in milliseconds since a custom epoch. It occupies 64 bits.
- **WorkerID**: Identifies the worker or node where the ID was generated. It uses 48 bits.
- **Sequence**: A number that ensures uniqueness within the same millisecond and node. It occupies 16 bits.

## Contents

- [Flake](#flake)
- [Getting Started](#Getting Started)
  - [Usage](#Usage)
    - [Generating a Flake ID](#Generating a Flake ID)
    - [Customizing the Worker ID](#Customizing the Worker ID)
    - [Parsing a Flake ID](#Parsing a Flake ID)
  - [Integrating with Your Application](#Integrating with Your Application)
  - [Usage Example Webserver Command](#Usage Example Webserver Command)
- [ID Overview](#ID Overview)


## Getting Started

Welcome to the Flake ID Generator! This guide will help you get started with installing, configuring, and using the Flake ID generator and CLI tool.

### Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Go](https://golang.org/dl/) (version 1.21 or later)
- [Git](https://git-scm.com/)

### Installation

#### 1. Clone the Repository

First, clone the `flake` repository to your local machine:

```sh
git clone https://github.com/turbolytics/flake.git
cd flake
```

#### 2. Build the CLI

Compile the `flake` CLI tool by running:

```sh
make build
```

The CLI binary will be available in the `./bin` directory.

### Usage

The `flake` CLI provides commands to generate and parse Flake IDs. Below are some common operations you can perform.

#### Generating a Flake ID

Generate a Flake ID using the default configuration:

```sh
./bin/flake generate
```

Example output:

```
id="000001906F401F94-000000000001-0000" timestamp="2024-07-01 13:03:11.764 -0400 EDT"
```

#### Customizing the Worker ID

Specify a worker ID when generating a Flake ID:

```sh
./bin/flake generate --worker 12345
```

Example output:

```
Generated Flake ID: 017D2E31244A4F1F-000000003039-0001
```

#### Parsing a Flake ID

Parse an existing Flake ID to view its components:

```sh
./bin/flake parse 000001906F401F94-000000000001-0000 
```

Example output:

```
Timestamp: 1625074800000
WorkerID: 0x1ABCD
Sequence: 0x0001
```

### Integrating with Your Application

The `flake` library can also be used in Go applications to generate and parse Flake IDs programmatically. Here's a quick guide on integrating it into your application.

#### Installation

Add the `flake` library to your Go project:

```sh
go get github.com/turbolytics/flake
```

#### Example Usage

Use the `flake` package to generate a new Flake ID:

```go
package main

import (
	"fmt"
	"github.com/turbolytics/flake/pkg/flake"
)

func main() {
	// Create a new generator with a specific worker ID
	gen, err := flake.NewGenerator(flake.GeneratorWithWorkerID(12345))
	if err != nil {
		panic(err)
	}

	// Generate a new Flake ID
	id, err := gen.GenerateFlakeID()
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated Flake ID:", id.String())
}
```

To create a usage example for the `webserver` command from the `http.go` file in the [turbolytics/flake](https://github.com/turbolytics/flake) repository, we'll include the following steps in the `README.md`:

1. **Starting the Webserver**: How to start the webserver using the CLI.
2. **Interacting with the Webserver**: Using `curl` to interact with the webserver.
3. **Expected Output**: What to expect as output from the webserver.

Here's how you can structure the README section:

---

## Usage Example: Webserver Command

The `http` command starts an HTTP server that generates Flake IDs. Below is a guide on how to start the server and interact with it using `curl`.

### Starting the Webserver

To start the webserver, use the `http` command. You can specify the port using the `--port` flag. The default port is `8080`.

```sh
./bin/flake http --port 8080
```

**Example Output:**

```sh
2024/07/01 13:11:48 Handlers listening on port 8080...
```

### Interacting with the Webserver

Once the server is running, you can interact with it to generate Flake IDs. The server provides an endpoint to generate a new Flake ID.

#### Generate a Flake ID

You can use `curl` to make an HTTP GET request to the `/generate` endpoint.

```sh
curl -X GET http://localhost:8080/generate
```

**Example Output:**

```json
{
  "id": "000001906F484A3E-000000000001-0000",
  "flake": {
    "timestamp": 1719853926974,
    "worker_id": 1,
    "sequence": 0
  }
}
```

### Details of the Response

- `id`: The string representation of the Flake ID.
- `flake`: An object containing:
    - `timestamp`: The timestamp component of the Flake ID.
    - `worker_id`: The worker ID component.
    - `sequence`: The sequence number.

This simple interaction demonstrates how you can generate unique Flake IDs via HTTP requests to the running webserver.

## ID Overview

### Flake IDs: Collision-Free, Coordination-Free Unique Identifiers

In distributed systems, generating unique identifiers is crucial for ensuring data consistency and avoiding conflicts. Traditional database systems often rely on sequential primary keys that necessitate coordination to maintain uniqueness. In contrast, Flake IDs and UUIDs provide coordination-free unique identifiers, each with its own approach and advantages.

#### Flake IDs

Flake IDs are designed to be unique without requiring any centralized coordination. They achieve this by incorporating a combination of:

- **Timestamp**: Provides the time-based component, ensuring chronological ordering.
- **Worker ID**: Uniquely identifies the machine or process generating the ID, ensuring that IDs generated simultaneously on different machines are unique.
- **Sequence**: A counter that increments with each ID generated, ensuring uniqueness even when multiple IDs are created within the same millisecond.

##### How Flake IDs Ensure Uniqueness

1. **Timestamp**: The first part of a Flake ID is derived from the current timestamp, allowing IDs to be sorted by the order of creation.
2. **Worker ID**: Each generator instance is assigned a unique worker ID, which ensures that IDs produced on different machines or processes do not collide.
3. **Sequence**: Within the same millisecond, a sequence counter is used to differentiate between multiple IDs generated by the same worker.

This combination guarantees that each ID is globally unique without needing any coordination between machines or processes. The design allows for high throughput, making Flake IDs particularly suitable for distributed systems and microservices architectures.

#### UUIDs

UUIDs (Universally Unique Identifiers) are another type of unique identifier that do not require coordination. They are designed to be unique across both space and time, using a variety of strategies depending on the UUID version:

- **Version 1 (Timestamp and MAC Address)**: Combines the current timestamp with the MAC address of the generating machine, ensuring uniqueness over time and across different machines.
- **Version 4 (Random)**: Uses random numbers to generate the ID, relying on the vast address space to avoid collisions.

##### How UUIDs Ensure Uniqueness

1. **Version 1**: By using the machine's MAC address and a timestamp, Version 1 UUIDs ensure that IDs generated at different times or on different machines are unique.
2. **Version 4**: The random approach of Version 4 UUIDs provides a high probability of uniqueness by leveraging a large namespace.

UUIDs are versatile and widely used, but the random nature of Version 4 UUIDs does not provide orderability based on creation time, unlike Flake IDs.

#### Comparison with Traditional Coordinated IDs

Traditional ID systems, such as those used in databases, often rely on sequential primary keys. These systems require coordination to ensure that each new ID is unique:

- **Centralized Coordination**: A central authority (e.g., the database) issues new IDs, incrementing a counter to generate the next ID.
- **Conflict Avoidance**: The central authority must manage access to the ID generator, often involving locks or other synchronization mechanisms to prevent conflicts.

While these systems ensure uniqueness, they introduce bottlenecks due to the need for coordination, which can limit scalability and increase latency.

##### Advantages of Coordination-Free IDs

1. **Scalability**: Flake IDs and UUIDs can be generated independently on multiple machines or processes without the need for synchronization, allowing for high scalability.
2. **Fault Tolerance**: Without reliance on a central authority, Flake IDs and UUIDs are more resilient to single points of failure.
3. **Performance**: By eliminating the need for coordination, these systems can generate IDs with lower latency, improving overall performance.

#### Conclusion

Flake IDs and UUIDs offer robust solutions for generating unique identifiers in distributed systems. They avoid the bottlenecks associated with traditional coordinated ID systems, providing a balance between uniqueness, performance, and scalability. Flake IDs, with their timestamp-based approach, also provide the added benefit of orderability, making them an excellent choice for systems that require both uniqueness and chronological sorting.

