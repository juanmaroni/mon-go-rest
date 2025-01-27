import csv
from pymongo import MongoClient

HOST = "localhost"
PORT = 27017
DB_NAME = "pokemon"
COLLECTION_NAME = "kanto"
CSVFILE_PATH = "./data/pokedata.csv"

client = MongoClient(HOST, PORT)
db = client[DB_NAME]
db.drop_collection(COLLECTION_NAME) # In case collection already exists
collection = db[COLLECTION_NAME]

with open(CSVFILE_PATH, mode="r", encoding="utf-8") as file:
    reader = csv.DictReader(file, delimiter=';')
    primary_key = "number"
    documents = []

    for row in reader:
        if primary_key in row:
            row["_id"] = int(row.pop(primary_key))

        documents.append(row)

    if documents:
        collection.insert_many(documents)
    
    print(f"{len(documents)} documents inserted.")
