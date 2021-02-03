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