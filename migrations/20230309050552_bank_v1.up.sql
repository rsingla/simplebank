DROP SCHEMA IF EXISTS "bank" CASCADE;

CREATE SCHEMA "bank";

CREATE TYPE "bank"."Currency" AS ENUM (
  'USD',
  'EUR'
);

CREATE TABLE "bank"."accounts" (
  "account_id" BIGSERIAL PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "bank"."entries" (
  "entry_id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "bank"."transfers" (
  "entry_id" BIGSERIAL PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "bank"."accounts" ("owner");

CREATE INDEX ON "bank"."transfers" ("from_account_id");

CREATE INDEX ON "bank"."transfers" ("to_account_id");

CREATE INDEX ON "bank"."transfers" ("from_account_id", "to_account_id");

ALTER TABLE "bank"."entries" ADD FOREIGN KEY ("account_id") REFERENCES "bank"."accounts" ("account_id");

ALTER TABLE "bank"."transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "bank"."accounts" ("account_id");

ALTER TABLE "bank"."transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "bank"."accounts" ("account_id");
