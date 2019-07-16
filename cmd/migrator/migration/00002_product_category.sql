-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."product_categories"(
	"id" uuid NOT NULL,
	"name" text NOT NULL,
	"photo" text ,
	"parent_category_id" uuid references product_categories,
	"slug" text NOT NULL,
	"color" text DEFAULT NULL,
	UNIQUE("slug"),
	"created_at" timestamptz DEFAULT now(),
  	"deleted_at" timestamptz,
  	"updated_at" timestamptz,
 	CONSTRAINT "product_categories_pkey" PRIMARY KEY ("id"),
	UNIQUE("slug")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."product_categories";