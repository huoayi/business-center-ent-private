-- reverse: set comment to column: "user_id" on table: "login_records"
COMMENT ON COLUMN "login_records" ."user_id" IS '';
-- reverse: set comment to column: "way" on table: "login_records"
COMMENT ON COLUMN "login_records" ."way" IS '';
-- reverse: set comment to column: "ip" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ip" IS '';
-- reverse: set comment to column: "ua" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ua" IS '';
-- reverse: set comment to column: "deleted_at" on table: "login_records"
COMMENT ON COLUMN "login_records" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "login_records"
COMMENT ON COLUMN "login_records" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "login_records"
COMMENT ON COLUMN "login_records" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "login_records"
COMMENT ON COLUMN "login_records" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "login_records"
COMMENT ON COLUMN "login_records" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "login_records"
COMMENT ON COLUMN "login_records" ."id" IS '';
-- reverse: set comment to table: "login_records"
COMMENT ON TABLE "login_records" IS '';
-- reverse: create index "loginrecord_user_id" to table: "login_records"
DROP INDEX "loginrecord_user_id";
-- reverse: create "login_records" table
DROP TABLE "login_records";
