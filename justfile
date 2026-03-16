awaken:
    go run cmd/gopher-wisdom/main.go 

fetch:
    curl http://localhost:8080/quotes

post:
    curl --include \
    --request POST \
    --header "Content-Type: application/json" \
    --data '{"id": 3, "content": "I am a failure, I will prove them wrong", "anime": "Naruto", "character": "Naruto"}' \
    http://localhost:8080/quotes