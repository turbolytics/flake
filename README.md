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


# Usage 

## Generate Flake ID from CLI

Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.Certainly! Below is a Markdown-formatted README with usage examples based on the changes in the pull request for the Flake project:

---

# Flake

Flake is a Go library and CLI tool for generating 128-bit globally unique identifiers (GUIDs) based on Twitter's Snowflake IDs.

## Installation

To install Flake, use `go get`:

```bash
go get github.com/turbolytics/flake
```

## Usage

### CLI Usage

Flake CLI provides commands to generate Flake IDs and manage configurations.

#### Generate Command

Generate one or more Flake IDs using the CLI.

```bash
flake generate --count 5
```

Options:
- `--count`: Number of IDs to generate (default is 1)

Example output:
```
Generated Flake ID: 000001AD5EF69192-0001-0001-00000000-0A37
Generated Flake ID: 000001AD5EF69193-0001-0001-00000000-0A38
Generated Flake ID: 000001AD5EF69194-0001-0001-00000000-0A39
Generated Flake ID: 000001AD5EF69195-0001-0001-00000000-0A3A
Generated Flake ID: 000001AD5EF69196-0001-0001-00000000-0A3B
```

### Library Usage

You can also use Flake as a library in your Go projects.

```go
package main

import (
	"fmt"
	"log"

	"github.com/turbolytics/flake"
)

func main() {
	// Initialize Flake Generator
	fg := flake.NewFlakeGenerator(1, 1) // Replace with your region and machine IDs

	// Generate a Flake ID
	id := fg.GenerateFlakeID()
	fmt.Println("Generated Flake ID:", id.String())
}
```

## Configuration

Flake can be configured using environment variables:

- `FLAKE_REGION_ID`: Region ID (default: 1)
- `FLAKE_MACHINE_ID`: Machine ID (default: 1)
- `FLAKE_LOG_ENABLED`: Enable logging (default: false)

Example:
```bash
export FLAKE_REGION_ID=2
export FLAKE_MACHINE_ID=3
export FLAKE_LOG_ENABLED=true
```

## Contributing

Contributions are welcome! Please fork the repository and submit pull requests to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

This Markdown README provides a structured overview of Flake, including CLI usage examples, library integration, configuration options, and guidance for contributing. Adjust the commands, examples, and details as per the specifics of your project and the changes made in your pull request.
