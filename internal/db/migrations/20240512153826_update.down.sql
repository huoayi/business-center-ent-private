-- reverse: set comment to column: "merchant_id" on table: "products"
COMMENT ON COLUMN "products" ."merchant_id" IS '';
-- reverse: modify "products" table
ALTER TABLE "products" DROP COLUMN "merchant_id";
-- reverse: drop index "products_business_id_key" from table: "products"
CREATE UNIQUE INDEX "products_business_id_key" ON "products" ("business_id");
