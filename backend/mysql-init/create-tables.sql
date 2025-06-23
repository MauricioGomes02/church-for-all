CREATE DATABASE IF NOT EXISTS church_for_all;

USE church_for_all;

CREATE TABLE IF NOT EXISTS church (
    id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phones VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);
