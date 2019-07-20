-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."product_orders"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"updated_at" timestamptz,
  	"deleted_at" timestamptz,
	"product_id" uuid references products,
	"order_id" uuid references orders,
	"order_quantity" int,
	"real_price" int default 0,
 	CONSTRAINT "product_orders_pkey" PRIMARY KEY ("id"),
	UNIQUE("product_id","order_id")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."product_orders";