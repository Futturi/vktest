CREATE TABLE actor
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    genre text,
    data date
);

CREATE TABLE cinema
(
    id BIGSERIAL PRIMARY KEY,
    name text,
    description text,
    data date,
    rating int
);


CREATE TABLE author_cinema
(
    actor_id int references actor(id) on delete cascade,
    cinema_id int references cinema(id) on delete cascade
);

CREATE TABLE users
(
    id BIGSERIAL PRIMARY KEY,
    username text unique,
    password text
);

CREATE TABLE admins
(
    id BIGSERIAL PRIMARY KEY,
    username text unique,
    password text
)

