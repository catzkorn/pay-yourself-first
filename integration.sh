#!/usr/bin/env bash

docker run --name pay-yourself-first-postgres -p 6969:5432 -e POSTGRES_PASSWORD=mysecretpassword --rm -d postgres
sleep 2
psql -f ./db/migrations/001_budget.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres
psql -f ./db/migrations/003_savings.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres
psql -f ./db/migrations/004_drop_income_id_add_date.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres
psql -f ./db/migrations/005_expenses.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres
DATABASE_CONN_STRING="postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres" go run ./cmd/pay-yourself-first/main.go &
cd ./frontend
$(npm bin)/cypress run
cd ../
kill $!
docker stop pay-yourself-first-postgres