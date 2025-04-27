
-- +migrate Up
CREATE TABLE products(
    "id" uuid NOT NULL,
    "name" text NOT NULL,
    "description" text,
    "price" DECIMAL(10,2) NOT NULL,
    "stock" INT NOT NULL,
    "user_id" uuid NOT NULL,
    "status" text,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id")
);

-- +migrate Down
DROP TABLE products;
