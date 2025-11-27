CREATE EXTENSION IF NOT EXISTS pgcrypto;

---- Create User
INSERT INTO users (role, first_name, last_name, email, password)
VALUES
('Admin', 'Hafiz', 'Arrahman', 'fizrahman47@gmail.com',
        crypt('password123', gen_salt('bf'))),
('Member', 'Hafiz', 'Arrahman', 'hfizrrhman@gmail.com',
        crypt('password123', gen_salt('bf')));