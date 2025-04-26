
-- +migrate Up
CREATE TABLE order_items(
    "id" uuid NOT NULL,
    "order_id" uuid NOT NULL,
    "product_id" uuid NOT NULL,
    "quantity" INT NOT NULL,
    "price" DECIMAL(10,2),
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_order_item_order" FOREIGN KEY ("order_id") REFERENCES "orders",
    CONSTRAINT "fk_order_item_product" FOREIGN KEY ("product_id") REFERENCES "products"
);

-- +migrate Down
DROP TABLE order_items;
