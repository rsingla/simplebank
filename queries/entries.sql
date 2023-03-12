
CREATE TABLE "bank"."entries" (
  "entry_id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamp DEFAULT (now())
);


-- name: CreateEntry :one
INSERT INTO "bank"."entries" (account_id, amount) VALUES ($1, $2) RETURNING *;

-- name: GetEntryByEntryId :one
select account_id, amount, created_at  from "bank"."entries" where entry_id = 1;




