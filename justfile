awaken:
    go run cmd/gopher-wisdom/main.go 

fetch:
    curl http://localhost:8080/quotes

post:
    curl --include \
    --request POST \
    --header "Content-Type: application/json" \
    --data '{"id": 3, "content": "Even the strongest of opponents always has a weakness", "anime": "Naruto", "character": "Itachi Uchiha"}' \
    http://localhost:8080/quotes