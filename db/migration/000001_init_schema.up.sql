CREATE TABLE "customers" (
  "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "address" varchar,
  "phone_number" varchar NOT NULL
);

CREATE TABLE "measures" (
  "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
  "customer_id" bigserial NOT NULL,
  "name" varchar NOT NULL,
  "number" varchar NOT NULL
);

CREATE TABLE "order_items" (
  "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
  "order_id" int UNIQUE NOT NULL,
  "product_id" int UNIQUE NOT NULL,
  "quantity" int NOT NULL DEFAULT 1
);

CREATE TABLE "orders" (
  "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
  "user_id" int NOT NULL,
  "status" varchar,
  "prepaid" bigint,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "products_type" (
  "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL
);

CREATE TABLE "products" (
  "id" SERIAL UNIQUE PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "price" int NOT NULL,
  "type_id" int,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE "measures" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("id") REFERENCES "order_items" ("order_id");

ALTER TABLE "products" ADD FOREIGN KEY ("type_id") REFERENCES "products_type" ("id");

CREATE INDEX ON "customers" ("full_name");

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "products" ("name");

COMMENT ON COLUMN "orders"."created_at" IS 'When order created';

COMMENT ON COLUMN "products"."price" IS 'must be positive';
