CREATE TABLE IF NOT EXISTS persons (
	id INT PRIMARY KEY,
	name VARCHAR(255),
	phone VARCHAR(20),
	age INT
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL primary key,
    age int,
    first_name text,
    last_name text,
    email text unique not null
);

INSERT INTO users (age, first_name, last_name, email)
    VALUES (99, 'Michael', 'Scott', 'mscott@dunder.com');
