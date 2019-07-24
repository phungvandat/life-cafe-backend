-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."products"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
  	"deleted_at" timestamptz,
  	"updated_at" timestamptz,
	"name" text NOT NULL,
	"main_photo" text,
	"quantity" int DEFAULT 0,
	"description" text,
	"slug" text NOT NULL,
	"price" text DEFAULT 0,
	"flag" int DEFAULT 1,
	"color" text,
	"barcode" text,
 	CONSTRAINT "products_pkey" PRIMARY KEY ("id"),
	UNIQUE("slug")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."products";