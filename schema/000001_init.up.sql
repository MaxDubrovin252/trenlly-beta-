CREATE TABLE users(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);
CREATE TABLE trening(
    id serial not null unique,
    time integer not null,
    main_group varchar(255) not null,
    weight integer not null,
    height integer not null
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    tren_id int references trening (id) on delete cascade not null
);