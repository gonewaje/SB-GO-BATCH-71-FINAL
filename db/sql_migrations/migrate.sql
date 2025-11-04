-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'customer', -- 'admin' | 'customer'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS restaurants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    phone VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS menu_items (
    id SERIAL PRIMARY KEY,
    restaurant_id INTEGER REFERENCES restaurants(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    available BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    restaurant_id INTEGER REFERENCES restaurants(id) ON DELETE SET NULL,
    total_price INTEGER NOT NULL DEFAULT 0,
    status VARCHAR(50) NOT NULL DEFAULT 'pending', -- 'pending' | 'paid' | 'cancelled'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id) ON DELETE CASCADE,
    menu_item_id INTEGER REFERENCES menu_items(id) ON DELETE SET NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    price_each INTEGER NOT NULL DEFAULT 0
);

-- Seed admin if not exists
INSERT INTO users (name, email, password, role)
SELECT 'Admin', 'admin@food.local', '$2a$10$HAUpF/iQ9ftBC1PCoaHnOOVCEJaLiAtGpxdPbm1o6dWTndHtfbUC2', 'admin'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email='admin@food.local');
-- Password (bcrypt): admin123

-- Seed a restaurant and a couple of menu items (if not exists)
INSERT INTO restaurants (name, address, phone)
SELECT 'Nasi Goreng Enak', 'Jalan Kenangan No. 1', '021-555-123'
WHERE NOT EXISTS (SELECT 1 FROM restaurants WHERE name='Nasi Goreng Enak');

INSERT INTO menu_items (restaurant_id, name, price, available)
SELECT r.id, 'Nasi Goreng Spesial', 30000, TRUE
FROM restaurants r WHERE r.name='Nasi Goreng Enak'
AND NOT EXISTS (SELECT 1 FROM menu_items m JOIN restaurants rr ON rr.id=m.restaurant_id WHERE m.name='Nasi Goreng Spesial' AND rr.name='Nasi Goreng Enak');

INSERT INTO menu_items (restaurant_id, name, price, available)
SELECT r.id, 'Mie Goreng', 25000, TRUE
FROM restaurants r WHERE r.name='Nasi Goreng Enak'
AND NOT EXISTS (SELECT 1 FROM menu_items m JOIN restaurants rr ON rr.id=m.restaurant_id WHERE m.name='Mie Goreng' AND rr.name='Nasi Goreng Enak');

-- +migrate Down
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS restaurants;
DROP TABLE IF EXISTS users;
