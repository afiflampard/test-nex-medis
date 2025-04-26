
-- +migrate Up
CREATE TABLE orders(
    "id" uuid NOT NULL,
    "customer_id" uuid NOT NULL,
    "order_date" timestamptz NOT NULL,
    "cart_id" uuid NOT NULL,
    "status" VARCHAR(50) NOT NULL CHECK (status IN ('pending', 'paid', 'shipped', 'completed', 'cancelled')),
    "total_amount" DECIMAL(10,2) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_orders_user" FOREIGN KEY ("customer_id") REFERENCES "users"
);

CREATE INDEX idx_orders_order_date ON orders(order_date);
CREATE INDEX idx_orders_customer_order_date ON orders(customer_id, order_date);


-- +migrate Down
DROP TABLE orders;
