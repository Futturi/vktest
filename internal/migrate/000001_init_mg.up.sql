CREATE TABLE actor
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    genre text,
    data int
);

CREATE TABLE cinema
(
    id BIGSERIAL PRIMARY KEY,
    name text,
    description text,
    data int,
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
    username text unique not null,
    password text not null
);

CREATE TABLE admins
(
    id BIGSERIAL PRIMARY KEY,
    username text unique not null,
    password text not null
)

