# Go Architecture Patterns

This repository demonstrates different architectural approaches in Go through consistent implementations of a simplified document processing app. Each implementation showcases the structural patterns, component organization, and dependency flows that characterize a specific architectural style.

## What This Is

- A simplified comparative demonstration of selected architectural patterns implemented in Go
- A learning resource for understanding structural differences between architectural approaches
- A practical guide to help with decision-making

## What This Is Not

- An exhaustive collection of possible architectural patterns and combinations
- A demonstration of domain modeling or information flow patterns
- A production-ready system

## App Overview
- Document uploads
- Document retrieval
- Http handler as input
- Mocked database as output
- Mocked eventbus as input/output
- App bootstrapping

## Prerequisites

- Go 1.24 or later

## How to Use

#### Setup
```bash
cd [clean|hexagonal|fcis|modular|vertical]
go mod tody
```
or
```bash
make setup
```

#### Running the Examples

Each architectural approach is implemented in its own directory. To run any specific example:
```bash
cd [clean|hexagonal|fcis|modular|vertical]
go run .
```

#### API
Available at `http://localhost:8080`. Available routes:
- `GET /health` - Health check
- `POST /documents/upload` - Upload a document
- `GET /documents/:id/` - Retrieve document by ID (always returns sample document)

### [TODO] Tests
To run tests for a specific architecture:

```bash
cd [clean|hexagonal|fcis|modular|vertical]
go test ./...
```
or
```bash
make test
```

## Patterns Overview

### Clean

Core Principles:
- Domain entities at the center
- Use cases orchestrate business operations
- Clear layer separation with dependencies pointing inward

### Hexagonal

Core Principles:
- Business logic isolated in the core
- Ports define interfaces for external communication
- Adapters implement port interfaces
- Clear distinction between "driving" and "driven" sides

### Functional Core, Imperative Shell (FCIS)

Core Principles:
- Pure functional core with domain logic
- Side effects isolated in the imperative shell
- Functional core has no dependencies on external systems

### Modular Architecture

Core Principles:
- Self-contained modules
- Internal cohesion within modules
- Clear public APIs between modules
- Shared kernel for common concepts

### Vertical Slice Architecture

Core Principles:
- Organized by feature rather than technical layer
- Each slice contains UI, business logic, and data access
- Minimal cross-slice dependencies
- Shared infrastructure for cross-cutting concerns

## Comparison

### Clean vs. Hexagonal

Similarities:
- Both separate domain logic from infrastructure
- Both use interfaces to define boundaries
- Both maintain unidirectional dependencies toward domain

Differences:
- Clean focuses on use case organization
- Hexagonal emphasizes the application boundary
- Clean typically has more explicit layers

### Hexagonal vs. FCIS

Similarities:
- Both isolate domain logic
- Both define clear boundaries

Differences:
- FCIS focuses on pure functions vs. side effects
- Hexagonal focuses on interface boundaries
- FCIS typically has more immutable data structures

### Modular vs. Vertical Slice

Similarities:
- Both provide organizational boundaries
- Both can contain multiple technical layers

Differences:
- Modular organizes by domain concept
- Vertical slice organizes by feature/user story
- Modular has more cross-module dependencies

## Strengths, Weaknesses, and Use Cases

### Clean

Strengths:
- Clear dependency rules
- Business logic isolation
- Framework independence
- Highly testable

Weaknesses:
- "Interface explosion"
- Potentially higher initial development time
- Overhead for simple applications

Use Cases:
- Complex business domains
- Long-lived enterprise applications
- Systems requiring high testability

### Hexagonal

Strengths:
- Clear separation of concerns
- Easily swappable external dependencies
- Good for systems with multiple interfaces

Weaknesses:
- Conceptually more complex
- Overengineering for simple applications
- Requires discipline to maintain boundaries

Use Cases:
- Applications with multiple UIs or data sources
- Systems requiring high adaptability to external changes

### Functional Core, Imperative Shell

Strengths:
- Highly testable domain logic
- Clear separation of side effects
- Simplified reasoning about business rules

Weaknesses:
- Requires functional programming discipline
- May feel unfamiliar to OOP-oriented teams
- Potential data copying overhead

Use Cases:
- Systems with complex business rules
- Applications benefiting from immutability
- Projects requiring extensive unit testing

### Modular Architecture

Strengths:
- Good for large codebases
- Enables parallel development
- Clear ownership boundaries

Weaknesses:
- Module boundaries need careful design
- Risk of duplication across modules
- Integration challenges

Use Cases:
- Larger applications with distinct functional areas
- Systems developed by multiple teams
- Applications where parts evolve at different rates

### Vertical Slice Architecture

Strengths:
- Feature isolation
- Rapid feature development
- Reduces cognitive load for feature work

Weaknesses:
- Potential code duplication
- Inconsistent patterns
- Requires discipline for shared concerns

Use Cases:
- Prioritizing feature velocity
- Diverse and rapidly changing feature requirements
- Teams organized around product features
