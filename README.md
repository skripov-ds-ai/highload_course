# highload_course

Server is available on `8888` port.

Postgres DB is available on `5432` port.


## Launch instructions
1. Build `docker compose build`.
2. Run Postgres DB `docker compose up -d db`.
3. Run DB migrations `docker compose up db_migration`.
4. Run server `docker compose up -d app`.
5. Shutdown all `docker compose down`.
