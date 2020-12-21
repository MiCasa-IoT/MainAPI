let users = [
    {
        user: "mongo",
        pwd: "mongo",
        roles: [
            {
                role: "dbOwner",
                db: "micasadb"
            }
        ]
    }
];

for (var i = 0, length = users.length; i < length; ++i) {
    db.createUser(users[i]);
}

db.createCollection('users');
