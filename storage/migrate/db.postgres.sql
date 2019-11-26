DROP SCHEMA public CASCADE;

DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;


CREATE SCHEMA public;


CREATE EXTENSION "uuid-ossp";

---------------------------------------------------------------------------------------------------
----------------------------------------- Creates tables views -----------------------------
---------------------------------------------------------------------------------------------------

CREATE TABLE review
(
    id        UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    answers   JSONB                     DEFAULT '[]',
    author    VARCHAR(256)              DEFAULT '',
    body      VARCHAR                   DEFAULT '',
    orderHash VARCHAR(256)              DEFAULT '',
    rated     VARCHAR(256)              DEFAULT '',
    rating    INTEGER                   DEFAULT 0,
    created   TIMESTAMP                 DEFAULT now(),
    updated   TIMESTAMP                 DEFAULT now()
);
