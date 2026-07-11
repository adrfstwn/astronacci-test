# Astronacci Test — Voucher Generator

An app for airline crew to check and generate vouchers for a given flight and date.

```
project/
├── frontend/          # React (Vite) application
├── backend/            # Go server (Gin) with SQLite
├── docker-compose.yml  # Runs both services together
└── README.md
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.25 or later
- [Node.js](https://nodejs.org/) 20 or later (includes npm)
- (Optional) [Docker](https://www.docker.com/) and Docker Compose, if you prefer running via containers

No external database is required — the backend uses an embedded SQLite database (file-based, created automatically on first run).

## 1. Install dependencies

**Backend**

```bash
cd backend
go mod download
```

**Frontend**

```bash
cd frontend
npm install
```

## 2. Run the backend

```bash
cd backend
go run ./cmd/server
```

- Starts the API on `http://localhost:8080`.
- On startup it automatically creates `vouchers.db` (SQLite) in the `backend/` folder and applies database migrations.
- Endpoints:
  - `POST /api/check` — `{ "flightNumber": "GA102", "date": "11-07-2026" }` → `{ "exists": bool }`
  - `POST /api/generate` — `{ "name", "id", "flightNumber", "date", "aircraft" }` → `{ "success": true, "seats": [...] }`

## 3. Run the frontend

In a separate terminal:

```bash
cd frontend
npm run dev
```

- Starts the app on `http://localhost:5173`.
- By default it calls the backend at `http://localhost:8080`. To point at a different backend URL, create a `frontend/.env` file:

  ```
  VITE_API_BASE_URL=http://localhost:8080
  ```

Open `http://localhost:5173` in your browser, fill in the crew/flight details, and click **Generate Vouchers**.

## 4. Docker instructions (optional)

To run both services together with Docker Compose:

```bash
docker compose up --build
```

- Frontend: `http://localhost:80`
- Backend: `http://localhost:8080`
- The SQLite database persists across restarts in the `backend-data` Docker volume.

To stop and remove the containers:

```bash
docker compose down
```

To build/run a single service's image manually:

```bash
# Backend
cd backend
docker build -t voucher-backend .
docker run -p 8080:8080 voucher-backend

# Frontend
cd frontend
docker build -t voucher-frontend .
docker run -p 80:80 voucher-frontend
```

## Tech stack

- **Frontend**: React 19 + Vite, Tailwind CSS
- **Backend**: Go, Gin, `modernc.org/sqlite` (pure-Go SQLite driver, no CGO required)
