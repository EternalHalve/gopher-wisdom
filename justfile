# Build and start the container in the background
awaken:
    docker compose up --build -d

# Stop and remove containers and networks
sleep:
    docker compose down

# Follow the application logs
logs:
    docker compose logs -f

# Run the project locally (legacy/dev mode)
up:
    go run cmd/gopher-wisdom/main.go 

# Fetch all quotes
fetch:
    curl http://localhost:8080/api/v1/quotes

# Post a new quote
post:
    curl --include \
    --request POST \
    --header "Content-Type: application/json" \
    --data '{"content": "Even the strongest of opponents always has a weakness", "anime": "Naruto", "character": "Itachi Uchiha"}' \
    http://localhost:8080/api/v1/quotes

# Clean up Docker images and the local data folder
clean:
    docker compose down --rmi all
    rm -rf ./data

test:
    go test -v ./internal/quotes