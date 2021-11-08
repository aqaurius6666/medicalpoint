# MEDICALPOINT GATEWAY

# Run
`docker-compose -f docker-compose-production.yaml up -d`
## Port
* 10998: gate-way api port
* 26259: database port. Postgresql, Cockroachdb using the same connection config. username: root, password: root, database: defaultdb.
* 9191: blockchain grpc port
* 26657: blockchain rpc port

# API doc
Go to `http://localhost:10998/swagger/index.html`

# Reset data
- Restart blockchain container
- Exec in gateway container, and `server seed-data --clean --super-admin-id <SUPER_ADMIN_ID>`
