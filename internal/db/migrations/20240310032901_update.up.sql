-- create "login_records" table
CREATE TABLE "login_records" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "ua" character varying NOT NULL DEFAULT '', "ip" character varying NOT NULL DEFAULT '', "way" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "loginrecord_user_id" to table: "login_records"
CREATE INDEX "loginrecord_user_id" ON "login_records" ("user_id");
-- set comment to table: "login_records"
COMMENT ON TABLE "login_records" IS '登录记录，记录用户登录的一些信息';
-- set comment to column: "id" on table: "login_records"
COMMENT ON COLUMN "login_records" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "login_records"
COMMENT ON COLUMN "login_records" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "login_records"
COMMENT ON COLUMN "login_records" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "login_records"
COMMENT ON COLUMN "login_records" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "login_records"
COMMENT ON COLUMN "login_records" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "login_records"
COMMENT ON COLUMN "login_records" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "ua" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ua" IS '用户登录时的 user-agent';
-- set comment to column: "ip" on table: "login_records"
COMMENT ON COLUMN "login_records" ."ip" IS '用户登录时的 ip 地址';
-- set comment to column: "way" on table: "login_records"
COMMENT ON COLUMN "login_records" ."way" IS '用户登录时的登录方式，比如手机号验证码等';
-- set comment to column: "user_id" on table: "login_records"
COMMENT ON COLUMN "login_records" ."user_id" IS '用户 id，外键关联用户';
