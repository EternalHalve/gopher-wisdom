<div align="center">
<h1>𝔖𝔨𝔶 𝔔𝔲𝔬𝔱𝔢𝔰</h1>

✧ ——— Code. Compile. OK. ——— ✧

<p>
A lightweight REST API built in Go using the Gin Gonic framework. This service serves a curated collection of impactful anime quotes to keep you inspired.
</p>
</div>

---

### Technical Dossier
* **Modern Engine**: Powered by `Gin Gonic` for high-concurrency routing and rapid JSON serialization.
* **Type Safety**: Strict Go struct definitions with `json` tags to ensure zero-leak data integrity.
* **Modular Architecture**: Clean separation between `internal/quotes` logic and the `cmd` entry point, following standard Golang project layouts.

---

### Public Endpoints
The following sectors are currently broadcasting on the public frequency:
- `GET /quotes` - Retrieve the entire collective consciousness of the archive.

---

### Technical Setup
```bash
# Clone
git clone https://github.com/EternalHalve/sky-quotes.git
cd sky-quotes

# Initialize and sync modules
go mod tidy

# Awaken the service
go run cmd/sky-quotes/main.go
```