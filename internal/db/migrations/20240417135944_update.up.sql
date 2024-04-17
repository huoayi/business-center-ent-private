-- modify "orders" table
ALTER TABLE "orders" ADD COLUMN "status" character varying NOT NULL DEFAULT 'canceled';
-- set comment to column: "status" on table: "orders"
COMMENT ON COLUMN "orders" ."status" IS '订单';
-- modify "products" table
ALTER TABLE "products" ADD COLUMN "count" bigint NOT NULL DEFAULT 0;
-- set comment to column: "count" on table: "products"
COMMENT ON COLUMN "products" ."count" IS '库存';
