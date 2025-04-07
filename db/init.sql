-- Создаем отдельные схемы для каждого микросервиса
CREATE SCHEMA IF NOT EXISTS inventory;
CREATE SCHEMA IF NOT EXISTS orders;

-- Создаем пользователей для каждого сервиса (опционально, но рекомендуется)
CREATE ROLE inventory_service WITH LOGIN PASSWORD 'inventory_password';
CREATE ROLE order_service WITH LOGIN PASSWORD 'order_password';

-- Назначаем права на схемы
GRANT USAGE ON SCHEMA inventory TO inventory_service;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA inventory TO inventory_service;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA inventory TO inventory_service;

GRANT USAGE ON SCHEMA orders TO order_service;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA orders TO order_service;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA orders TO order_service;

-- Таблица категорий (остается без изменений)
CREATE TABLE inventory.categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица продуктов (добавляем поле для блокировки при заказе)
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

-- Представление для доступного количества (stock - reserved)
CREATE VIEW inventory.available_products AS
SELECT id, name, (stock - reserved) as available 
FROM inventory.products;

-- Триггеры (остаются без изменений)
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

-- Типы статусов
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

-- Таблица заказов
CREATE TABLE orders.orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status orders.order_status NOT NULL DEFAULT 'pending',
    payment_status orders.payment_status NOT NULL DEFAULT 'pending',
    total_amount DECIMAL(12, 2) NOT NULL CHECK (total_amount >= 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица элементов заказа
CREATE TABLE orders.order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL REFERENCES orders.orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    price_at_purchase DECIMAL(12, 2) NOT NULL CHECK (price_at_purchase >= 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Функция для обновления updated_at
CREATE OR REPLACE FUNCTION orders.update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггер для заказов
CREATE TRIGGER update_orders_updated_at
BEFORE UPDATE ON orders.orders
FOR EACH ROW EXECUTE FUNCTION orders.update_updated_at_column();

-- Индексы для ускорения запросов
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

-- Даем доступ Order Service только к этой функции
GRANT EXECUTE ON FUNCTION inventory.reserve_product(INT, INT) TO order_service;