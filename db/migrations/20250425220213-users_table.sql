-- +migrate Up
CREATE TABLE users(
    "id" uuid NOT NULL,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "name" text NOT NULL,
    "role_id" uuid NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz,
    PRIMARY KEY("id"),
    CONSTRAINT "fk_role" FOREIGN KEY ("role_id") REFERENCES "role"
); 

CREATE UNIQUE INDEX idx_users_username ON ssers(username);
CREATE UNIQUE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_created_at ON users(created_at);

-- +migrate Down
DROP TABLE user;