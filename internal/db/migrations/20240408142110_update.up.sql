-- modify "merchants" table
ALTER TABLE "merchants" ADD COLUMN "provence" character varying NOT NULL DEFAULT '河南';
-- set comment to column: "provence" on table: "merchants"
COMMENT ON COLUMN "merchants" ."provence" IS '省份';
-- modify "products" table
ALTER TABLE "products" ADD COLUMN "produce_type" character varying NOT NULL DEFAULT 'tea';
-- set comment to column: "produce_type" on table: "products"
COMMENT ON COLUMN "products" ."produce_type" IS '产品类型';
