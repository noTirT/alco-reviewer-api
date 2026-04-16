CREATE
OR REPLACE FUNCTION trigger_set_timestamp() RETURNS trigger AS
$$
BEGIN
NEW.updated_at = NOW();

RETURN NEW;

END;

$$
language plpgsql;

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT gen_random_uuid(),
    email varchar NOT NULL UNIQUE,
    username varchar NOT NULL UNIQUE,
    PASSWORD varchar NOT NULL UNIQUE,
    follower_count int NOT NULL DEFAULT 0,
    following_count int NOT NULL DEFAULT 0,
    tokenhash varchar(15) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS drinks (
    id uuid DEFAULT gen_random_uuid(),
    name varchar NOT NULL,
    alcohol boolean NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS locations (
    id uuid DEFAULT gen_random_uuid(),
    name varchar NOT NULL,
    TYPE varchar NOT NULL,
    address varchar NOT NULL,
    city varchar NOT NULL,
    zip_code varchar NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS reviews (
    id uuid DEFAULT gen_random_uuid(),
    reviewer_id uuid NOT NULL,
    rating int NOT NULL,
    review_text text,
    drink_id uuid NOT NULL,
    location_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id),
    FOREIGN KEY(reviewer_id) REFERENCES users(id),
    FOREIGN KEY(drink_id) REFERENCES drinks(id),
    FOREIGN KEY(location_id) REFERENCES locations(id),
    CONSTRAINT rating_range CHECK (
        rating >= 1
        AND rating <= 5
    )
);

CREATE TABLE IF NOT EXISTS drinks_to_locations(
    drink_id uuid NOT NULL,
    location_id uuid NOT NULL,
    PRIMARY KEY(drink_id, location_id),
    FOREIGN KEY(drink_id) REFERENCES drinks(id),
    FOREIGN KEY(location_id) REFERENCES locations(id)
);

CREATE TABLE IF NOT EXISTS following(
    follower_id uuid NOT NULL,
    followed_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(follower_id, followed_id),
    FOREIGN KEY(follower_id) REFERENCES users(id),
    FOREIGN KEY(followed_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS following_location(
    follower_id uuid NOT NULL,
    location_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY(follower_id, location_id),
    FOREIGN KEY(follower_id) REFERENCES users(id),
    FOREIGN KEY(location_id) REFERENCES locations(id)
);

CREATE TRIGGER set_timestamp_user BEFORE
UPDATE
    ON users FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_review BEFORE
UPDATE
    ON reviews FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_drinks BEFORE
UPDATE
    ON drinks FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_locations BEFORE
UPDATE
    ON locations FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();
