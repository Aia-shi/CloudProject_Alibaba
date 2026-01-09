-- name: GetUserPeriods :many
SELECT * FROM periods WHERE periods.user_id = $1;

-- name: GetUserExpenses :many
SELECT * FROM expenses WHERE expenses.user_id = $1;

-- name: GetUserIncome :many
SELECT * FROM incomes WHERE incomes.user_id = $1;

-- name: GetUserPeriodsById :one
SELECT * FROM periods WHERE periods.id = $1;

-- name: GetUserExpensesById :one
SELECT * FROM expenses WHERE expenses.id = $1;

-- name: GetUserIncomeById :one
SELECT * FROM incomes WHERE incomes.id = $1;

-- name: DeleteUserPeriods :exec
DELETE FROM periods WHERE periods.id = $1;

-- name: DeleteUserExpenses :exec
DELETE FROM expenses WHERE expenses.id = $1;

-- name: DeleteUserIncomes :exec
DELETE FROM incomes WHERE incomes.id = $1;

-- name: DeleteUserExpensesWithPeriod :exec
DELETE FROM expenses WHERE expenses.period_id = $1;

-- name: DeleteUserIncomesWithPeriod :exec
DELETE FROM incomes WHERE incomes.period_id = $1;

-- name: UpdateUserPeriods :exec
UPDATE periods SET
  name = @new_name::text
WHERE periods.id = @id::int;

-- name: UpdateUserExpenses :exec
UPDATE expenses SET
  title = @new_title::text,
  amount = @new_amount::float,
  description = @new_description::text,
  date = @new_date::text,
  status = @new_status::text,
  category = @new_category::text
WHERE expenses.id = @id::int;

-- name: UpdateUserIncomes :exec
UPDATE incomes SET
  title = @new_title::text,
  amount = @new_amount::float,
  description = @new_description::text,
  date = @new_date::text,
  category = @new_category::text
WHERE incomes.id = @id::int;

-- name: InsertUserPeriods :exec
INSERT INTO periods (user_id, name) VALUES (
    $1,
    $2
  );

-- name: InsertUserExpenses :exec
INSERT INTO expenses (period_id, user_id, title, amount, description, date, status, category) VALUES (
    $1,
    $2,  
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
  );

-- name: InsertUserIncomes :exec
INSERT INTO incomes (period_id, user_id, title, amount, description, date, category) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
  );
