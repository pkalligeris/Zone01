# Task 09: Banner Selection (Developer 1 or Shared)

**Objective:** Ensure the correct banner file is loaded based on user input.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: Path Resolution
1.  **RED (Write Test):**
    *   Create `internal/banner/path_test.go` (or add to existing).
    *   Write `TestGetBannerPath`.
    *   Assert "shadow" -> "assets/banners/shadow.txt".
    *   Assert "standard" -> "assets/banners/standard.txt".
2.  **GREEN (Write Code):**
    *   Implement a helper function to map names to paths.

### Cycle 2: Integration
1.  **RED (Write Test):**
    *   Add integration test case GT-11 ("hello" shadow).
    *   Assert output matches shadow golden file.
2.  **GREEN (Write Code):**
    *   Update `cmd/ascii-art/main.go`.
    *   Use `Config.BannerFile` to call the path resolver, then `LoadBanner`.

## Acceptance Criteria
*   [x] Supports `standard`, `shadow`, `thinkertoy`.
*   [x] Returns error if banner is unknown or file missing.
*   [x] Defaults to `standard` if not specified.
