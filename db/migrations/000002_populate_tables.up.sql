INSERT INTO todos (name, description) VALUES
    ('Teste 01', 'Primeiro Teste'),
    ('Teste 02', 'Segundo Teste');

INSERT INTO users (first_name, last_name, email, password) VALUES
    ('Ramon', 'Carvalho', 'ramon@email.com', '123'),
    ('Emilly', 'Coeli', 'emilly@email.com', '123');

INSERT INTO todos_users VALUES 
    (1, 1),
    (1, 2),
    (2, 1),
    (2, 2);
