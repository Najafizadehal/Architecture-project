
# Von Neumann Machine Simulator

This project simulates a Von Neumann machine architecture using Go and the Gin framework. It provides an API to load instructions, execute them, and interact with the memory and registers.

## Project Structure

- **main.go**: Initializes the components and starts the Gin server.
- **controllers**: Contains the controller logic to handle API requests.
- **controlunit**: Implements the control unit responsible for fetching, decoding, and executing instructions.
- **memory**: Handles the memory and registers.
- **alu**: Arithmetic Logic Unit (ALU) for performing arithmetic and logic operations.
- **bus**: Implements the data bus for data transfer between memory, registers, and ALU.

## Getting Started

### Prerequisites

- Go 1.16+
- Gin framework

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/Najafizadehal/Architecture-project.git
    cd Architecture-project
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

### Running the Server

Start the Gin server:
```bash
go run main.go
```

The server will start on port `8080`.

### API Endpoints

#### Load Instructions

Load multiple instructions into memory.

```bash
curl -X POST "http://localhost:8080/load_instructions" -H "Content-Type: application/json" -d '[
    {"address": "0100", "value": "0005"},  # First number
    {"address": "0101", "value": "0003"},  # Second number
    {"address": "0102", "value": "0000"},  # Result storage
    {"address": "0103", "value": "3100"},  # LDA 0100
    {"address": "0104", "value": "1101"},  # ADD 0101
    {"address": "0105", "value": "4102"}   # STA 0102
]'
```

#### Write to Register

Set a specific register to a given value.

```bash
curl -X POST "http://localhost:8080/write/register" -H "Content-Type: application/json" -d '{"register": "PC", "value": 259}'
```

#### Run a Cycle

Execute a full fetch-decode-execute cycle.

```bash
curl -X GET "http://localhost:8080/run_cycle"
```

#### Read Memory

Read the value stored at a specific memory address.

```bash
curl -X GET "http://localhost:8080/memory/read/0102"
```

#### Fetch Stage

Perform the fetch stage of the instruction cycle.

```bash
curl -X GET "http://localhost:8080/fetch"
```

#### Decode Stage

Perform the decode stage of the instruction cycle.

```bash
curl -X GET "http://localhost:8080/decode"
```

#### Execute Stage

Perform the execute stage of the instruction cycle.

```bash
curl -X GET "http://localhost:8080/execute?opcode=1&address=0100"
```

### Example Workflow

1. **Load Instructions and Data**:

    ```bash
    curl -X POST "http://localhost:8080/load_instructions" -H "Content-Type: application/json" -d '[
        {"address": "0100", "value": "0005"},  # First number
        {"address": "0101", "value": "0003"},  # Second number
        {"address": "0102", "value": "0000"},  # Result storage
        {"address": "0103", "value": "3100"},  # LDA 0100
        {"address": "0104", "value": "1101"},  # ADD 0101
        {"address": "0105", "value": "4102"}   # STA 0102
    ]'
    ```

2. **Set Program Counter (PC)**:

    ```bash
    curl -X POST "http://localhost:8080/write/register" -H "Content-Type: application/json" -d '{"register": "PC", "value": 259}'
    ```

3. **Run Instruction Cycles**:

    Fetch-Decode-Execute Cycle 1:
    ```bash
    curl -X GET "http://localhost:8080/run_cycle"
    ```

    Fetch-Decode-Execute Cycle 2:
    ```bash
    curl -X GET "http://localhost:8080/run_cycle"
    ```

    Fetch-Decode-Execute Cycle 3:
    ```bash
    curl -X GET "http://localhost:8080/run_cycle"
    ```

4. **Read Result**:

    ```bash
    curl -X GET "http://localhost:8080/memory/read/0102"
    ```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
