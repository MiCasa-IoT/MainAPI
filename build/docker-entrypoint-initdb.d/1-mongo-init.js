var users = [
    {
        user: "mongo",
        pwd: "mongo",
        roles: [
            {
                role: "dbOwner",
                db: "mongo_example"
            }
        ]
    }
];

for (var i = 0, length = users.length; i < length; ++i) {
    db.createUser(users[i]);
}

db.createCollection('staffs');
