# Task 12: HTML Template & Styling

**Objective:** Create the HTML template with a dark-themed responsive design, form inputs, and conditional result/error display.

## Steps

1.  **Create `templates/index.html`**
    *   Dark-themed layout with background image (`assets/unnamed (1).png`).
    *   Responsive container centered on the page.

2.  **Form Implementation**
    *   Textarea for `text` input with placeholder.
    *   Radio buttons for `banner` selection: `standard`, `shadow`, `thinkertoy`.
    *   Submit button styled with accent color.

3.  **Template Logic**
    *   Use Go template actions (`{{if}}`, `{{eq}}`) for:
        *   Retaining user input (`{{.Input}}`) after submission.
        *   Keeping the selected banner checked (`{{if eq .Banner "..."}}`).
        *   Conditionally displaying `.Result` in a `<pre>` block.
        *   Conditionally displaying `.Error` in a styled error block.

4.  **Styling**
    *   Dark background with semi-transparent container.
    *   Monospace font for textarea and result output.
    *   Styled error messages with distinct color and border.
    *   Hover effects on the submit button.

## Acceptance Criteria
*   [x] Form displays with text input and three banner radio buttons.
*   [x] Dark theme with background image renders correctly.
*   [x] User input and banner selection are retained after submission.
*   [x] Result is displayed in a monospace `<pre>` block.
*   [x] Errors are displayed in a styled error container.
