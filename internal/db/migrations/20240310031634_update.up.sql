-- modify "users" table
ALTER TABLE "users" ADD COLUMN "parent_id" bigint NULL DEFAULT 0;
-- set comment to column: "parent_id" on table: "users"
COMMENT ON COLUMN "users" ."parent_id" IS '邀请人用户 id';
-- create "vx_socials" table
CREATE TABLE "vx_socials" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "app_id" character varying NOT NULL DEFAULT '', "open_id" character varying NOT NULL DEFAULT '', "union_id" character varying NOT NULL DEFAULT '', "scope" character varying NOT NULL DEFAULT 'base', "session_key" character varying NOT NULL DEFAULT '', "access_token" character varying NOT NULL DEFAULT '', "refresh_token" character varying NOT NULL DEFAULT '', "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- create index "vxsocial_user_id" to table: "vx_socials"
CREATE INDEX "vxsocial_user_id" ON "vx_socials" ("user_id");
-- set comment to table: "vx_socials"
COMMENT ON TABLE "vx_socials" IS '微信社会源信息，记录用户在微信方的身份信息';
-- set comment to column: "id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "app_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."app_id" IS '微信应用 id';
-- set comment to column: "open_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."open_id" IS '微信身份源的 open_id';
-- set comment to column: "union_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."union_id" IS '微信身份源的 union_id';
-- set comment to column: "scope" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."scope" IS '账户的权限级别，一般为 base';
-- set comment to column: "session_key" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."session_key" IS '小程序专用的会话密钥，不可以返回给前端';
-- set comment to column: "access_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."access_token" IS '微信能力访问凭证';
-- set comment to column: "refresh_token" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."refresh_token" IS '刷新微信凭证的刷新凭证';
-- set comment to column: "user_id" on table: "vx_socials"
COMMENT ON COLUMN "vx_socials" ."user_id" IS '外键用户 id';
