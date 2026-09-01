# Task 05: Input Configuration (Developer 1)

**Objective:** Update the shared model and implement robust flag parsing to support the new features.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: Config Model & Defaults
1.  **RED (Write Test):**
    *   Create `internal/input/flags_test.go`.
    *   Write `TestParseArgs_Defaults`. Call `ParseArgs([]string{"hello"})`.
    *   Assert `Config.BannerFile` is "standard", `Align` is "left", `Color` is empty, `OutputFile` is empty.
2.  **GREEN (Write Code):**
    *   Create `pkg/model/config.go` with the `Config` struct.
    *   Create `internal/input/flags.go`. Implement `ParseArgs` to return the default config for simple input.

### Cycle 2: Flag Extraction
1.  **RED (Write Test):**
    *   Write `TestParseArgs_Flags`.
    *   Input: `[]string{"--color=red", "--align=right", "--output=out.txt", "hello"}`.
    *   Assert `Config` fields match the flags.
2.  **GREEN (Write Code):**
    *   Implement loop in `ParseArgs` to detect strings starting with `--`.
    *   Split by `=` to get key/value.
    *   Populate `Config`.

### Cycle 3: Positional Arguments & Validation
1.  **RED (Write Test):**
    *   Write `TestParseArgs_Positional`.
    *   Test 2 args: `["hello", "shadow"]` -> Banner should be "shadow".
    *   Test Color ambiguity: `["--color=red", "hello", "shadow"]` -> Input "hello", Banner "shadow".
    *   Test Color substring: `["--color=red", "he", "hello"]` -> Substring "he", Input "hello".
    *   Test Invalid Flag format -> Expect specific usage error.
    *   Test Invalid Banner -> Expect specific usage error listing banners.
2.  **GREEN (Write Code):**
    *   Implement logic to handle remaining non-flag arguments.
    *   Implement `isBanner` check to disambiguate the second argument in color mode.
    *   Add validation logic for flag formats and banner existence.
3.  **REFACTOR:**
    *   Clean up the parsing loop. Ensure specific error messages match PRD exactly.

## Acceptance Criteria
*   [x] `pkg/model/config.go` exists.
*   [x] `ParseArgs` returns a populated `Config` object.
*   [x] Correctly identifies `[STRING]` vs `[BANNER]`.
*   [x] Correctly parses `--color`, `--output`, `--align`.
*   [ ] Returns specific error messages for bad flags as per PRD.
