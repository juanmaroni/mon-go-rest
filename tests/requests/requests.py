import requests

"""
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
"""
