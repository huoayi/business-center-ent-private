-- reverse: set comment to column: "user_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."user_id" IS '';
-- reverse: set comment to column: "refresh_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."refresh_token" IS '';
-- reverse: set comment to column: "access_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."access_token" IS '';
-- reverse: set comment to column: "session_key" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."session_key" IS '';
-- reverse: set comment to column: "scope" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."scope" IS '';
-- reverse: set comment to column: "union_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."union_id" IS '';
-- reverse: set comment to column: "open_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."open_id" IS '';
-- reverse: set comment to column: "app_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."app_id" IS '';
-- reverse: set comment to column: "deleted_at" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."id" IS '';
-- reverse: set comment to table: "vx_socials"
COMMENT ON TABLE "vx_socials" IS '';
-- reverse: create index "vxsocial_user_id" to table: "vx_socials"
DROP INDEX "vxsocial_user_id";
-- reverse: create "vx_socials" table
DROP TABLE "vx_socials";
-- reverse: set comment to column: "parent_id" on table: "users"
COMMENT ON COLUMN "users" ."parent_id" IS '';
-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "parent_id";
