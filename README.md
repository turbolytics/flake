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

| Field     | Description                  | Size     |
|-----------|------------------------------|----------|
| Timestamp | Time in milliseconds         | 64 bits  |
| WorkerID  | Worker or node identifier    | 48 bits  |
| Sequence  | Unique sequence number       | 16 bits  |

### Explanation:
- **Timestamp**: Represents the time in milliseconds since a custom epoch. It occupies 64 bits.
- **WorkerID**: Identifies the worker or node where the ID was generated. It uses 48 bits.
- **Sequence**: A number that ensures uniqueness within the same millisecond and node. It occupies 16 bits.
