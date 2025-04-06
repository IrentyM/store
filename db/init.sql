-- Создание схемы (если нужно)
CREATE SCHEMA IF NOT EXISTS inventory;

-- Таблица категорий
CREATE TABLE inventory.categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

-- Индекс для быстрого поиска по имени категории
CREATE INDEX idx_categories_name ON inventory.categories (name);

-- Таблица продуктов
CREATE TABLE inventory.products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(12, 2) NOT NULL CHECK (price >= 0),
    category_id INT NOT NULL REFERENCES inventory.categories(id) ON DELETE RESTRICT,
    stock INTEGER NOT NULL DEFAULT 0 CHECK (stock >= 0),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для продуктов
CREATE INDEX idx_products_name ON inventory.products (name);
CREATE INDEX idx_products_category ON inventory.products (category_id);
CREATE INDEX idx_products_price ON inventory.products (price);
CREATE INDEX idx_products_stock ON inventory.products (stock);

-- Функция для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггеры для автоматического обновления updated_at
CREATE TRIGGER update_categories_updated_at
BEFORE UPDATE ON inventory.categories
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_products_updated_at
BEFORE UPDATE ON inventory.products
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();