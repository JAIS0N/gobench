# GoBench: Dockerized REST API Benchmarking (Go vs Node vs Python)

GoBench is a backend performance benchmarking project comparing equivalent REST APIs implemented in **Go (Gin)**, **Node.js (Express)**, and **Python (FastAPI)**.

The goal is to evaluate how different runtimes handle concurrent HTTP workloads under identical conditions.

---

## 🚀 Tech Stack

- Go (Gin)
- Node.js (Express)
- Python (FastAPI)
- Docker & Docker Compose
- `hey` (load testing)

---

## 📦 Project Structure

```
gobench/
├── services/
│   ├── go-gin/
│   ├── node-express/
│   └── python-fastapi/
├── benchmarks/
├── docker-compose.yml
└── README.md
```

---

## ⚙️ Running the Project

### 1. Start all services

```bash
docker compose up --build
```

### 2. Access endpoints

| Service | URL |
|---------|-----|
| Go | http://localhost:8081/ping |
| Node | http://localhost:8082/ping |
| Python | http://localhost:8083/ping |

---

## 🔁 API Endpoints

| Endpoint | Description |
|----------|-------------|
| `/ping` | Lightweight response (used for benchmarking) |
| `/compute` | CPU-bound computation workload |

---

## 📊 Benchmark Setup

- **Tool:** `hey`
- **Requests:** 10,000
- **Concurrency:** 100
- **Environment:** Docker (local machine)

---

## 📈 Benchmark Results

| Service | Requests/sec | Avg Latency | p95 Latency | p99 Latency |
|---------|-------------|-------------|-------------|-------------|
| Go (Gin) | 15,472 | 6.3 ms | 10.2 ms | 66.3 ms |
| Node (Express) | 9,416 | 9.7 ms | 16.2 ms | 51.3 ms |
| Python (FastAPI) | 3,328 | 29.8 ms | 40.0 ms | 98.1 ms |

---

## 🧠 Key Observations

- **Go** achieved the highest throughput (~15K req/sec) with the lowest average latency.
- **Node.js** performed moderately well with good latency consistency.
- **Python FastAPI** showed higher latency under concurrent load.
- Differences highlight runtime efficiency and concurrency model impact.

---

## 🔗 How Services Communicate

External access uses host ports:
- `8081` → Go
- `8082` → Node
- `8083` → Python

Internal Docker communication uses service names:
- `http://go-gin:8080`
- `http://node-express:8080`
- `http://python-fastapi:8080`

---

## 🐳 Docker Architecture

Each service runs in its own container with:
- Isolated runtime
- Independent port (`8080` internally)
- Host port mapping via Docker Compose

---

## 🎯 Key Learnings

- Containerized multi-service architecture using Docker Compose
- Performance benchmarking under concurrent workloads
- Differences in concurrency models (Go vs Node vs Python)
- Practical understanding of port mapping and container networking

---

## 📌 Future Improvements

- [ ] Add Grafana dashboard for visualization
- [ ] Automate benchmarking pipeline
- [ ] Add distributed load testing
- [ ] Extend with gRPC comparison

---

## 👤 Author

**Jaison Felix Menezes**
