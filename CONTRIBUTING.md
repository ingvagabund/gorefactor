# Contributing to gorefactor

Thank you for your interest in contributing to gorefactor! This document provides guidelines and instructions for contributing.

## Development Setup

1. **Fork and clone the repository:**
   ```bash
   git clone https://github.com/your-username/gorefactor.git
   cd gorefactor
   ```

2. **Ensure you have Go 1.24 or later:**
   ```bash
   go version
   ```

3. **Build the project:**
   ```bash
   make build
   ```

4. **Run tests:**
   ```bash
   make test
   ```

## Making Changes

1. **Create a branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes** following Go best practices:
   - Use `gofmt` or `goimports` for formatting
   - Write table-driven tests for new functionality
   - Add documentation comments for exported functions and types
   - Follow the existing code style

3. **Run tests and checks:**
   ```bash
   make test
   make vet
   ```

4. **Commit your changes** following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification (see below)

## Commit Messages

This project follows the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification. Every commit message must be structured as:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Types

| Type | When to use |
|------|-------------|
| `feat` | A new feature or capability |
| `fix` | A bug fix |
| `docs` | Documentation-only changes |
| `test` | Adding or updating tests |
| `refactor` | Code change that neither fixes a bug nor adds a feature |
| `ci` | Changes to CI configuration or scripts |
| `chore` | Maintenance tasks (deps, build config, etc.) |
| `perf` | Performance improvement |

### Examples

```
feat: add magic number detection analyzer
feat(cli): support scanning entire packages
fix: handle empty string literals in duplicate detection
docs: update README with JSON output examples
test: add edge case for single-character strings
refactor(duplicates): extract AST walking into helper
ci: add golangci-lint step to workflow
```

### Rules

- Use lowercase for the type and description
- Do not end the description with a period
- Use the imperative mood ("add" not "adds" or "added")
- Reference issue numbers in the footer when applicable (e.g., `Refs: #42`)
- Mark breaking changes with `!` after the type or with a `BREAKING CHANGE:` footer

## Pull Request Process

1. **Ensure your code:**
   - Passes all tests (`make test`)
   - Passes `go vet` (`make vet`)
   - Follows the project's code style
   - Includes tests for new functionality

2. **Open a Pull Request:**
   - Provide a clear description of your changes
   - Reference any related issues
   - Ensure CI checks pass

3. **Respond to feedback:**
   - Address review comments promptly
   - Make requested changes in new commits (we'll squash on merge)

## Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting
- Write clear, concise code
- Add comments for exported functions and types
- Use table-driven tests where appropriate

## Testing

- All new features should include tests
- Run `make test` before submitting a PR
- Aim for good test coverage (80%+)

## Questions?

Feel free to open an issue for questions or discussions about the project.
