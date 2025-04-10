CREATE SCHEMA IF NOT EXISTS inventory;
CREATE SCHEMA IF NOT EXISTS orders;

CREATE TABLE inventory.categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE inventory.products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(12, 2) NOT NULL CHECK (price >= 0),
    category_id INT NOT NULL REFERENCES inventory.categories(id) ON DELETE RESTRICT,
    stock INTEGER NOT NULL DEFAULT 0 CHECK (stock >= 0),
    reserved INTEGER NOT NULL DEFAULT 0 CHECK (reserved >= 0 AND reserved <= stock),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE VIEW inventory.available_products AS
SELECT id, name, (stock - reserved) as available 
FROM inventory.products;

CREATE OR REPLACE FUNCTION inventory.update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_categories_updated_at
BEFORE UPDATE ON inventory.categories
FOR EACH ROW EXECUTE FUNCTION inventory.update_updated_at_column();

CREATE TRIGGER update_products_updated_at
BEFORE UPDATE ON inventory.products
FOR EACH ROW EXECUTE FUNCTION inventory.update_updated_at_column();

CREATE TYPE orders.order_status AS ENUM (
    'pending',
    'processing',
    'completed',
    'cancelled',
    'refunded'
);

CREATE TYPE orders.payment_status AS ENUM (
    'pending',
    'paid',
    'failed',
    'refunded'
);

CREATE TABLE orders.orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status orders.order_status NOT NULL DEFAULT 'pending',
    payment_status orders.payment_status NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(12, 2) NOT NULL CHECK (total_amount >= 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders.order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders.orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_at_purchase DECIMAL(12, 2) NOT NULL CHECK (price_at_purchase >= 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION orders.update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_orders_updated_at
BEFORE UPDATE ON orders.orders
FOR EACH ROW EXECUTE FUNCTION orders.update_updated_at_column();

CREATE INDEX idx_orders_user_id ON orders.orders(user_id);
CREATE INDEX idx_orders_status ON orders.orders(status);
CREATE INDEX idx_order_items_order_id ON orders.order_items(order_id);
CREATE INDEX idx_order_items_product_id ON orders.order_items(product_id);

CREATE OR REPLACE FUNCTION inventory.reserve_product(
    product_id INT,
    quantity INT
) RETURNS BOOLEAN AS $$
DECLARE
    success BOOLEAN;
BEGIN
    UPDATE inventory.products
    SET reserved = reserved + quantity
    WHERE id = product_id AND (stock - reserved) >= quantity
    RETURNING TRUE INTO success;
    
    RETURN COALESCE(success, FALSE);
END;
$$ LANGUAGE plpgsql;

-- Insert mock data into inventory.categories
INSERT INTO inventory.categories (name, description, created_at, updated_at)
VALUES
    ('Fruits', 'Fresh and organic fruits', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Vegetables', 'Fresh vegetables', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Dairy', 'Milk, cheese, and other dairy products', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Bakery', 'Bread, cakes, and pastries', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert mock data into inventory.products
INSERT INTO inventory.products (name, description, price, category_id, stock, reserved, created_at, updated_at)
VALUES
    ('Apple', 'Fresh red apples', 1.50, 1, 100, 10, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Banana', 'Organic bananas', 0.80, 1, 200, 20, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Carrot', 'Crunchy carrots', 0.60, 2, 150, 5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Milk', '1L whole milk', 1.20, 3, 50, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('Bread', 'Whole grain bread', 2.50, 4, 30, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert mock data into orders.orders
INSERT INTO orders.orders (user_id, status, payment_status, total_amount, created_at, updated_at)
VALUES
    (1, 'pending', 'pending', 15.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (2, 'processing', 'paid', 25.50, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (3, 'completed', 'paid', 10.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (4, 'cancelled', 'failed', 0.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert mock data into orders.order_items
INSERT INTO orders.order_items (order_id, product_id, quantity, price_at_purchase, created_at)
VALUES
    (1, 1, 5, 1.50, CURRENT_TIMESTAMP),
    (1, 2, 10, 0.80, CURRENT_TIMESTAMP),
    (2, 3, 15, 0.60, CURRENT_TIMESTAMP),
    (2, 4, 5, 1.20, CURRENT_TIMESTAMP),
    (3, 5, 4, 2.50, CURRENT_TIMESTAMP);