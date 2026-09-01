# Proposed Workflow & Collaboration Strategy

The project is explicitly designed for **3 developers** to work in parallel. The task files (`task_01`, `task_02`, `task_03`) are decoupled so that each developer owns a specific layer of the application (Data, Input, Logic) without stepping on each other's toes.

Here is the breakdown of the workflow and a simulation of how your team can execute this using Git.

---

## 1. The Work Split

* **Developer 1 (Data Layer):** Works on `task_01_banner.md`. Responsible for reading files.
* **Developer 2 (Input Layer):** Works on `task_02_input.md`. Responsible for CLI args & validation.
* **Developer 3 (Logic Layer):** Works on `task_03_rendering.md`. Responsible for the ASCII generation algorithm.

---

## 2. Recommended Branching Strategy

Use the **Feature Branch Workflow**.

* **`main`**: The stable version of the code.
* **`feat/banner-loader`**: Dev 1's branch.
* **`feat/input-parser`**: Dev 2's branch.
* **`feat/rendering-engine`**: Dev 3's branch.

---

## 3. Workflow Simulation

Here is a simulation of how the team interacts with the repository.

### Phase 1: Foundation (Task 00 - Joint)
*The team meets. One person (or the lead) initializes the repo.*

```bash
# Lead Developer on 'main'
git init
go mod init ascii-art
# Creates directories and pkg/model/banner.go (The Shared Contract)
git add .
git commit -m "chore: init project structure and shared models"
git push origin main
```

> **Status:** The `pkg/model` folder exists. Everyone now agrees on what a `Banner` looks like.

### Phase 2: Parallel Development (Tasks 01, 02, 03)

Now, everyone pulls the latest `main` and creates their own branch.

**Developer 1 (Banner Loader):**
```bash
git checkout -b feat/banner-loader
# Implements internal/banner/loader.go
# Writes internal/banner/loader_test.go
git commit -m "feat: implement banner loading logic"
git push origin feat/banner-loader
```

**Developer 2 (Input Parser):**
```bash
git checkout -b feat/input-parser
# Implements internal/input/parser.go
# Writes internal/input/parser_test.go
git commit -m "feat: implement input validation and escape sequences"
git push origin feat/input-parser
```

**Developer 3 (Rendering Engine):**
```bash
git checkout -b feat/rendering-engine
# Implements internal/render/renderer.go
# Writes internal/render/renderer_test.go
# Note: Dev 3 doesn't need Dev 1's code yet. They mock the Banner data in their tests!
git commit -m "feat: implement rendering logic"
git push origin feat/rendering-engine
```

### Phase 3: Code Review & Merge

The developers open Pull Requests (PRs) to merge into `main`.

1.  **Dev 1's PR is merged.** `main` now has the Banner Loader.
2.  **Dev 2's PR is merged.** `main` now has Input Parsing.
3.  **Dev 3's PR is merged.** `main` now has the Renderer.

### Phase 4: Integration (Task 04)

The team (or the Tester) pulls the latest `main` which now contains all three pieces.

```bash
git checkout main
git pull origin main
# Now 'main' has internal/banner, internal/input, and internal/render
```

The Tester wires them together in `cmd/ascii-art/main.go`:

```go
// cmd/ascii-art/main.go
func main() {
    // Call Dev 2's code
    str := input.ParseInput(os.Args)

    // Call Dev 1's code
    font := banner.LoadBanner("standard.txt")

    // Call Dev 3's code
    output := render.Render(str, font)

    fmt.Print(output)
}
```

```bash
git commit -m "feat: wire application in main.go"
git push origin main
```

### Phase 5: Extended Features (4 Developers)

Once the core is stable, the team splits again to implement the advanced features.

**Prerequisite:** Define `pkg/model/config.go` on `main` first so everyone has the struct.

*   **Developer 1 (Input & Config):** `feat/input-config` (Task 05 & 09)
    *   Implements flag parsing.
    *   Implements banner selection logic.
*   **Developer 2 (Color):** `feat/color` (Task 06)
    *   Updates Renderer to handle ANSI codes.
*   **Developer 3 (Output):** `feat/file-output` (Task 07)
    *   Updates Main to write to files.
*   **Developer 4 (Align):** `feat/alignment` (Task 08)
    *   Updates Renderer to handle padding/terminal size.

**Integration:**

1.  Dev 1 merges `feat/input-config` (The Parser).
2.  Dev 2 merges `feat/color` (The Renderer update).
3.  Dev 4 merges `feat/alignment` (The Renderer update - watch for conflicts with Dev 2).
4.  Dev 3 merges `feat/file-output` (The Main update).

**Conflict Management:**

Dev 2 and Dev 4 both modify `internal/render`. They should communicate.
*   Dev 2 adds color codes *around* characters.
*   Dev 4 adds spaces *before* lines.
*   These changes are largely orthogonal but touch the same file.
