# Contributing to gf

Thank you for your interest in contributing to `gf`!

## What We're Looking For

The most valuable contributions are **new pattern files** for the `examples/` directory. If you have a grep pattern you use regularly that others might find useful, please share it!

Bug fixes and improvements to the codebase are also welcome.

## Setting Up a Development Environment

```bash
# Clone the repository
git clone https://github.com/geeknik/gf
cd gf

# Run tests
make test

# Build locally
make build

# Install locally
make install
```

## Making Changes

1. Fork the repository
2. Create a branch for your changes: `git checkout -b feature/my-changes`
3. Make your changes and add tests if applicable
4. Run tests: `make test`
5. Run the linter: `make lint`
6. Commit your changes with a clear message
7. Push to your fork: `git push origin feature/my-changes`
8. Open a pull request

## Adding Pattern Files

Pattern files should be placed in the `examples/` directory with a descriptive name in `kebab-case.json` format.

Example pattern file format:

```json
{
  "flags": "-HnrE",
  "pattern": "your-pattern-here"
}
```

Or with multiple patterns:

```json
{
  "flags": "-HnrE",
  "patterns": [
    "pattern-one",
    "pattern-two"
  ]
}
```

## Code Style

- Follow standard Go formatting (`go fmt`)
- Run `go vet` before committing
- Add tests for new functionality
- Keep changes minimal and focused

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
