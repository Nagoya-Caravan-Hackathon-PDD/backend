ALTER TABLE "encounters" DROP CONSTRAINT "encounters_to_user_id_fkey";
ALTER TABLE "encounters" DROP CONSTRAINT "encounters_from_user_id_fkey";
DROP TABLE IF EXISTS "encounters";
