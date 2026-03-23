<div align="center">
<h1>𝔊𝔬𝔭𝔥𝔢𝔯 𝔚𝔦𝔰𝔡𝔬𝔪</h1>

✧ ——— Code. Compile. Collapse. ——— ✧

<p>
A lightweight, modular REST API built in Go. This service serves a curated collection of anime quotes.
</p>

</div>

---

### ✦ Technical Dossier

- **Modern Engine**  
  Powered by `Gin Gonic` for high-concurrency routing and fast JSON serialization.

- **Traffic Control**  
  Integrated rate limiting via `limiter` to prevent over-extraction.

- **Containerized Environment**  
  Fully Dockerized with multi-stage builds for a minimal runtime footprint.

- **Persistent Storage**  
  Uses `SQLite` with `GORM`, including automated schema migrations.

- **Asynchronous Workers**  
  Background `statusWorker` handles real-time monitoring and graceful shutdowns.

---

### ✦ Public Endpoints (v1)

- `GET /api/v1/quotes` - Retrieve the entire collective quotes from the database.
- `GET /api/v1/quotes/:id` - Query the database for the chosen quote by its associated ID.
- `GET /api/v1/quotes/:id?format=alien` - [NEW] Decrypt the transmission into Zorgon-7 dialect.
- `POST /api/v1/quotes` - Append a new transmission to the database.

---

### ✦ Command Interface

This project utilizes `just` as a command runner. Execute these from the root directory.

| Command | Action |
|:--|:--|
| `just awaken` | **[Primary]** Build and start the Dockerized service (detached) |
| `just sleep` | Stop the containerized environment |
| `just logs` | Stream real-time server logs |
| `just up` | Run the service locally |
| `just clean` | Remove Docker images and local data |
| `just fetch` | Retrieve all stored quotes |
| `just post` | Insert a predefined quote |

---

### ✦ Technical Setup
```bash
# Clone
git clone [https://github.com/EternalHalve/gopher-wisdom.git](https://github.com/EternalHalve/gopher-wisdom.git)
cd gopher-wisdom

# Configure Environment
cp .env.example .env

# Awaken the service
just awaken

# Cleaning up the trash
just clean
