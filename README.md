# gRPC PubSub Broker (Go)

This project implements a lightweight **Publish-Subscribe** system using **pure Go and gRPC** — with no external brokers!

## Features
- Subscribe to topics
- Publish messages
- Target individual subscribers
- Broadcast to a selected group
- List all subscribers
- List all topics
- Unsubscribe subscribers
- Capture subscriber IP addresses
- Automatic reconnection with exponential backoff
- Docker Compose setup for multi-server deployment

## Quick Start

### 1. Build Protobuf Files
```bash
make proto-build

