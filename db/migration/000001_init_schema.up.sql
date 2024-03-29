CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE SEQUENCE global_id_sequence;
CREATE OR REPLACE FUNCTION id_generator(OUT result bigint) AS $$
DECLARE
    our_epoch bigint := 1314220021721;
    seq_id bigint;
    now_millis bigint;
    -- the id of this DB shard, must be set for each
    -- schema shard you have - you could pass this as a parameter too
    shard_id int := 1;
BEGIN
    SELECT nextval('global_id_sequence') % 1024 INTO seq_id;

    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
    result := (now_millis - our_epoch) << 23;
    result := result | (shard_id << 10);
    result := result | (seq_id);
END;
$$ LANGUAGE 'plpgsql';

CREATE TABLE "categories" (
  "id" bigint PRIMARY KEY NOT NULL DEFAULT (id_generator()),
  "name" varchar(100) NOT NULL,
  "parent_id" bigint NOT NULL DEFAULT 0 CHECK (parent_id <> id),
  "color" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);
CREATE TRIGGER update_categories_update_at BEFORE UPDATE ON categories FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE "labels" (
  "id" bigint PRIMARY KEY NOT NULL DEFAULT (id_generator()),
  "name" varchar(100) NOT NULL,
  "color" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);
CREATE TRIGGER update_labels_update_at BEFORE UPDATE ON labels FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE "users" (
  "id" bigint PRIMARY KEY NOT NULL DEFAULT (id_generator()),
  "username" varchar(100) UNIQUE NOT NULL,
  "password" varchar(60) NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "nickname" varchar(100) NOT NULL,
  "avatar" varchar(255),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);
CREATE TRIGGER update_users_update_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE "articles" (
  "id" bigint PRIMARY KEY NOT NULL DEFAULT (id_generator()),
  "title" varchar(255) NOT NULL,
  "excerpt" text NOT NULL,
  "content" text NOT NULL,
  "user_id" bigint NOT NULL REFERENCES "users" (id),
  "status" varchar(20) NOT NULL DEFAULT 'draft',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);
COMMENT ON COLUMN articles.status IS 'draft, published, archived';
CREATE TRIGGER update_articles_update_at BEFORE UPDATE ON articles FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE OR REPLACE FUNCTION update_article_on_relation_update()
RETURNS TRIGGER AS $$
BEGIN
  IF (TG_OP = 'INSERT') THEN
    UPDATE articles SET updated_at = now() WHERE id = NEW.article_id;
    RETURN NEW;
  ELSIF (TG_OP = 'DELETE') THEN
    UPDATE articles SET updated_at = now() WHERE id = OLD.article_id;
    RETURN OLD;
  END IF;
END;
$$ language 'plpgsql';

CREATE TABLE "articles_categories" (
  "article_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  PRIMARY KEY ("article_id", "category_id")
);
CREATE TRIGGER update_article_on_categories_change
AFTER INSERT OR DELETE ON articles_categories
FOR EACH ROW EXECUTE PROCEDURE update_article_on_relation_update();

CREATE TABLE "articles_labels" (
  "article_id" bigint NOT NULL,
  "label_id" bigint NOT NULL,
  PRIMARY KEY ("article_id", "label_id")
);
CREATE TRIGGER update_article_on_labels_change
AFTER INSERT OR DELETE ON articles_labels
FOR EACH ROW EXECUTE PROCEDURE update_article_on_relation_update();

ALTER TABLE "articles" ADD CONSTRAINT "fk_articles_user_id_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;