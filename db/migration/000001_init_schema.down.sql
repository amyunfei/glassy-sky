DROP TRIGGER IF EXISTS "update_categories_update_at" ON "categories";
DROP TABLE IF EXISTS "categories";
DROP TRIGGER IF EXISTS "update_labels_update_at" ON "labels";
DROP TABLE IF EXISTS "labels";

DROP TRIGGER IF EXISTS "update_article_on_labels_change" ON "articles_labels";
DROP TABLE IF EXISTS "articles_labels";
DROP TRIGGER IF EXISTS "update_article_on_categories_change" ON "articles_categories";
DROP TABLE IF EXISTS "articles_categories";
DROP FUNCTION IF EXISTS "update_article_on_relation_update";
DROP TRIGGER IF EXISTS "update_articles_update_at" ON "articles";
DROP TABLE IF EXISTS "articles";

DROP TRIGGER IF EXISTS "update_users_update_at" ON "users";
DROP TABLE IF EXISTS "users";
DROP FUNCTION IF EXISTS "update_modified_column";
DROP SEQUENCE IF EXISTS "global_id_sequence";
DROP FUNCTION IF EXISTS "id_generator";