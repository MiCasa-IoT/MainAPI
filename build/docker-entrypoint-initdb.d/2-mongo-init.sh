mongoimport -u mongo -p mongo --db mongo_example --collection users --drop --file /docker-entrypoint-initdb.d/connection.json --jsonArray
