-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."categories"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
  	"deleted_at" timestamptz,
  	"updated_at" timestamptz,
	"name" text NOT NULL,
	"photo" text ,
	"parent_category_id" uuid references categories,
	"slug" text NOT NULL,
	"color" text DEFAULT NULL,
 	CONSTRAINT "categories_pkey" PRIMARY KEY ("id"),
	UNIQUE("slug")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."categories";