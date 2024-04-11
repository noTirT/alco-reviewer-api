create or replace function trigger_set_timestamp()
returns trigger
as $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
language plpgsql
;

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT gen_random_uuid(),
    email varchar not null unique,
    username varchar not null unique,
    password varchar not null unique,
    tokenhash varchar(15) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    primary key(id)
);

CREATE TABLE IF NOT EXISTS drinks (
    id uuid DEFAULT gen_random_uuid(),
    name varchar not null,
    alcohol boolean not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    primary key(id)
    );

CREATE TABLE IF NOT EXISTS locations (
    id uuid DEFAULT gen_random_uuid(),
    name varchar not null,
    type varchar not null,
    address varchar not null,
    city varchar not null,
    zip_code varchar not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    primary key(id)
    );

CREATE TABLE IF NOT EXISTS reviews (
        id uuid DEFAULT gen_random_uuid(),
    reviewer_id int not null,
    rating int not null,
    review_text text,
    drink_id int not null,
    location_id int not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    primary key(id),
    foreign key(reviewer_id) references users(id),
    foreign key(drink_id) references drinks(id),
    foreign key(location_id) references locations(id),
    constraint rating_range check (rating >= 1 and rating <= 5)
);


CREATE TABLE IF NOT EXISTS drinks_to_locations(
    drink_id int not null,
    location_id int not null,
    primary key(drink_id, location_id),
    foreign key(drink_id) references drinks(id),
    foreign key(location_id) references locations(id)
);

CREATE TRIGGER set_timestamp_user
BEFORE UPDATE ON users
FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_review
BEFORE UPDATE ON reviews
FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_drinks
BEFORE UPDATE ON drinks
FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp_locations
BEFORE UPDATE ON locations
FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

