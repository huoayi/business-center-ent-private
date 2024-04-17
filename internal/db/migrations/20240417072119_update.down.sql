-- reverse: set comment to column: "user_id" on table: "orders"
COMMENT ON COLUMN "orders" ."user_id" IS '';
-- reverse: set comment to column: "products_id" on table: "orders"
COMMENT ON COLUMN "orders" ."products_id" IS '';
-- reverse: set comment to column: "address" on table: "orders"
COMMENT ON COLUMN "orders" ."address" IS '';
-- reverse: set comment to column: "amount" on table: "orders"
COMMENT ON COLUMN "orders" ."amount" IS '';
-- reverse: set comment to column: "count" on table: "orders"
COMMENT ON COLUMN "orders" ."count" IS '';
-- reverse: set comment to column: "deleted_at" on table: "orders"
COMMENT ON COLUMN "orders" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "orders"
COMMENT ON COLUMN "orders" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "orders"
COMMENT ON COLUMN "orders" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "orders"
COMMENT ON COLUMN "orders" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "orders"
COMMENT ON COLUMN "orders" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "orders"
COMMENT ON COLUMN "orders" ."id" IS '';
-- reverse: create "orders" table
DROP TABLE "orders";
-- reverse: create index "products_business_id_key" to table: "products"
DROP INDEX "products_business_id_key";
