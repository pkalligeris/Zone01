# 🧩 Tetris Optimizer — Product Requirements Document

> **Version:** 1.0 &nbsp;|&nbsp; **Language:** Go &nbsp;|&nbsp; **Type:** CLI Application

---

## 1. Project Overview

| Field | Detail |
|---|---|
| **Name** | Tetris Optimizer |
| **Language** | Go (Golang) |
| **Type** | Command-line application |
| **Objective** | Read a list of tetrominoes from a text file and assemble them into the **smallest possible square grid** |

---

## 2. Scope & Constraints

| Constraint | Description |
|---|---|
| **Language** | Go (Golang) |
| **Dependencies** | Strictly limited to Go standard library packages — `fmt`, `os`, `strings`, `bufio`, etc. No third-party dependencies. |
| **Quality** | Must adhere to Go best practices (Effective Go), compile without warnings, and include unit tests. |

---

## 3. Functional Requirements

### 3.1 Input Processing

**CLI Arguments**
The program must accept **exactly one argument**: the relative or absolute path to a text file.

**File Parsing**
The program must read the text file and parse the tetromino shapes.

**Format Validation**

| Rule | Specification |
|---|---|
| Grid size | Each tetromino is defined in a `4×4` grid of characters |
| Empty space | Represented by a period (`.`) |
| Solid block | Represented by a hash (`#`) |
| Separator | Tetrominoes are separated by **exactly one** empty line |
| Minimum count | The file must contain **at least one** valid tetromino |

---

### 3.2 Core Logic

**Shape Validation**
Each parsed `4×4` grid must contain exactly **4 hashes (`#`)** that are contiguous — forming a valid Tetris piece.

**Assembly Algorithm**
The program calculates the **smallest possible square grid** that can contain all provided tetrominoes.

**Placement Rules**

- Tetrominoes **cannot overlap**
- If a perfect square cannot be formed, remaining empty spaces are left as periods (`.`)
- Placement priority favours the **top-leftmost available position** (standard backtracking approach)

---

### 3.3 Output Formatting

| Requirement | Detail |
|---|---|
| **Identification** | Each tetromino is mapped to an uppercase Latin letter by input order — `A` for first, `B` for second, etc. |
| **Display** | The final solved square is printed to **standard output** |
| **Empty spaces** | Unfilled cells are printed as `.` |
| **Line endings** | Each row ends with a newline character (`\n`) |

---

## 4. Error Handling

> The program must output strictly the word `ERROR` (followed by a newline) under **any** of the following conditions:

| # | Condition |
|---|---|
| 1 | Wrong number of command-line arguments (`0` or `>1`) |
| 2 | File does not exist or cannot be read |
| 3 | The file contains `0` tetrominoes |
| 4 | Invalid characters found (anything other than `.`, `#`, and newlines) |
| 5 | Invalid grid dimensions (a tetromino grid that is not `4×4`) |
| 6 | Invalid block count (a grid containing more or fewer than `4` `#` characters) |
| 7 | Invalid shape (the `4` `#` characters do not form a contiguous tetromino) |
| 8 | Incorrect spacing (e.g., multiple blank lines between shapes, trailing blank lines) |

---

## 5. Technical Specifications

### 5.1 Proposed Architecture

```
main ──► parser ──► solver ──► printer
```

| Component | Responsibility |
|---|---|
| **Main** | Entry point — handles CLI arguments and orchestrates the flow |
| **Parser** | Reads the file, validates the `4×4` structure, validates block connections, returns a slice of `Tetromino` objects |
| **Solver** | Implements a backtracking algorithm to find the smallest valid square |
| **Printer** | Renders the solved 2D grid to the terminal with the correct alphabetical mapping |

**Size Optimisation**

The minimum starting square size is computed as:

```
size = ⌈√(N × 4)⌉
```

where `N` is the number of valid tetrominoes.

---

### 5.2 Maximum Limits

> Because pieces are mapped to uppercase Latin letters (`A–Z`), the system supports a **maximum of 26 tetrominoes**.
> The program must return `ERROR` if more than **26** valid tetrominoes are supplied.

---

## 6. Testing Strategy

### Unit Tests (`_test.go` files)

- Test file parsing with **valid** and **invalid** inputs
- Test connection validation — ensure disjointed blocks fail
- Test algorithm performance and correctness on known small sets (`1–5` pieces)

### Integration Tests

- End-to-end execution testing CLI arguments and stdout against expected baseline files
- **Stress test** with `26` pieces — ensure the backtracking algorithm resolves within a reasonable timeframe (no infinite loops or excessive memory consumption)