# Task 15: Dockerize Ascii-Art Web

**Objective:** Containerize the ASCII Art Web application using Docker to ensure portability and consistency across different environments. You must follow the same principles as the core subject.

## Steps

1.  **Create the Dockerfile**
    *   Create a file named `Dockerfile` at the root of the repository.
    *   Use an appropriate base image for the Go web server.
    *   Follow Docker best practices: utilize multi-stage builds to produce a small footprint, and ensure the container runs as a non-root user.
    *   Ensure that only the standard Go packages are used.

2.  **Apply Metadata**
    *   Apply metadata to your Docker objects using the `LABEL` instruction within the `Dockerfile` (e.g., author, description, and version).

3.  **Build the Docker Image**
    *   Use Docker client utilities to build the image (e.g., `docker build`).
    *   Ensure the web server code respects good practices and builds without issues.
    *   Understand containerizing an application, services, dependencies, and creating images.

4.  **Run the Docker Container**
    *   Launch a container from the image.
    *   Map the necessary ports so the web server can receive and output data over HTTP.
    *   Test the basics of the web application (HTML, HTTP) to ensure behavior is identical to the native setup.

5.  **Implement Garbage Collection (Clean Up)**
    *   Regularly clean up the environment.
    *   Take caution of unused objects (often referred to as "garbage collection") utilizing commands like `docker image prune`, `docker container prune`, or `docker system prune`.

## Acceptance Criteria

*   [x] A `Dockerfile` exists at the root of the repository.
*   [x] The `Dockerfile` adheres to best practices (minimal image size, multi-stage build, non-root user).
*   [x] The Docker image builds successfully without warnings (using `docker build`).
*   [x] The `Dockerfile` includes metadata `LABEL`s (e.g., maintainer, version, description).
*   [x] The web application runs successfully inside a container (using `docker run` and port mapping).
*   [x] The application behaves exactly the same in the container as it does natively.
*   [x] Developer is knowledgeable on how to clean up unused Docker objects.
