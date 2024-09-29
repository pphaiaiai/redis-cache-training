DROP TABLE IF EXISTS product CASCADE;

CREATE TABLE product (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    image_url VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

INSERT INTO product (name, description, price, image_url)
VALUES
('Product 1', 'Description for product 1', 100.00, 'https://placehold.it/200x200'),
('Product 2', 'Description for product 2', 150.00, 'https://placehold.it/200x200'),
('Product 3', 'Description for product 3', 200.00, 'https://placehold.it/200x200'),
('Product 4', 'Description for product 4', 250.00, 'https://placehold.it/200x200'),
('Product 5', 'Description for product 5', 300.00, 'https://placehold.it/200x200'),
('Product 6', 'Description for product 6', 350.00, 'https://placehold.it/200x200'),
('Product 7', 'Description for product 7', 400.00, 'https://placehold.it/200x200'),
('Product 8', 'Description for product 8', 450.00, 'https://placehold.it/200x200'),
('Product 9', 'Description for product 9', 500.00, 'https://placehold.it/200x200'),
('Product 10', 'Description for product 10', 550.00, 'https://placehold.it/200x200');
