# Mon-Go-REST
This is a side project to learn how to design and serve REST API using Go standard library.
Data is stored in MongoDB, initially imported with a Python script and a CSV.

Allowed HTTP operations: GET (one or all).

## Run with Docker
From "devops/dev" directory, run:
```
docker compose --env-file .env.dev.compose -f docker-compose.dev.yml up --build -d 
```

And check endpoints:
```
# Homepage
curl http://localhost:3333/

# Get all
curl http://localhost:3333/api/v1/pokemon

# Get one
curl http://localhost:3333/api/v1/pokemon/151
```
