mongoimport -u mongo -p mongo --db mongo_example --collection users --drop --file /docker-entrypoint-initdb.d/users.json --jsonArray
