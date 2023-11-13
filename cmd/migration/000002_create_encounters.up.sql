CREATE TABLE "encounters" (
  "encounter_id" varchar PRIMARY KEY,
  "to_user_id" varchar NOT NULL,
  "from_user_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL
);

ALTER TABLE "encounters" ADD FOREIGN KEY ("to_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "encounters" ADD FOREIGN KEY ("from_user_id") REFERENCES "users" ("user_id");
