# Stage 1: Build modules
FROM golang:1.22 as modules-build

# Set the working directory
WORKDIR /app

# Copy the module sources relative to the build context
COPY modules/v2/common /modules/common
COPY modules/v2/database/d7redis /modules/database/d7redis
COPY modules/v2/database/d7mysql /modules/database/d7mysql

# Copy go workspace script and packages.json
COPY scripts/init_workspace.sh /app/init_workspace.sh
COPY auth-service/v2/workspace.packages.json /app/workspace.packages.json

# Install dependencies
RUN apt-get update && apt-get install -y jq

# Initialize and install common modules
RUN chmod +x init_workspace.sh && ./init_workspace.sh

# Stage 2: Build main app
FROM golang:1.22 as main-build

# Set the working directory
WORKDIR /app

# Copy the modules and go workspace from the previous stage
COPY --from=modules-build /modules /modules
COPY --from=modules-build /app/go.work /app/go.work

# Copy the main app source code relative to the build context
COPY auth-service/v2 /app

# Ensure go.mod and go.sum exist and print them for debugging
COPY auth-service/v2/go.mod /app/go.mod
COPY auth-service/v2/go.sum /app/go.sum
RUN ls -la /app

# Download Go modules with verbose output for debugging
RUN go mod download

# Copy the remaining source code and build the Go app
COPY auth-service/v2 /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/main main.go

# Stage 3: Create a minimal Docker image with the built binary
FROM scratch

# Copy the built binary from the builder stage
COPY --from=main-build /bin/main /bin/main

# Expose the port the app runs on
EXPOSE 3001

# Command to run the executable
CMD ["/bin/main"]
