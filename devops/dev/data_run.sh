python -m venv venv && venv/bin/pip install --no-cache-dir -r requirements.txt
. venv/bin/activate
python insert_data_mongo.py mongo 27017 pokemon kanto csv/pokemon_kanto.csv y
python insert_data_mongo.py mongo 27017 pokemon johto csv/pokemon_johto.csv y
