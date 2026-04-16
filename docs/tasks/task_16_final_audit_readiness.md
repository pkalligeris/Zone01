# Task 16: Final Audit Readiness

**Objective:** Complete the final polish required for project submission by aligning the Docker workflow, runtime container inspection, documentation, and verification steps with the expected audit experience.

## Steps

1.  **Validate the Docker Audit Flow End-to-End**
    *   Re-run the full auditor flow using the project `Dockerfile`:
        *   `docker image build -f Dockerfile -t <image-name> .`
        *   `docker images`
        *   `docker container run -p <host-port>:8080 --detach --name <container-name> <image-name>`
        *   `docker ps -a`
        *   `docker exec -it <container-name> /bin/bash`
    *   Confirm the application responds correctly on port `8080` from inside the containerized environment.

2.  **Polish the Runtime Filesystem Layout**
    *   Ensure the running container clearly exposes only the runtime artifacts needed by the application.
    *   Confirm the executable, templates, and static asset directories are easy for an auditor to locate during `docker exec`.
    *   If the runtime layout is confusing or inconsistent, standardize it and update the documentation accordingly.

3.  **Verify the Docker Helper Script**
    *   Confirm `dockerize.sh` remains in sync with the real image name, container name, exposed port, and shell instructions.
    *   Make sure the script provides a simple one-command workflow for building and running the project.

4.  **Close Documentation Gaps**
    *   Review `README.md`, `docs/prd.md`, and `docs/architecture.md` for any mismatch with the implemented Docker, API, and web behavior.
    *   Ensure the docs explicitly mention:
        *   standard-library-only Go dependencies
        *   the JSON API endpoint
        *   the Docker helper workflow
        *   the expected verification commands for the audit

5.  **Run Final Verification**
    *   Run `go test ./...`.
    *   Rebuild the Docker image after any final edits.
    *   Confirm the web page loads, the API responds, and the container stays up.
    *   Clean up temporary audit containers/images if they are no longer needed.

## Acceptance Criteria

*   [x] The full Docker audit flow succeeds using the repository `Dockerfile`.
*   [x] The runtime container filesystem is easy to inspect and contains the required executable and runtime directories.
*   [x] `dockerize.sh` correctly automates the current build/run workflow.
*   [x] `README.md`, `docs/prd.md`, and `docs/architecture.md` match the implemented project behavior.
*   [x] `go test ./...` passes after the final audit-readiness changes.
*   [x] The project is ready for submission without undocumented setup steps or ambiguous auditor instructions.
