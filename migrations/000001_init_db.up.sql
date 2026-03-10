CREATE TABLE IF NOT EXISTS "users" (
    "user_id" SERIAL PRIMARY KEY,
    "full_name" VARCHAR,
    "email" VARCHAR,
    "password" VARCHAR,
    "address" TEXT,
    "phone" VARCHAR,
    "profile_pic" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "kategory" (
    "kategory_id" SERIAL PRIMARY KEY,
    "name" VARCHAR
);

CREATE TABLE IF NOT EXISTS "products" (
    "product_id" SERIAL PRIMARY KEY,
    "kategory_id" INT REFERENCES "kategory"("kategory_id"),
    "name" VARCHAR,
    "description" TEXT,
    "price" INT,
    "image_url" VARCHAR
);

CREATE TABLE IF NOT EXISTS "product_variant" (
    "variant_id" SERIAL PRIMARY KEY,
    "product_id" INT REFERENCES "products"("product_id"),
    "temperature" VARCHAR,
    "add_price" INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "product_size" (
    "product_size_id" SERIAL PRIMARY KEY,
    "product_id" INT REFERENCES "products"("product_id"),
    "name" VARCHAR,
    "add_price" INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS "product_images" (
    "product_images_id" SERIAL PRIMARY KEY,
    "product_id" INT REFERENCES "products"("product_id"),
    "path" VARCHAR
);

CREATE TABLE IF NOT EXISTS "discount" (
    "discount_id" SERIAL PRIMARY KEY,
    "product_id" INT REFERENCES "products"("product_id"),
    "flash_sale" BOOLEAN DEFAULT FALSE,
    "description" TEXT,
    "discount_rate" INT
);

CREATE TABLE IF NOT EXISTS "promo" (
    "promo_id" SERIAL PRIMARY KEY,
    "title" VARCHAR,
    "description" TEXT,
    "promo_type" VARCHAR,
    "discount_value" INT
);

CREATE TABLE IF NOT EXISTS "cart" (
    "cart_id" SERIAL PRIMARY KEY,
    "user_id" INT REFERENCES "users"("user_id")
);

CREATE TABLE IF NOT EXISTS "cart_item" (
    "cart_item_id" SERIAL PRIMARY KEY,
    "cart_id" INT REFERENCES "cart"("cart_id"),
    "product_id" INT REFERENCES "products"("product_id"),
    "variant_id" INT REFERENCES "product_variant"("variant_id"),
    "product_size_id" INT REFERENCES "product_size"("product_size_id"),
    "quantity" INT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS "transaction" (
    "transaction_id" SERIAL PRIMARY KEY,
    "user_id" INT REFERENCES "users"("user_id"),
    "promo_id" INT REFERENCES "promo"("promo_id"),
    "fullname" VARCHAR,
    "email" VARCHAR,
    "address" TEXT,
    "delivery_type" VARCHAR,
    "subtotal" INT,
    "tax" INT,
    "total" INT,
    "tanggal" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transaction_product" (
    "transaction_product_id" SERIAL PRIMARY KEY,
    "transaction_id" INT REFERENCES "transaction"("transaction_id"),
    "product_id" INT REFERENCES "products"("product_id"),
    "variant_id" INT REFERENCES "product_variant"("variant_id"),
    "product_size_id" INT REFERENCES "product_size"("product_size_id"),
    "quantity" INT,
    "price_at_purchase" INT
);

CREATE TABLE IF NOT EXISTS "reviews" (
    "reviews_id" SERIAL PRIMARY KEY,
    "product_id" INT REFERENCES "products"("product_id"),
    "user_id" INT REFERENCES "users"("user_id"),
    "message" TEXT,
    "rating" INT
);