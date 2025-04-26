
-- +migrate Up
CREATE TABLE cart_items(
    "id" uuid NOT NULL,
    "cart_id" uuid NOT NULL,
    "product_id" uuid NOT NULL,
    "quantity" INT NOT NULL,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_cart_cart_items" FOREIGN KEY ("cart_id") REFERENCES "carts",
    CONSTRAINT "fk_product_cart_items" FOREIGN KEY ("product_id") REFERENCES "products"
);

-- +migrate Down
DROP TABLE cart_items;
