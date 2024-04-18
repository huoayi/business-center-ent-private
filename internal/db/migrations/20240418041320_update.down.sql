-- reverse: set comment to column: "pay_url" on table: "merchants"
COMMENT ON COLUMN "merchants" ."pay_url" IS '';
-- reverse: modify "merchants" table
ALTER TABLE "merchants" DROP COLUMN "pay_url";
