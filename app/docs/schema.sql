CREATE DATABASE IF NOT EXISTS salesdb;
USE salesdb;
CREATE TABLE customers (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL,
    name VARCHAR(255),
    email VARCHAR(255),
    address VARCHAR(255),
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (id),
    UNIQUE KEY `uuid` (`uuid`)
);

CREATE TABLE products (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL,
    name VARCHAR(255),
    category VARCHAR(255),
    description VARCHAR(255),
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (id),
    UNIQUE KEY `uuid` (`uuid`)
);

CREATE TABLE orders (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL,
    customer_uuid VARCHAR(36),
    date_of_sale DATE,
    payment_method VARCHAR(255),
    shipping_cost DOUBLE,
    discount DOUBLE,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (id),
    UNIQUE KEY `uuid` (`uuid`),
    FOREIGN KEY (customer_uuid) REFERENCES customers(uuid) ON DELETE SET NULL
);

CREATE TABLE order_items (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    uuid VARCHAR(36) NOT NULL,
    order_uuid VARCHAR(36),
    product_uuid VARCHAR(36),
    quantity INT,
    unit_price DOUBLE,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (id),
    UNIQUE KEY `uuid` (`uuid`),
    FOREIGN KEY (order_uuid) REFERENCES orders(uuid) ON DELETE CASCADE,
    FOREIGN KEY (product_uuid) REFERENCES products(uuid) ON DELETE SET NULL
);


CREATE TABLE regions (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,s
    order_uuid VARCHAR(36),
    region_name VARCHAR(255),
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (id),
    UNIQUE KEY `uuid` (`uuid`),
    FOREIGN KEY (order_uuid) REFERENCES orders(uuid) ON DELETE CASCADE
);