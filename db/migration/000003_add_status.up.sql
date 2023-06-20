CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL
);

INSERT INTO "roles" ("id", "name") VALUES (0, 'ADMIN');
INSERT INTO "roles" ("id", "name") VALUES (1, 'USER');

ALTER TABLE "users"
ADD COLUMN "status" boolean NOT NULL DEFAULT false;

ALTER TABLE "users"
ADD COLUMN "role" int NOT NULL;

ALTER TABLE "users" ADD FOREIGN KEY ("role") REFERENCES "roles" ("id");

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");