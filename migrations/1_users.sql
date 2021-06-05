CREATE TABLE IF NOT EXISTS users (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    city VARCHAR(255) NOT NULL,
    dorm VARCHAR(255) NOT NULL,
    room_number INT,
    university VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);