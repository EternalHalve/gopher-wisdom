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
* **Traffic Control**: Integrated Rate Limiting via `limiter` to protect the archives from over-extraction.
* **Containerized Environment**: Fully Dockerized using multi-stage builds for a minimal runtime footprint
* **Persistent Storage**: Integrated with `SQLite` via `GORM` with automated schema migrations.
* **Asynchronous Workers**: Features a background `statusWorker` for real-time service monitoring and graceful shutdown handling.

---

### Public Endpoints
The following sectors are accessible (API Version: `v1`):
- `GET /api/v1/quotes` - Retrieve the entire collective quotes from the database.
- `GET /api/v1/quotes/:id` - Query the database for the chosen quote by its associated ID.
- `GET /api/v1/quotes/:id?format=alien` - [NEW] Decrypt the transmission into Zorgon-7 dialect.
- `POST /api/v1/quotes` - Append a new transmission to the database.

---

### Command

This project utilizes `just` as a command runner. Execute these from the root directory:

| Command | Action |
| :--- | :--- |
| `just awaken` | [Primary] Builds and awakens the Dockerized service in detached mode. |
| `just sleep` | Collapses the containerized environment. |
| `just logs` | Streams the real-time heartbeat and server logs from the container. |
| `just up` | Initiates the REST service locally. |
| `just clean` | Clean up Docker images and the local data folder |
| `just fetch` | Queries the database for all stored quotes. |
| `just post` | Transmits the predefined quote to the database. |

---

### Technical Setup
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