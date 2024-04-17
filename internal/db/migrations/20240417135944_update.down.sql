-- reverse: set comment to column: "count" on table: "products"
COMMENT ON COLUMN "products" ."count" IS '';
-- reverse: modify "products" table
ALTER TABLE "products" DROP COLUMN "count";
-- reverse: set comment to column: "status" on table: "orders"
COMMENT ON COLUMN "orders" ."status" IS '';
-- reverse: modify "orders" table
ALTER TABLE "orders" DROP COLUMN "status";
