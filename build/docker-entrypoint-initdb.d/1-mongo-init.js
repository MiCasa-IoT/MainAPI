let users = [
    {
        user: "mongo",
        pwd: "mongo",
        roles: [
            {
                role: "dbOwner",
                db: "ble_connections"
            }
        ]
    }
];

for (var i = 0, length = users.length; i < length; ++i) {
    db.createUser(users[i]);
}

db.createCollection('users');
