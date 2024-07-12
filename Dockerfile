FROM golang:1.18

# Install necessary tools
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /workspace

# Copy the current directory contents into the container at /workspace
COPY . /workspace

# Download go modules
RUN go mod tidy

# Expose the application on port 8080
EXPOSE 8080

# Start a shell session by default
CMD ["bash"]
