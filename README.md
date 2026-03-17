<div align="center">
<h1>𝔊𝔬𝔭𝔥𝔢𝔯 𝔚𝔦𝔰𝔡𝔬𝔪</h1>

✧ ——— Code. Compile. OK. ——— ✧

<p>
A lightweight REST API built in Go using the Gin Gonic framework. This service serves a curated collection of impactful anime quotes to keep you inspired.
</p>

</div>

---

### Technical Dossier
* **Modern Engine**: Powered by `Gin Gonic` for high-concurrency routing and rapid JSON serialization.
* **Persistent Storage**: Integrated with `SQLite` via `GORM` for seamless data persistence and automated schema migrations.
* **Type Safety**: Strict Go struct definitions with `json` and `gorm` tags to ensure zero-leak data integrity.
* **Modular Architecture**: Clean separation between `internal/quotes` logic and the `cmd` entry point, following standard Golang project layouts.

---

### Public Endpoints
The following sectors are accessible for the general public:
- `GET /quotes` - Retrieve the entire collective quotes from the database.
- `GET /quotes/:id` - Query the database for the chosen quote by its associated ID.
- `POST /quotes` — Append a new transmission to the database.

---

### Command

This project utilizes `just` as a command runner for streamlined operations. Execute these from the root directory:

| Command | Action |
| :--- | :--- |
| `just awaken` | Initiates the REST service. |
| `just fetch` | Queries the database for all stored quotes. |
| `just post` | Transmits the predefined quote to the database. |

---

### Technical Setup
```bash
# Clone
git clone https://github.com/EternalHalve/gopher-wisdom.git
cd gopher-wisdom

# Initialize and sync modules
go mod tidy

# Awaken the service (manually)
go run cmd/gopher-wisdom/main.go
```