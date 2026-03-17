<div align="center">
<h1>𝔊𝔬𝔭𝔥𝔢𝔯 𝔚𝔦𝔰𝔡𝔬𝔪</h1>

✧ ——— Code. Compile. Collapse. ——— ✧

<p>
A lightweight, modular REST API built in Go. This service serves a curated collection of anime quotes.
</p>

</div>

---

### Technical Dossier
* **Modern Engine**: Powered by `Gin Gonic` for high-concurrency routing and rapid JSON serialization.
* **Persistent Storage**: Integrated with `SQLite` via `GORM` with automated schema migrations.
* **Asynchronous Workers**: Features a background `statusWorker` for real-time service monitoring and graceful shutdown handling.
* **Modular Architecture**: Follows the Standard Go Project Layout with logic encapsulated in `internal/` to ensure a clean, decoupled dependency graph.

---

### Public Endpoints
The following sectors are accessible (API Version: `v1`):
- `GET /api/v1/quotes` - Retrieve the entire collective quotes from the database.
- `GET /api/v1/quotes/:id` - Query the database for the chosen quote by its associated ID.
- `POST /api/v1/quotes` — Append a new transmission to the database.

---

### Command

This project utilizes `just` as a command runner. Execute these from the root directory:

| Command | Action |
| :--- | :--- |
| `just awaken` | Initiates the REST service and background workers. |
| `just fetch` | Queries the database for all stored quotes. |
| `just post` | Transmits the predefined quote to the database. |

---

### Technical Setup
```bash
# Clone
git clone [https://github.com/EternalHalve/gopher-wisdom.git](https://github.com/EternalHalve/gopher-wisdom.git)
cd gopher-wisdom

# Configure Environment
cp .env.example .env  # Ensure DB_NAME is defined

# Initialize and sync modules
go mod tidy

# Awaken the service (manually)
go run cmd/server/main.go