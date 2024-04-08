-- reverse: set comment to column: "produce_type" on table: "products"
COMMENT ON COLUMN "products" ."produce_type" IS '';
-- reverse: modify "products" table
ALTER TABLE "products" DROP COLUMN "produce_type";
-- reverse: set comment to column: "provence" on table: "merchants"
COMMENT ON COLUMN "merchants" ."provence" IS '';
-- reverse: modify "merchants" table
ALTER TABLE "merchants" DROP COLUMN "provence";
