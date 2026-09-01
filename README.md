# Zone01 Athens — Project Portfolio

A collection of projects developed during my studies at [Zone01 Athens](https://zone01.org), an intensive peer-to-peer programming school. Each project explores a different domain — from networking and algorithms to web development and containerization — all built with **Go** and following clean architecture principles.

## About Me

**Panagiotis Kalligeris** — Software engineering student at Zone01 Athens, focused on building robust, well-tested systems in Go.

[![GitHub](https://img.shields.io/badge/GitHub-pkalligeris-181717?style=flat&logo=github)](https://github.com/pkalligeris)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Panagiotis%20Kalligeris-0A66C2?style=flat&logo=linkedin)](https://www.linkedin.com/in/panagiotis-kalligeris-6ab69b398)

---

## Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![TCP](https://img.shields.io/badge/TCP%2FIP-Networking-4B8BBE?style=for-the-badge)
![Leaflet](https://img.shields.io/badge/Leaflet-199900?style=for-the-badge&logo=leaflet&logoColor=white)

---

## Projects

### 1. Net-Cat — TCP Group Chat

A TCP-based group chat server and client inspired by the Unix `nc` utility. Supports concurrent connections, multiple chat rooms, direct messages, spam protection, IP banning, and an optional terminal UI built with `gocui`.

| | |
|---|---|
| **Language** | Go |
| **Concepts** | TCP sockets, concurrency (goroutines/channels), rate limiting, terminal UI |
| **Highlights** | Up to 10 concurrent clients, 256 rooms, per-room message history, nickname system, graceful shutdown |

**In-chat commands:** `/nick`, `/switch`, `/dm`, `/rooms`, `/users`, `/history`, `/stats`, `/leave`

[View Project →](./Net-cat)

---

### 2. ASCII Art Web — Dockerize

A feature-rich ASCII art generator that converts text into graphical ASCII art using banner files. Available as both a CLI tool and a web application with live preview, served from a production-ready Docker container.

| | |
|---|---|
| **Language** | Go, HTML/CSS, JavaScript |
| **Concepts** | HTTP server, HTML templating, REST API, multi-stage Docker builds, AJAX live preview |
| **Highlights** | 3 banner fonts, text colorization (Hex/RGB/HSL), alignment options, file export, containerized with non-root runtime |

**CLI features:** `--color`, `--output`, `--align` flags with multiple banner support (`standard`, `shadow`, `thinkertoy`)

[View Project →](./ascii-art-web-dockerize)

---

### 3. Groupie Tracker — Visualizations

A web application that consumes a RESTful API to display information about musical artists and bands, featuring interactive map visualizations, live search, and a multi-criteria filtering system.

| | |
|---|---|
| **Language** | Go, HTML/CSS, JavaScript |
| **Concepts** | REST API consumption, geocoding (OpenStreetMap Nominatim), interactive maps (Leaflet), client-side filtering |
| **Highlights** | Chronological tour path animation, real-time search with autocomplete, multi-criteria filters (date, members, location) |

**Filter criteria:** creation date, first album year, number of members, concert locations

[View Project →](./groupie-tracker-visualizations)

---

### 4. Tetris Optimizer

A command-line tool that reads a set of tetrominoes from a text file and assembles them into the smallest possible square grid using a backtracking algorithm. Supports up to 26 pieces (A–Z) with comprehensive input validation.

| | |
|---|---|
| **Language** | Go |
| **Concepts** | Backtracking algorithm, recursive search, input parsing and validation, golden file testing |
| **Highlights** | Optimal square computation, 40 unit + integration tests, strict validation (contiguity, dimensions, character set) |

**How it works:** Parse → Solve (backtracking with grid growth) → Print (letter-mapped grid)

[View Project →](./tetris-optimizer)

---

## Running Any Project

All projects use Go modules and can be run with:

```bash
cd <project-directory>
go run .
```

For the Dockerized project:

```bash
cd ascii-art-web-dockerize
docker build -t ascii-art-web .
docker run --rm -p 8080:8080 ascii-art-web
```

To run tests across any project:

```bash
cd <project-directory>
go test ./... -v
```

---

## License

These projects were completed as part of the [Zone01 Athens](https://zone01.org) curriculum.
