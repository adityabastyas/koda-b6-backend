CREATE TABLE IF NOT EXISTS "forgot_password"(
  "id" SERIAL PRIMARY KEY,
  "email" VARCHAR(255) UNIQUE NOT NULL,
  "code" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP DEFAULT now() NOT NULL,
  "update_at" TIMESTAMP,
  "delete_at" TIMESTAMP
);