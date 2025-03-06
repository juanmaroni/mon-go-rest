from sys import exit, argv
from config.logger import setup_logger
from mongo.driver import connect


def main(host, port, db_name, collection_name, csv_path, overwrite_collection=False):
    conn, err = connect(host, port, db_name, collection_name)

    if err:
        logger.error(err)
        return 1
    
    logger.info("Connected to MongoDB.")
    
    if conn.collection is None:
        logger.warning(f"Collection '{collection_name}' does not exist in database {db_name}.")
        conn.collection = conn.db.create_collection(collection_name)
        logger.info(f"New collection '{collection_name}' created in database {db_name}.")

    logger.info(f"Inserting data in collection '{collection_name}'...")
    code, msg = conn.insert_data_csv(csv_path, overwrite_collection, "number")

    if code == 1:
        logger.error(msg)
    else:
        logger.info(msg)
    
    conn.disconnect()
    logger.info("Disconnected from MongoDB.")
    
    return code


if __name__ == "__main__":
    logger = setup_logger()
    args = argv[1:]
    logger.info(f"List of arguments: {args}")
    n_args = len(args)

    if n_args != 5 and n_args != 6:
        logger.error("Not enough arguments.")
        exit(1)
    
    if not all(args):
        logger.error("Empty arguments not allowed.")
        exit(1)

    # Check optional parameters
    overwrite_collection = False

    if n_args == 6 and args[5] == "y":
        overwrite_collection = True      
    
    # Example: python insert_data_mongo.py localhost 27017 pokemon kanto csv/pokemon_kanto.csv y
    code = main(
        host=args[0],
        port=int(args[1]),
        db_name=args[2],
        collection_name=args[3],
        csv_path=args[4],
        overwrite_collection=overwrite_collection
    )
    exit(code)
