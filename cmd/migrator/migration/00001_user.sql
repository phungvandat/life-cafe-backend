-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."users"(
	"id" uuid NOT NULL,
	"created_at" timestamptz DEFAULT now(),
  	"deleted_at" timestamptz,
  	"updated_at" timestamptz,
	"username" text NOT NULL,
	"fullname" text NOT NULL,
	"password" text,
	"role" text NOT NULL,
	"active" boolean DEFAULT true,
	"phone_number" text,
	"address" text,
	"email" text,
 	CONSTRAINT "users_pkey" PRIMARY KEY ("id"),
	UNIQUE("username")
)
WITH (oids = false);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE "public"."users";