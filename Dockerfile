FROM golang:1.23.0

WORKDIR /app

COPY . .

# Cleans up 'go.mod' file by removing any dependencies that are not used and adding any missing dependencies
RUN go mod tidy
