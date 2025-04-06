# Mon-Go-REST
![Go](https://img.shields.io/badge/Go-1.24.0-blue)
![MongoDB](https://img.shields.io/badge/MongoDB-%5E7.0-green)
![Docker](https://img.shields.io/badge/Docker-ready-blue)

This is a side project to learn how to design and serve REST API using Go standard library.
Data is stored in MongoDB, initially imported with Python scripts and CSV files.

Allowed HTTP operations: GET (one or all).

## How to run
### Docker
Make sure you have Docker and Docker Compose installed.

From "devops/dev" directory, run:
```
docker compose --env-file .env.dev.compose -f docker-compose.dev.yml up --build -d 
```

### Check endpoints
With cURL:
```
# Homepage
curl http://localhost:3333/

# Get all
curl http://localhost:3333/api/v1/pokemon

# Get all by region
curl http://localhost:3333/api/v1/pokemon/kanto

# Get one by Pokédex number (id)
curl http://localhost:3333/api/v1/pokemon/151

# Errors
curl http://localhost:3333/api/v1/pokemon/otnak
curl http://localhost:3333/api/v1/pokemon/9999
curl -X POST http://localhost:3333/api/v1/pokemon/9999
```

Or running "tests/check_endpoints.py".
