CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_at = now();
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