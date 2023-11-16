CREATE TABLE "gitmons" (
  "gitmon_id" varchar PRIMARY KEY,
  "owner_id" varchar NOT NULL,
  "gitmon_name" varchar NOT NULL,
  "exp" int NOT NULL,
  "base_hp" int NOT NULL,
  "current_hp" int NOT NULL,
  "base_attack" int NOT NULL,
  "current_attack" int NOT NULL,
  "base_defence" int NOT NULL,
  "current_defence" int NOT NULL,
  "base_speed" int NOT NULL,
  "current_speed" int NOT NULL
);

CREATE TABLE "gitmon_skills" (
  "gitmon_id" varchar NOT NULL,
  "skill_id" int NOT NULL,
  "is_active" boolean NOT NULL
);


CREATE TABLE "skills" (
  "skill_id" serial PRIMARY KEY,
  "skill_name" varchar NOT NULL,
  "required_bp" int NOT NULL,
  "skilltype" varchar NOT NULL,
  "value" FLOAT NOT NULL,
  "description" text NOT NULL
);

ALTER TABLE "gitmon_skills" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("skill_id");

ALTER TABLE "gitmon_skills" ADD FOREIGN KEY ("gitmon_id") REFERENCES "gitmons" ("gitmon_id");

ALTER TABLE "gitmons" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("user_id");