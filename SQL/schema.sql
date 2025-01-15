CREATE TABLE users (
    user_id   SERIAL PRIMARY KEY,
    name text      NOT NULL,
    password text NOT NULL
);

CREATE TABLE groups (
    group_id SERIAL PRIMARY KEY,
    groupName text NOT NULL
);

CREATE  TABLE roles (
    role_id SERIAL PRIMARY KEY,
    roleName text NOT NULL
);

CREATE TABLE user_groups (
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    group_id INT NOT NULL REFERENCES groups(group_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, group_id)
);

CREATE TABLE user_roles (
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    role_id INT NOT NULL REFERENCES roles(role_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE group_roles (
    group_id INT NOT NULL REFERENCES groups(group_id) ON DELETE CASCADE,
    role_id INT NOT NULL REFERENCES roles(role_id) ON DELETE CASCADE,
    PRIMARY KEY (group_id, role_id)
);
