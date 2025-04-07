-- 1 Order Service
CREATE TABLE IF NOT EXISTS orders(
    id bigserial PRIMARY KEY,
    customername VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',  --ending, completed, cancelled
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS order_items (
    orderID INT,
    productID INT,
    quantity INT NOT NULL CHECK(quantity > 0),
    PRIMARY KEY (orderID, productID),
    FOREIGN KEY (orderID) REFERENCES orders(id) ON DELETE CASCADE
);

-- -- 2 Inventory service
-- CREATE TABLE products (
--     productID bigserial PRIMARY KEY,
--     name VARCHAR(50) NOT NULL,
--     description TEXT NOT NULL,
--     price NUMERIC(10, 2) NOT NULL CHECK(price > 0),
--     available INT DEFAULT 0 CHECK(available > 0)
-- );

