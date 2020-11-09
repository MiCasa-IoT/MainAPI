mongoimport -u mongo -p mongo --db micasadb --collection users --drop --file /docker-entrypoint-initdb.d/connection.json --jsonArray
