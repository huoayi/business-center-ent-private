-- modify "merchants" table
ALTER TABLE "merchants" ADD COLUMN "pay_url" character varying NOT NULL DEFAULT '';
-- set comment to column: "pay_url" on table: "merchants"
COMMENT ON COLUMN "merchants" ."pay_url" IS '支付路径';
