
-- +migrate Up
CREATE TABLE carts(
    "id" uuid NOT NULL,
    "user_id" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_carts_user" FOREIGN KEY ("user_id") REFERENCES "users"
);

-- +migrate Down
DROP TABLE carts;