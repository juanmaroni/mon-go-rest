from csv import DictReader
from pymongo import MongoClient


class MongoConnection:
    def __init__(self, client, db, collection):
        self.client = client
        self.db = db
        self.collection = collection

    def disconnect(self):
        self.client.close()

    def insert_data_csv(self, csv_path, overwrite_collection, primary_key=None, primary_key_type=int):
        collection = self.collection
        count_coll_docs = collection.estimated_document_count()
        
        if count_coll_docs != 0 and not overwrite_collection:
            return 1, "Collection is not empty."
        
        collection.drop()

        with open(csv_path, mode="r", encoding="utf-8") as file:
            reader = DictReader(file, delimiter=';')
            documents = []

            for row in reader:
                if primary_key in row:
                    row["_id"] = primary_key_type(row.pop(primary_key))

                documents.append(row)

            if documents:
                collection.insert_many(documents)
        
        return 0, f"Inserted {len(documents)} documents in collection '{collection.name}'."


def connect(host, port, db_name, collection_name):
    client = MongoClient(host, port)
        
    if not client:
        return None, f"Couldn't connect to MongoDB with host '{host}' and port '{port}'."
    
    db_names = client.list_database_names()

    if db_name not in db_names:
        return None, f"Database '{db_name}' doesn't exist."
        
    db = client.get_database(db_name)
    collection_names = db.list_collection_names()

    if collection_name not in collection_names:
        return MongoConnection(client, db, None), None
    
    collection = db.get_collection(collection_name)
    
    return MongoConnection(client, db, collection), None
