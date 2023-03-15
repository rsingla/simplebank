CREATE TABLE "bank"."users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT (now()),
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "bank"."users" ("email");

ALTER TABLE "bank"."accounts" ADD FOREIGN KEY ("owner") REFERENCES "bank"."users" ("username");

ALTER TABLE "bank"."users" ADD CONSTRAINT "email_unique" UNIQUE ("email");

ALTER TABLE "bank"."accounts" ADD CONSTRAINT "owner_currency_unique" UNIQUE ("owner", "currency");