-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."orders"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
  	"deleted_at" timestamptz,
  	"updated_at" timestamptz,
	"type" text,
	"creator_id" uuid references users,
	"note" text,
	"customer_id" uuid references users,
	"status" text,
	"implementer_id" uuid references users,
	"receiver_phone_number" text,
	"receiver_address" text NOT NULL,
	"receiver_fullname" text DEFAULT 0,
 	CONSTRAINT "orders_pkey" PRIMARY KEY ("id")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."orders";