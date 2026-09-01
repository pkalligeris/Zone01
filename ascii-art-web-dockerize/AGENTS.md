# ASCII Art Web — Agent Instructions

## Mandatory: Activity Logging

Every conversation **must** end with an append to `docs/logs/ai.log` documenting all actions taken. Follow the established format exactly:

```
Timestamp: YYYY-MM-DD User Request: "<quoted user request>"

AI Assistant Output & Actions:

<Category>:
- <bullet point describing each action taken>
- <include file paths, commands run, and verification results>

Log Update: Appended this action to `ai.log`.
```

### Rules

1. **Read `docs/logs/ai.log` at the start** of every conversation to understand the current state.
2. **Append** — never overwrite or modify existing log entries.
3. **Use the exact format** shown above (Timestamp, User Request, AI Assistant Output & Actions, category headers, bullet items, closing "Log Update" line).
4. **Group actions** under descriptive category headers (e.g., "Bug Fix:", "Documentation Update:", "Feature Implementation:", "Git Workflow:", "Verification:").
5. **Include verification results** when tests or commands are run (e.g., "Ran `go test ./...` → PASS").
6. **Log every request** — even clarifications, reviews, or confirmations get a log entry.

## Project Context

Before making changes, review these docs for context:

- `docs/prd.md` — Product requirements
- `docs/architecture.md` — System architecture
- `docs/tasks/` — Task definitions and acceptance criteria
- `docs/milestones.md` — Project milestones
