CREATE TABLE "customers" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar,
  "created_at" timestamp DEFAULT (now()),
  "address" varchar,
  "phone_number" varchar
);

CREATE TABLE "measures" (
  "code" int PRIMARY KEY,
  "customer_id" bigserial,
  "name" varchar,
  "number" varchar
);

CREATE TABLE "order_items" (
  "order_id" int,
  "product_id" int,
  "quantity" int DEFAULT 1
);

CREATE TABLE "orders" (
  "id" int PRIMARY KEY,
  "user_id" int UNIQUE NOT NULL,
  "status" varchar,
  "prepaid" bigint,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "products_type" (
  "id" int PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "products" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "price" int NOT NULL,
  "type_id" int,
  "created_at" timestamp DEFAULT (now())
);

ALTER TABLE "measures" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "customers" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("type_id") REFERENCES "products_type" ("id");

CREATE INDEX ON "customers" ("full_name");

CREATE INDEX ON "orders" ("user_id");

CREATE INDEX ON "products" ("name");

COMMENT ON COLUMN "orders"."created_at" IS 'When order created';

COMMENT ON COLUMN "products"."price" IS 'must be positive';
