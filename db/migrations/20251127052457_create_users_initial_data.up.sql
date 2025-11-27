CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Insert Admin
INSERT INTO users (role, first_name, last_name, email, password)
VALUES ('Admin', 'Hafiz', 'Arrahman', 'fizrahman47@gmail.com', crypt('password123', gen_salt('bf')));

INSERT INTO userverified (user_id, phonenumber, verified)
VALUES (
           (SELECT user_id FROM users WHERE email = 'fizrahman47@gmail.com'),
           '6287801756611',
           true
       );

-- Insert Member
INSERT INTO users (role, first_name, last_name, email, password)
VALUES ('Member', 'Hafiz', 'Arrahman', 'hfizrrhman@gmail.com', crypt('password123', gen_salt('bf')));

INSERT INTO userverified (user_id, phonenumber, verified)
VALUES (
           (SELECT user_id FROM users WHERE email = 'hfizrrhman@gmail.com'),
           '6287801756622',
           true
       );