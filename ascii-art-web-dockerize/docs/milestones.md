# Project Milestones

## Milestone 1: Foundation
**Focus:** Project Setup & Shared Contracts
- [x] Task 00: Core Infrastructure
    - Project initialized.
    - Directory structure created.
    - Shared `Banner` model defined.

## Milestone 2: Data & Input Handling
**Focus:** Parsing & Validation
- [x] Task 01: Banner Management (Data Layer)
    - Banner loader implemented.
    - Unit tests for file reading and parsing passed.
- [x] Task 02: Input Processing (Input Layer)
    - CLI argument parser implemented.
    - Escape sequence handling (`\n`) working.
    - ASCII validation working.

## Milestone 3: Core Logic
**Focus:** Rendering Engine
- [x] Task 03: Rendering Engine (Logic Layer)
    - Single line rendering working.
    - Multi-line rendering working.
    - Edge cases (empty strings) handled.

## Milestone 4: Delivery
**Focus:** Integration & Quality Assurance
- [x] Task 04: Integration & Golden Tests
    - Application wired in `main.go`.
    - Golden files generated and verified.
    - Integration test suite passing.
    - Final `go test ./...` check passed.

## Milestone 5: Extended Features (Color, Output, Align)
**Focus:** Advanced Functionality
- [x] Task 05: Advanced Input Parsing
    - Implement flag parsing (`--color`, `--output`, `--align`).
    - Update `Config` model.
- [x] Task 06: Enhanced Rendering
    - Implement ANSI color application.
    - Implement text alignment (terminal size detection).
- [x] Task 07: File Output & Integration
    - Implement file writing in `main`.
    - Update integration tests for new features.

## Milestone 6: Web, API, and Delivery Readiness
**Focus:** Browser Experience, Containerization, and Final Audit Flow
- [x] Task 10: Web Interface
    - Browser UI implemented.
- [x] Task 11: Web Handlers
    - Form submission and HTTP routing implemented.
- [x] Task 12: Web Templates
    - Template rendering and static assets integrated.
- [x] Task 13: API
    - JSON API endpoint implemented and tested.
- [x] Task 14: Export
    - Downloadable text export implemented.
- [x] Task 15: Dockerize
    - Multi-stage Dockerfile implemented and verified.
- [x] Task 16: Final Audit Readiness
    - Docker audit flow, helper script, runtime layout, and docs aligned for submission.
