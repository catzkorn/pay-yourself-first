<!-- Development -->

psql -d pay_yourself_first -f ./db/migrations/001_budget.up.sql
psql -d pay_yourself_first -f ./db/migrations/002_insert_test_income.up.sql
psql -d pay_yourself_first -f ./db/migrations/003_savings.up.sql
psql -d pay_yourself_first -f ./db/migrations/004_drop_income_id_add_date.up.sql
psql -d pay_yourself_first -f ./db/migrations/005_expenses.up.sql

<!-- Testing -->

psql -d pay_yourself_first_test -f ./db/migrations/001_budget.up.sql
psql -d pay_yourself_first_test -f ./db/migrations/003_savings.up.sql
psql -d pay_yourself_first_test -f ./db/migrations/004_drop_income_id_add_date.up.sql
psql -d pay_yourself_first_test -f ./db/migrations/005_expenses.up.sql

<!-- Docker PSQL -->

docker run --name pay-yourself-first-postgres -p 6969:5432 -e POSTGRES_PASSWORD=mysecretpassword --rm -d postgres
psql -f ./db/migrations/001_budget.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres
psql -f ./db/migrations/003_savings.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres
psql -f ./db/migrations/004_drop_income_id_add_date.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres
psql -f ./db/migrations/005_expenses.up.sql postgresql://postgres:mysecretpassword@0.0.0.0:6969/postgres

Generating React:

```shell
esbuild --bundle ./frontend/src/app.jsx --outfile=web/js/index.js --define:process.env.NODE_ENV=\"development\" --minify --watch --sourcemap
esbuild --bundle ./frontend/src/app.jsx --outfile=web/js/index.js --define:process.env.NODE_ENV=\"production\" --minify --watch --sourcemap
```

Typescript typechecker

```shell
npx tsc
```

Open Cypress

```shell
$(npm bin)/cypress open
```
