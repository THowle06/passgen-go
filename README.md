# Go Password Generator

A simple and secure command-line password generator written in Go.

## Features

- Crytopgraphically secure password generation (`crypto/rand`)
- Configurable length and count
- Optional execution of symbols
- Guaranteed inclusion of character classes
- Entropy calculation with strength classification
- Optional clipboard support

## Installation

Clone the repository and build:

```bash
git clone https://github.com/THowle06/passgen-go.git
cd passgen-go
go build -o passgen-go
```

Run the binary:

```bash
./passgen-go
```

## Usage

### Generate a password

```bash
./passgen-go
```

### Generate multiple passwords

```bash
./passgen-go -count=5
```

### Custom length

```bash
./passgen-go -length=20
```

### Exclude symbols

```bash
./passgen-go -no-symbols
```

### Require all character classes

```bash
./passgen-go -require-each-class
```

### Show entropy

```bash
./passgen-go -show-entropy
```

### Copy to clipboard

```bash
./passgen-go -clipboard
```

## Example Output

```text
_aLAo80qo!55emTc
po^0w=gvoXS=O%w(

Entropy: 100.27 bits (Strong)
```

## Security Notes

- Passwords are generated using Go's `crypto/rand` for secure randomness.
- Entropy is calculated based on password length and character set size.
- Clipboard functionality depends on system envrionment support.

## Development

Run tests:

```bash
go test ./...
```
