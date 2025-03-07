# Mon-Go-REST
This is a side project to learn how to design and serve REST API using Go standard library.
Data is stored in MongoDB, initially imported with a Python script and a CSV.

Allowed HTTP operations: GET (one or all).

## Docker
From "devops/dev" directory, run:
```
docker compose --env-file .env.compose -f docker-compose.dev.yml up --build -d 
```

```
curl http://localhost:3333/
curl http://localhost:3333/api/v1/pokemon
curl http://localhost:3333/api/v1/pokemon/151
```
