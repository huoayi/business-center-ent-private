-- create index "products_business_id_key" to table: "products"
CREATE UNIQUE INDEX "products_business_id_key" ON "products" ("business_id");
-- create "orders" table
CREATE TABLE "orders" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "count" bigint NOT NULL DEFAULT 0, "amount" bigint NOT NULL DEFAULT 0, "address" character varying NOT NULL DEFAULT '', "products_id" bigint NOT NULL DEFAULT 0, "user_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "orders"
COMMENT ON COLUMN "orders" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "orders"
COMMENT ON COLUMN "orders" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "orders"
COMMENT ON COLUMN "orders" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "orders"
COMMENT ON COLUMN "orders" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "orders"
COMMENT ON COLUMN "orders" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "orders"
COMMENT ON COLUMN "orders" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "count" on table: "orders"
COMMENT ON COLUMN "orders" ."count" IS '数量，按照商品计量单位计算';
-- set comment to column: "amount" on table: "orders"
COMMENT ON COLUMN "orders" ."amount" IS '总价';
-- set comment to column: "address" on table: "orders"
COMMENT ON COLUMN "orders" ."address" IS '收货地址';
-- set comment to column: "products_id" on table: "orders"
COMMENT ON COLUMN "orders" ."products_id" IS '产品 id';
-- set comment to column: "user_id" on table: "orders"
COMMENT ON COLUMN "orders" ."user_id" IS '用户 id';
