DROP DATABASE IF exists products;
CREATE DATABASE IF NOT EXISTS products;

drop table if exists products CASCADE;
create table products(
	id  VARCHAR(100) PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
	image_url VARCHAR(500) NOT NULL,
	price FLOAT NOT NULL,
	rating FLOAT NOT NULL,
	offer BOOL NOT NULL DEFAULT FALSE,
	available BOOL NOT NULL,
    sales_days INT NOT NULL,
    sales_amount FLOAT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users(
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    permissions ENUM("1", "2", "4", "8", "16", "32", "64"),
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
