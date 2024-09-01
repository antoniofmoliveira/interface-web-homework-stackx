CREATE DATABASE stackx;
CREATE USER * * * * * * *;
USE stackx;
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name STRING NOT NULL,
    email STRING NOT NULL,
    dob STRING NOT NULL,
    age INT NOT NULL
)
GRANT ALL ON stackx.* TO * * * * * * *;