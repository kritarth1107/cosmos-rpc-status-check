# RPC Status Checker

## Overview

The RPC Status Checker is a Go application designed to periodically check the connectivity status of a list of RPC (Remote Procedure Call) endpoints. It utilizes the Cosmos SDK's client library to attempt connections and logs the outcomes to both a console and an SQLite database. The application is configurable via a `.env` file, allowing for easy adjustment of the RPC endpoints without modifying the source code.

## Features

- Periodic checking of RPC endpoint connectivity
- Logging of connection status to both console and SQLite database
- Configuration of RPC endpoints via `.env` file
- Use of UTC timestamps in log entries

## Requirements

- Go 1.15 or newer
- SQLite3
- External Go packages:
  - `github.com/mattn/go-sqlite3` for SQLite operations
  - `github.com/joho/godotenv` for loading environment variables

## Installation

1. **Clone the repository:**

```bash
git clone https://github.com/kritarth1107/cosmos-rpc-status-check.git
cd cosmos-rpc-status-check
```


2. **Install dependencies:**

Ensure you have Go installed on your system and then run:
```go
go mod tidy
```
This command will download and install the necessary Go packages.

## Configuration
Before running the application, configure the RPC endpoints in a `.env` file. Create a `.env` file in the root directory of the project and specify your RPC endpoints as follows:

### `.env` Demo
```javascript
ADDRESS_PREFIX=cosmos
RPC_ENDPOINTS=<RPC_URL_1>,<RPC_URL_2>,<RPC_URL_3>
```

## Running the Application
To start the RPC Status Checker, execute the following command in the terminal:

```go
go run main.go
```
This command will initiate the application, which will then periodically check each configured RPC endpoint every `3600 seconds`, logging the connection status to both the console and the SQLite database named `rpc_status.db`.

