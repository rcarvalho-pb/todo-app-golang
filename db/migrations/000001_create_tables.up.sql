CREATE TABLE IF NOT EXISTS todos (
    id INTEGER,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    status INTEGER DEFAULT 0,
    created_at DATE DEFAULT CURRENT_TIMESTAMP,
    last_modified_date DATE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_todos PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    create_at DATE DEFAULT CURRENT_TIMESTAMP,
    last_modified_date DATE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT pk_users PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS todos_users (
    todo_id INTEGER,
    user_id INTEGER,
    CONSTRAINT fk_todo FOREIGN KEY (todo_id) REFERENCES todos (id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);
