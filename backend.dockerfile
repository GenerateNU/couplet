FROM golang:latest

# Install Task
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

# Copy the backend into the container
COPY ./backend ./backend
COPY ./openapi.yaml ./openapi.yaml

# Set the container's working directory
WORKDIR ./backend

# Build and run the backend server
CMD ["task","build"]
CMD ["./server"]