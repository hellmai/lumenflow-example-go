# LumenFlow Example: Go

Example Go project demonstrating [LumenFlow](https://lumenflow.dev) workflow setup.

## What is LumenFlow?

LumenFlow is a delivery framework for AI-native software development. It provides:

- **Work Units (WUs)**: Atomic units of work with clear acceptance criteria
- **Lanes**: Parallel delivery streams with WIP=1 enforcement
- **Gates**: Quality checks (format, lint, typecheck, test) that must pass
- **TDD-First**: Test-driven development with hexagonal architecture

## Quick Start

```bash
# Run all gates
make gates

# Or run individually:
make fmt-check   # Check formatting
make lint        # Run golangci-lint
make vet         # Run go vet
make test        # Run tests with coverage

# Build the binary
make build

# Run the example
./bin/calculator
```

## Project Structure

```
.
├── .lumenflow.config.yaml    # LumenFlow configuration
├── .beacon/
│   └── stamps/               # WU completion stamps
├── docs/
│   └── 04-operations/
│       └── tasks/
│           ├── wu/           # Work Unit specs
│           ├── backlog.md    # Backlog of all WUs
│           └── status.md     # Current status board
├── calculator/
│   ├── calculator.go         # Calculator package
│   └── calculator_test.go    # Tests
├── cmd/
│   └── calculator/
│       └── main.go           # Entry point
├── Makefile                  # Build and gate commands
└── go.mod                    # Go module definition
```

## Gates Configuration

This project uses the following quality gates:

| Gate | Tool | Command |
|------|------|---------|
| Format | gofmt | `make fmt-check` |
| Lint | golangci-lint | `make lint` |
| Vet | go vet | `make vet` |
| Test | go test | `make test` |

All gates must pass before a WU can be marked done.

## Working with WUs

### View the Backlog

See [docs/04-operations/tasks/backlog.md](docs/04-operations/tasks/backlog.md) for available work.

### Claim a WU

```bash
# With LumenFlow CLI installed:
pnpm wu:claim --id WU-001 --lane Core
cd worktrees/core-wu-001
```

### Complete a WU

```bash
# Run gates first
make gates

# Return to main and complete
cd ../..
pnpm wu:done --id WU-001
```

## Documentation

- [LumenFlow Documentation](https://lumenflow.dev) - Complete framework reference
- [Getting Started](https://lumenflow.dev/getting-started/quickstart) - Quick start guide
- [Configuration Reference](https://lumenflow.dev/reference/config) - Config options
- [WU Schema](https://lumenflow.dev/reference/wu-schema) - Work Unit specification

## License

MIT
