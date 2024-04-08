-- create "merchants" table
CREATE TABLE "merchants" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "merchant_name" character varying NOT NULL DEFAULT '默认商户', "jpg_url" character varying NOT NULL DEFAULT '', "comment" character varying NOT NULL DEFAULT '', "amount" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "merchants"
COMMENT ON COLUMN "merchants" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "merchants"
COMMENT ON COLUMN "merchants" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "merchants"
COMMENT ON COLUMN "merchants" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "merchants"
COMMENT ON COLUMN "merchants" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "merchants"
COMMENT ON COLUMN "merchants" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "merchants"
COMMENT ON COLUMN "merchants" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "merchant_name" on table: "merchants"
COMMENT ON COLUMN "merchants" ."merchant_name" IS '商户姓名';
-- set comment to column: "jpg_url" on table: "merchants"
COMMENT ON COLUMN "merchants" ."jpg_url" IS '商户头像';
-- set comment to column: "comment" on table: "merchants"
COMMENT ON COLUMN "merchants" ."comment" IS '商户介绍';
-- set comment to column: "amount" on table: "merchants"
COMMENT ON COLUMN "merchants" ."amount" IS '上架商品总数';
-- set comment to column: "user_id" on table: "merchants"
COMMENT ON COLUMN "merchants" ."user_id" IS '外键用户 id';
-- create "products" table
CREATE TABLE "products" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "product_name" character varying NOT NULL DEFAULT '未知产品', "jpg_url" character varying NOT NULL DEFAULT '', "comment" character varying NOT NULL DEFAULT '', "price" bigint NOT NULL DEFAULT 0, "unit" character varying NOT NULL DEFAULT '', "business_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "products"
COMMENT ON COLUMN "products" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "products"
COMMENT ON COLUMN "products" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "products"
COMMENT ON COLUMN "products" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "products"
COMMENT ON COLUMN "products" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "products"
COMMENT ON COLUMN "products" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "products"
COMMENT ON COLUMN "products" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "product_name" on table: "products"
COMMENT ON COLUMN "products" ."product_name" IS '产品名称';
-- set comment to column: "jpg_url" on table: "products"
COMMENT ON COLUMN "products" ."jpg_url" IS '产品照片';
-- set comment to column: "comment" on table: "products"
COMMENT ON COLUMN "products" ."comment" IS '产品介绍';
-- set comment to column: "price" on table: "products"
COMMENT ON COLUMN "products" ."price" IS '单价';
-- set comment to column: "unit" on table: "products"
COMMENT ON COLUMN "products" ."unit" IS '单价使用单位';
-- set comment to column: "business_id" on table: "products"
COMMENT ON COLUMN "products" ."business_id" IS '外键商户用户 id';
