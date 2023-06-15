ALTER TABLE "blogs"
ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT (now());

ALTER TABLE "reviews"
ADD COLUMN "updated_at" timestamptz NOT NULL DEFAULT (now());

ALTER TABLE "reviews"
ADD COLUMN "created_at" timestamptz NOT NULL DEFAULT (now());

ALTER TABLE "reviews"
ADD COLUMN "p_id" int NOT NULL;

ALTER TABLE "reviews"
ADD FOREIGN KEY ("p_id") REFERENCES "packages" ("id");