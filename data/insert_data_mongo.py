import sys
import csv
from pymongo import MongoClient

def insert_data_mongo(host="localhost", port=27017, db_name="pokemon", collection_name="kanto", csv_path="./pokedata.csv"):
    client = MongoClient(host, port)
    db = client[db_name]
    db.drop_collection(collection_name) # In case collection already exists
    collection = db[collection_name]

    with open(csv_path, mode="r", encoding="utf-8") as file:
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

    client.close()

if __name__ == "__main__":
    # Example: python insert_data_mongo.py "localhost" 27017 "pokemon" "kanto" "./pokedata.csv"
    args = sys.argv
    insert_data_mongo(host=args[1], port=int(args[2]), db_name=args[3], collection_name=args[4], csv_path=args[5])

