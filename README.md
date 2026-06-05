# 🧩 Tetris Optimizer

> A Go command-line tool that reads a set of tetrominoes from a text file and assembles them into the **smallest possible square grid** using a backtracking algorithm.

---

## 📦 Project Structure

```
tetris-optimizer/
├── cmd/
│   └── tetris-optimizer/
│       ├── main.go               # Entry point — CLI argument handling & orchestration
│       └── main_test.go          # Integration tests
├── internal/
│   ├── parser/
│   │   ├── parser.go             # File reading, parsing & validation
│   │   └── parser_test.go        # Unit tests
│   ├── solver/
│   │   ├── solver.go             # Backtracking algorithm
│   │   └── solver_test.go        # Unit tests
│   └── printer/
│       ├── printer.go            # Grid renderer (stdout)
│       └── printer_test.go       # Unit tests
├── testdata/                     # Input fixtures for tests
├── docs/
│   ├── PRD.md                    # Product Requirements Document
│   ├── Architecture.md           # Architecture & TDD test scenarios
│   └── tasks/                    # Per-feature task files
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or later
- No third-party dependencies — standard library only

### Build

```bash
go build -o tetris-optimizer ./cmd/tetris-optimizer/
```

### Run

```bash
./tetris-optimizer path/to/input.txt
```

---

## 📄 Input Format

The input file contains one or more tetrominoes, each defined as a `4×4` grid:

- `.` — empty cell
- `#` — solid block
- Tetrominoes are separated by **exactly one blank line**
- Each tetromino must have **exactly 4 contiguous `#` blocks**
- Maximum of **26 tetrominoes** (mapped to letters `A–Z`)

**Example input (`input.txt`):**

```
####
....
....
....

#...
#...
#...
#...

##..
##..
....
....
```

---

## 📤 Output

Each tetromino is assigned an uppercase letter (`A`, `B`, `C`, …) in input order. The solved square is printed to stdout, with empty cells shown as `.`:

```
AAAA
BBCC
BBCC
....
```

---

## ❌ Error Handling

The program prints `ERROR` (followed by a newline) and exits with code `1` under any of the following conditions:

| Condition |
|---|
| Wrong number of arguments (not exactly 1) |
| File does not exist or cannot be read |
| File contains 0 tetrominoes |
| Invalid characters (anything other than `.`, `#`, newlines) |
| Grid dimensions are not `4×4` |
| A grid has more or fewer than 4 `#` blocks |
| Blocks in a grid are not contiguous |
| Incorrect spacing (double blank lines, trailing blank lines) |
| More than 26 tetrominoes |

---

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with race detector
go test -race ./...

# Static analysis
go vet ./...

# All checks together
go vet ./... && go test -race ./...
```

**Test coverage:**

| Package | Tests |
|---|---|
| `internal/parser` | 16 unit tests |
| `internal/solver` | 13 unit tests |
| `internal/printer` | 5 unit tests |
| `cmd/tetris-optimizer` | 6 integration tests |

---

## ⚙️ How It Works

1. **Parse** — reads and validates the input file; returns an ordered `[]Tetromino`
2. **Solve** — computes the minimum square size `⌈√(N × 4)⌉`, then uses backtracking to place all pieces; grows the grid by 1 if no solution is found at the current size
3. **Print** — maps piece index to letter (`1→A`, `2→B`, …) and renders the grid to stdout

---

## 📚 Documentation

- [PRD.md](docs/PRD.md) — Product Requirements Document
- [Architecture.md](docs/Architecture.md) — Architecture, data flow diagrams & TDD test scenarios
