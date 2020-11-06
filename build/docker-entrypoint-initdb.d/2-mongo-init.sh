mongoimport -u mongo -p mongo --db ble_connections --collection users --drop --file /docker-entrypoint-initdb.d/connection.json --jsonArray
