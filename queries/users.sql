CREATE TABLE "bank"."users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password_changed_at" timestamp NOT NULL DEFAULT (now()),
  "created_at" timestamp DEFAULT (now())
);

-- name: CreateUser :one
INSERT INTO "bank"."users" (username, hashed_password, first_name, last_name, email) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUserByUsername :one
select *  from "bank"."users" where username = $1;

-- name: GetUserByEmail :one
select *  from "bank"."users" where email = $1;

-- name: GetUserForUpdate :one
select *  from "bank"."users" where username = $1 FOR NO KEY UPDATE;

-- name: ListAllUsers :many
SELECT * FROM "bank"."users";

-- name: ListUsers :many
SELECT * FROM "bank"."users" ORDER BY username LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE  "bank"."users" SET hashed_password = $1, password_changed_at = $2 WHERE username = $3 RETURNING *;

-- name: DeleteUser :one
DELETE FROM "bank"."users" WHERE username = $1 RETURNING *;


