-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."product_categories"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"updated_at" timestamptz,
  	"deleted_at" timestamptz,
	"product_id" uuid references products,
	"category_id" uuid,
 	CONSTRAINT "product_categories_pkey" PRIMARY KEY ("id"),
	UNIQUE("product_id","category_id")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."product_categories";