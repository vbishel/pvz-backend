CREATE TABLE acceptance_products (
    acceptance_id INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY(acceptance_id, product_id),
    FOREIGN KEY (acceptance_id) REFERENCES acceptances(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);

CREATE INDEX idx_acceptance_products_acceptance_id ON acceptance_products(acceptance_id);
