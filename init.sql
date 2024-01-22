CREATE DATABASE test;

\connect test;

CREATE TABLE users(
                      email VARCHAR(255) NOT NULL PRIMARY KEY,
                      first VARCHAR(255),
                      last VARCHAR(255),
                      country VARCHAR(255),
                      city VARCHAR(255),
                      age int
);