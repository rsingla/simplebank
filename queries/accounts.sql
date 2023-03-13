CREATE TABLE "bank"."accounts" (
  "account_id" BIGSERIAL PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);


-- name: GetAccountByOwnerName :many
select *  from "bank"."accounts" where owner = $1;

-- name: GetAccountByAccountId :one
select *  from "bank"."accounts" where account_id = $1;


-- name: GetAccountForUpdate :one
select *  from "bank"."accounts" where account_id = $1 FOR NO KEY UPDATE;

-- name: ListAllAccounts :many
SELECT * FROM "bank"."accounts";

-- name: ListAccounts :many
SELECT * FROM "bank"."accounts" ORDER BY account_id LIMIT $1 OFFSET $2;

-- name: CreateAccount :one
INSERT INTO "bank"."accounts" (owner, balance, currency) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAccount :one
UPDATE  "bank"."accounts" SET balance = $1 WHERE account_id = $2 RETURNING *;

-- name: DeleteAccount :one
DELETE FROM "bank"."accounts" WHERE account_id = $1 RETURNING *;



