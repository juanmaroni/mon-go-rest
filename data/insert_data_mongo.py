import sys
from mongo.driver import connect


def main(host, port, db_name, collection_name, csv_path, overwrite_collection=False):
    conn, err = connect(host, port, db_name, collection_name)

    if err:
        print(err) # Log
        return
    
    print("Info: connected to MongoDB.") # Log
    
    if conn.collection is None:
        print(f"Warn: collection '{collection_name}' does not exist.")
        conn.collection = conn.db.create_collection(collection_name)
        print(f"Warn: new collection '{collection_name}' created.")

    msg = conn.insert_data_csv(csv_path, overwrite_collection, "number")
    print(msg) # Log
    conn.disconnect()
    print("Info: disconnected from MongoDB.") # Log


if __name__ == "__main__":
    args = sys.argv
    n_args = len(args)

    if n_args != 6 and n_args != 7:
        print("Error: not enough arguments.") # Log
        sys.exit(1)
    
    if not all(args):
        print("Error: empty value arguments not allowed.") # Log
        sys.exit(1)

    # Check optional parameters
    overwrite_collection = False

    if n_args == 7 and args[6] == "y":
        overwrite_collection = True      
    
    # Example: python insert_data_mongo.py localhost 27017 pokemon kanto csv/pokemon_kanto.csv y 
    main(
        host=args[1],
        port=int(args[2]),
        db_name=args[3],
        collection_name=args[4],
        csv_path=args[5],
        overwrite_collection=overwrite_collection
    )
    sys.exit(0)
