# Your first GitHub Action

A sample repo that you can copy with the template button. Then you can use the GitHub Actions UI to add basic build, test, scan and docker workflows to.

The app is inspired by fortune cookies or the magic 8-ball in that it gives you back a random common code error message as your "fortune". It may or may not have abilities to predict your next error. Built with Go and Cobra, this project serves as a repo template for adding common GitHub Actions to, and demonstrates basic CLI application development.

## Project Description

This application provides random error messages from software over the years. Run it to receive your error fortune!

The project is built using:
- **Go** - The main programming language
- **Cobra** - A powerful CLI library for Go

## Installation

### Prerequisites
- Go 1.24.4 or later
- Git

### Install from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/bretfisher/first-action.git
   cd first-action
   ```

2. Build the application:
   ```bash
   go build -o eball
   ```

3. (Optional) Install globally:
   ```bash
   go install
   ```

### Install using Go Install

```bash
go install github.com/bretfisher/first-action@latest
```

## Usage Examples

### Basic Usage

Run the application to get a random fortune:

```bash
./eball
```

Output example:
```
🥠 Your Fortune:
   It was always DNS.
```

### Command Line Options

View help information:
```bash
./eball --help
```

## Development

### Project Structure

```
first-action/
├── cmd/
│   └── root.go          # Main command logic and fortune definitions
├── main.go              # Application entry point
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── LICENSE              # Project license
└── README.md            # This file
```

