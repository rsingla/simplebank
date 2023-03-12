CREATE TABLE "bank"."transfers" (
  "entry_id" BIGSERIAL PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamp DEFAULT (now())
);


-- name: CreateTransfer :one
INSERT INTO "bank"."transfers" (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTransferByEntryId :one
select from_account_id, to_account_id, amount, created_at  from "bank"."transfers" where entry_id = 1;
