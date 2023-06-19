CREATE TABLE "roles" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL
);

ALTER TABLE "users"
ADD COLUMN "status" boolean NOT NULL DEFAULT false;

ALTER TABLE "users"
ADD COLUMN "role" int NOT NULL;

ALTER TABLE "users" ADD FOREIGN KEY ("role") REFERENCES "roles" ("id");