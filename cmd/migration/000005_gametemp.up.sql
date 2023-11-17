CREATE TABLE "game_temp" (
  "game_id" varchar PRIMARY KEY,
  "owner_id" varchar NOT NULL,
  "owner_is_ok" bool NOT NULL,
  "enemy_id" varchar,
  "enemy_is_ok" bool NOT NULL,
  "is_block" bool NOT NULL,
  "owner_speed" int NOT NULL,
  "is_first" bool NOT NULL,
  "game_index" int NOT NULL,
  "is_end" bool NOT NULL
);
