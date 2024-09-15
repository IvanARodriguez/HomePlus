# HomePlus

This project consists of a Go backend and a React frontend using Vite, both containerized using Docker. The setup includes hot-reloading for both the backend and frontend using Air for Go and pnpm for the React app.

## Prerequisite

Make sure you have the following installed on your system:

- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/#download)

## Running the Project

To start both the backend and frontend for development using Docker Compose, you can run:

```bash
  make dev
```

### This will:

- Build and spin up the Go backend and React frontend in separate Docker containers.
- Enable hot-reloading for both the backend (via Air) and frontend (via pnpm with Vite).

## Accessing the Application

The React frontend will be available at http://localhost:3000.
The Go backend will be available at http://localhost:8000.

## Conclusion

This setup allows for rapid development with hot-reloading in both the Go backend and React frontend environments. Adjust the Dockerfile or Makefile as needed for production deployment.
