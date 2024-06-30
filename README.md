# flake

flake is a Go-based project that generates 128-bit globally unique identifiers (GUIDs) using a system inspired by Twitter's Snowflake IDs. The project offers both an HTTP server and a client library for flexible integration into various applications.
Features

- 128-bit GUID Generation: Generates unique 128-bit IDs with time-based ordering.
- HTTP Server: Provides RESTful endpoints for GUID generation.
- Client Library: Simple Go library for embedding GUID generation in your own applications.
- Time-Ordered IDs: Ensures IDs are roughly ordered by generation time.
- Scalable: Suitable for distributed systems with high ID generation rates.


```
|-----------------------+-------------+-------------+----------------+-------------+
| 48 bits (Timestamp)   | 16 bits (Region ID) | 16 bits (Machine ID) | 32 bits (Sequence) | 16 bits (Randomness) |
+-----------------------+-------------+-------------+----------------+-------------+
```


## 128-bit Flake ID Structure

The 128-bit Flake ID is composed of the following components:

| Bits | Field Name   | Description                                                       |
|------|--------------|-------------------------------------------------------------------|
| 48   | Timestamp    | Milliseconds since a custom epoch. Provides ample range for time representation. |
| 16   | Region ID    | Identifies the region or data center.                             |
| 16   | Machine ID   | Identifies the specific machine or node.                          |
| 32   | Sequence     | Allows for a high number of unique IDs per millisecond per machine. |
| 16   | Randomness   | Adds additional uniqueness and helps mitigate potential collisions. |
| **128** | **Total**   |                                                                 |


