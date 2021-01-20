<!-- Development -->
psql -d pay_yourself_first -f ./db/migrations/001_budget.up.sql
psql -d pay_yourself_first -f ./db/migrations/002_insert_test_income.up.sql

<!-- Testing -->
psql -d pay_yourself_first_test -f ./db/migrations/001_budget.up.sql

