# Golden File Testing Strategy

## Objective
To ensure regression testing by comparing the program's actual output against pre-verified "golden" output files. This guarantees that refactoring or new features do not break existing functionality.

## Test Suite

The following test cases define the input and the expected behavior.

| Case ID | Input String | Description | Expected Output File |
| :--- | :--- | :--- | :--- |
| **GT-01** | `"hello"` | Basic lowercase word | `test/golden/hello.txt` |
| **GT-02** | `"HELLO"` | Basic uppercase word | `test/golden/HELLO.txt` |
| **GT-03** | `"HeLlo WoRlD"` | Mixed case with spaces | `test/golden/mixed_case.txt` |
| **GT-04** | `"1234567890"` | Numbers | `test/golden/numbers.txt` |
| **GT-05** | `"!@#$%^&*()"` | Special characters | `test/golden/special_chars.txt` |
| **GT-06** | `"Hello\nThere"` | Embedded newline | `test/golden/multiline.txt` |
| **GT-07** | `"\n"` | Single newline (should print empty line) | `test/golden/newline_only.txt` |
| **GT-08** | `""` | Empty string (should be handled gracefully) | `test/golden/empty.txt` |
| **GT-09** | `"Hello\n\nWorld"` | Multiple consecutive newlines | `test/golden/double_newline.txt` |
| **GT-10** | `"ABCDEFGHIJKLMNOPQRSTUVWXYZ"` | Full uppercase alphabet | `test/golden/all_upper.txt` |
| **GT-11** | `"hello" shadow` | Banner selection (Shadow) | `test/golden/shadow_hello.txt` |
| **GT-12** | `"hello" thinkertoy` | Banner selection (Thinkertoy) | `test/golden/thinkertoy_hello.txt` |
| **GT-13** | `--align=right "hello"` | Right alignment (Fixed 80 cols) | `test/golden/align_right.txt` |
| **GT-14** | `--align=center "hello"` | Center alignment (Fixed 80 cols) | `test/golden/align_center.txt` |
| **GT-15** | `--align=justify "A B"` | Justify alignment (Fixed 80 cols) | `test/golden/align_justify.txt` |
| **GT-16** | `--color=red "hello"` | Color full string (ANSI) | `test/golden/color_red.txt` |
| **GT-17** | `--color=green "l" "hello"` | Color substring (ANSI) | `test/golden/color_substring.txt` |
| **GT-18** | `--output=test_out.txt "hello"` | File output verification | `test/golden/output_file.txt` |

## Implementation Plan

> Bash note: for inputs containing `!` (for example GT-05), use single quotes to avoid history expansion errors:
> `go run ./cmd/ascii-art '!@#$%^&*()'`

1.  **Generate Golden Files:**
    *   **Constraint:** For Alignment tests (GT-13 to GT-15), the test runner must mock or force a terminal width of **80 columns** to ensure deterministic output.
    *   **Constraint:** For Color tests (GT-16, GT-17), the output file will contain raw ANSI escape codes.
    *   **Constraint:** For Output tests (GT-18), the verification compares the content of the *created file* against the golden file.
    *   Run a verified version of the code.
    *   Redirect output: `go run . "hello" > test/golden/hello.txt`.
    *   Manually inspect `hello.txt` to confirm it is correct.

2.  **Automated Test Runner (`integration_test.go`):**
    *   Iterate through the table of test cases.
    *   For each case:
        1.  Capture `stdout` of the `main()` function or `Render()` function.
        2.  Read the content of the corresponding `.txt` file.
        3.  Compare `actual_output` vs `expected_output`.
        4.  Fail the test if they differ.

## Maintenance
*   If the rendering logic changes intentionally (e.g., fixing a bug in the font), regenerate the golden files and verify them manually.
