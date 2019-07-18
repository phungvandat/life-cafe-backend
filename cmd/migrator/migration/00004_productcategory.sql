-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."productcategories"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
	"updated_at" timestamptz,
  	"deleted_at" timestamptz,
	"product_id" uuid references products,
	"product_category_id" uuid references product_categories,
 	CONSTRAINT "productcategories_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."productcategories";