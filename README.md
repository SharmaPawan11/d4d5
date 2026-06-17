## Overview

d4d5 is a moderately optimized, headless chess move verifier written in Go. It acts as an absolute gatekeeper for game state validation, designed to process live broadcast data or power frontend interfaces via WebAssembly.

It takes as an input:
1. Current board position (FEN string)
2. Move which the player intends to make (Source and Target squares)

It returns:
- A new game state (`State`) if the move is mathematically valid and legal.
- An `error` if the move is invalid, illegal, or corrupted.

## Developer Setup

### Environment Setup

Developers are expected to run `task setup` after cloing to ensure that the hooks are installed.

Once the hooks are installed, they are updated automatically due to `post-merge` hook.

## Project and Directory Structure

The repository is structured to separate the execution layer from the core mathematical engine.

```text
d4d5/
├── .github/workflows/   # CI/CD pipelines for automated perft validation
├── cmd/                 # Application entrypoints
│   └── main.go          # The main executable launchpad
├── internal/
│   ├── core/            # The chess engine (8x8 array representation, piece logic, movegen)
│   └── tests/           # Heavy validation suites (Perft EPD files and diagnostic runners)
├── scripts/             # Shell scripts for local git hooks (pre-push, post-merge)
├── go.mod               # Go module definition
└── Taskfile.yml         # Task runner configurations
```

## Dependency Management

The project operates with zero external dependencies. The entire move generator, FEN parser, and testing framework are built strictly using the Go Standard Library.

## How to build

## How to test

The testing architecture is split into two tiers to balance developer velocity with strict mathematical verification.

1. Fast Unit Tests (Local)
   Runs instantaneously. Tests piece movement logic, check detection, and board state transitions. This is automatically run locally on git push via the pre-push hook.
    
    `go test -v ./...`

2. Automated Perft Validation Suite (CI/CD)
      A heavy, multi-minute mathematical proof that traverses millions of nodes to guarantee no illegal moves (ghost pawns, invalid castling, etc.) are possible. This is gated behind a build tag and runs automatically in GitHub Actions.
   
    `go test -tags=perft -timeout=20m -v ./internal/tests/...`


## How to run

## Observability

## CLI Tools and their usage

The project utilizes Task as a modern alternative to Makefiles for environment setup.

`task setup`: Installs local git hooks and prepares the development environment.

`task install-hooks`: Forces a sync of the shell scripts to the hidden `.git/hooks` directory.


## Database and Schema

**Not Applicable**. d4d5 is a purely in-memory, stateless mathematical engine. It calculates legality instantly and does not require a database, caching layer, or persistent storage.

## Credentials and Security

**Not Applicable**. The engine does not interact with third-party APIs, handle user authentication, or require environment variables containing secrets.