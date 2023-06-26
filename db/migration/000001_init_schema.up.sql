CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "destinations" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "price" int NOT NULL,
  "country" varchar NOT NULL,
  "visa_require" varchar,
  "language" varchar,
  "currency" varchar,
  "area" varchar,
  "location" varchar
);

CREATE TABLE "galleries" (
  "id" BIGSERIAL PRIMARY KEY,
  "des_id" int NOT NULL,
  "url" varchar NOT NULL
);

CREATE TABLE "types" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "packages" (
  "id" SERIAL PRIMARY KEY,
  "des_id" int NOT NULL,
  "name" varchar NOT NULL,
  "details" varchar NOT NULL,
  "price" bigint NOT NULL,
  "img_url" varchar NOT NULL,
  "duration" varchar NOT NULL,
  "number_people" int NOT NULL
);

CREATE TABLE "packages_types" (
  "p_id" int,
  "t_id" int,
  PRIMARY KEY ("p_id", "t_id")
);

CREATE TABLE "trips" (
  "id" BIGSERIAL PRIMARY KEY,
  "user" varchar NOT NULL,
  "p_id" int NOT NULL,
  "total" bigint NOT NULL,
  "start_date" timestamptz NOT NULL DEFAULT (now()),
  "notes" varchar
);

CREATE TABLE "blogs" (
  "id" SERIAL PRIMARY KEY,
  "author" varchar NOT NULL,
  "destination_id" int,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "details" varchar NOT NULL,
  "tags" varchar,
  "facebook" varchar,
  "instagram" varchar,
  "twitter" varchar,
  "img_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reviews" (
  "id" SERIAL PRIMARY KEY,
  "author" varchar NOT NULL,
  "rate" float NOT NULL,
  "detail" varchar
);

CREATE INDEX ON "destinations" ("id");

CREATE INDEX ON "destinations" ("name");

CREATE INDEX ON "galleries" ("des_id");

CREATE INDEX ON "packages" ("id");

CREATE INDEX ON "packages" ("name");

CREATE UNIQUE INDEX ON "packages_types" ("p_id", "t_id");

ALTER TABLE "galleries" ADD FOREIGN KEY ("des_id") REFERENCES "destinations" ("id");

ALTER TABLE "packages" ADD FOREIGN KEY ("des_id") REFERENCES "destinations" ("id");

ALTER TABLE "packages_types" ADD FOREIGN KEY ("p_id") REFERENCES "packages" ("id");

ALTER TABLE "packages_types" ADD FOREIGN KEY ("t_id") REFERENCES "types" ("id");

ALTER TABLE "trips" ADD FOREIGN KEY ("user") REFERENCES "users" ("username");

ALTER TABLE "trips" ADD FOREIGN KEY ("p_id") REFERENCES "packages" ("id");

ALTER TABLE "blogs" ADD FOREIGN KEY ("author") REFERENCES "users" ("username");

ALTER TABLE "blogs" ADD FOREIGN KEY ("destination_id") REFERENCES "destinations" ("id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("author") REFERENCES "users" ("username");
