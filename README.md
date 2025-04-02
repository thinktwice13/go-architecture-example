# go-architecture-example

This repo demonstrates different architectural approaches to building the same application: a document processing app. The focus is on code organisation and component interaction, showcasing the trade-offs between architectures. Included:
- Clean Architecture
- Functional Core, Imperative Shell
- Hexagonal Architecture
- Modular
- Vertical Slices

## Overview

Each implementation delivers identical functionality:
- [x] Uploading documents
- [ ] Retrieving document details

## Architectural Implementations

### Clean Architecture

#### Key Characteristics
- Clear separation of concerns with concentric layers
- Usecases orchestrate domain operations
- Adapters tie use cases to infrastructure
- Strong encapsulation of business rules

#### When to Use
- For complex business domains with rich behavior
- For systems expected to evolve over many years
- When maximizing testability of business logic is critical
- For teams with clear separation between domain experts and infrastructure specialists

### FCIS (Functional Core, Imperative Shell)

#### Key Characteristics
- Core logic is functional and pure
- Imperative shell handles side effects
- Separation of concerns between core and shell
- Core is testable and reusable
- Shell is responsible for I/O and state management
- Encourages immutability and functional programming principles
- Dependencies flow outward from the core

#### When to Use
- For systems with complex business logic that can be expressed functionally
- When the team is comfortable with functional programming paradigms
- When the system has many side effects (I/O, state changes)
- For applications that require high testability and reusability of core logic
- When the system is expected to evolve over time with changing requirements

### Hexagonal Architecture (Ports & Adapters)

#### Key Characteristics
- Application core defines ports (interfaces) for what it needs
- Adapters implement these interfaces on both sides
- Clear distinction between "driving" (input) and "driven" (output) sides
- Core has no knowledge of specific technologies
- Emphasis on pluggability and interchangeability of components

#### When to Use
- When the system must support multiple interfaces (UI, API, CLI)
- When external dependencies might change over time
- For systems requiring high testability through substitution
- When building systems that will integrate with many external services
- For teams focused on interface-driven development

### Modular

### Vertical Slice Architecture

#### Key Characteristics
- Organized around business capabilities rather than technical layers
- Each feature contains all technical aspects (API, business logic, data access)
- Minimal cross-feature dependencies
- Optimized for feature development and maintenance
- Often uses mediator pattern for cross-feature coordination

#### When to Use
- For systems with distinct, loosely-coupled features
- When teams are organized around business capabilities
- For rapid feature development and deployment
- When different features have different technical requirements or complexities
- For teams with full-stack developers responsible for entire features

## Prerequisites
- Go 1.24 or later
- Make (optional, for convenience)

### Building and Running

To run implementations change to the respective directory and run `go run .`:

The servers is on http://localhost:8080. All implementations expose identical API endpoints:

- `POST /documents` - Upload a document
- `GET /documents/:id` - Get document details

## [TODO] Testing

Each implementation includes examples of:
- Unit test
- Integration test
- E2E test

Run tests with:

```bash
make test-all
```