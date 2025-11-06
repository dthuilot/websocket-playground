# Contributing to WebSocket Playground

Thank you for considering contributing to this project! Here are some guidelines to help you get started.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/websocket-playground.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes
6. Commit your changes: `git commit -am 'Add new feature'`
7. Push to the branch: `git push origin feature/your-feature-name`
8. Create a Pull Request

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Make (optional, but recommended)

### Installation

```bash
# Clone the repository
git clone https://github.com/dthuilot/websocket-playground.git
cd websocket-playground

# Install dependencies
go mod download

# Run tests
go test ./...

# Run the server
make server
```

## Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Write meaningful commit messages
- Add tests for new features

## Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting a PR
- Test WebSocket connections manually when making changes to the handler

## Pull Request Process

1. Update the README.md with details of changes if applicable
2. Update documentation for any new features
3. Ensure all tests pass
4. Request review from maintainers

## Reporting Bugs

When reporting bugs, please include:

- Go version
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Any relevant logs or error messages

## Feature Requests

We welcome feature requests! Please:

- Check if the feature has already been requested
- Provide a clear description of the feature
- Explain why it would be useful
- Consider submitting a PR if you can implement it

## Code of Conduct

Be respectful and constructive in all interactions.

## Questions?

Feel free to open an issue for any questions or concerns.
