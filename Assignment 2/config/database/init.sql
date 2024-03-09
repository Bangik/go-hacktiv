CREATE DATABASE orders_by;

CREATE TABLE orders (
  order_id SERIAL PRIMARY KEY,
  customer_name VARCHAR(255) NOT NULL,
  ordered_at TIMESTAMP NOT NULL
);

CREATE TABLE items (
  item_id SERIAL PRIMARY KEY,
  item_code VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  order_id INT NOT NULL,

  FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE
);