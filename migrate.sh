#!/bin/bash
docker compose exec app bash -c "cd ./database && sql-migrate up"