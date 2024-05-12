-- drop index "products_business_id_key" from table: "products"
DROP INDEX "products_business_id_key";
-- modify "products" table
ALTER TABLE "products" ADD COLUMN "merchant_id" bigint NOT NULL DEFAULT 0;
-- set comment to column: "merchant_id" on table: "products"
COMMENT ON COLUMN "products" ."merchant_id" IS '外键商户用户 id';
