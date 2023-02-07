DROP TRIGGER IF EXISTS "update_categories_update_at" ON "categories";
DROP TABLE IF EXISTS "categories";

DROP FUNCTION IF EXISTS "update_modified_column";
DROP SEQUENCE IF EXISTS "global_id_sequence";
DROP FUNCTION IF EXISTS "id_generator";