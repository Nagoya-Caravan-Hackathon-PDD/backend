ALTER TABLE "gitmons" DROP CONSTRAINT "gitmons_owner_id_fkey";
ALTER TABLE "gitmon_skills" DROP CONSTRAINT "gitmon_skills_gitmon_id_fkey";
ALTER TABLE "gitmon_skills" DROP CONSTRAINT "gitmon_skills_skill_id_fkey";

DROP TABLE "skills";
DROP TABLE "gitmon_skills";
DROP TABLE "gitmons";
