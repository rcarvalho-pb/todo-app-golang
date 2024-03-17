INSERT INTO todos (name, description) VALUES
    ('Teste 01', 'Primeiro Teste'),
    ('Teste 02', 'Segundo Teste'),
    ('Teste 03', 'Terceiro Teste');

INSERT INTO users (first_name, last_name, email, password, role) VALUES
    ('Ramon', 'Carvalho', 'ramon@email.com', '123', 2),
    ('Emilly', 'Coeli', 'emilly@email.com', '123', 1);

INSERT INTO todos_users VALUES 
    (1, 1),
    (2, 1),
    (2, 2),
    (3, 1);
