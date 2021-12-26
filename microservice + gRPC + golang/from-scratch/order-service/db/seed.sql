use order;

CREATE TABLE orders (
    id INT NOT NULL AUTO_INCREMENT,
    token VARCHAR(50) NOT NULL,
    created_at TIMESTAMP ,
    PRIMARY KEY (id)
);

CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    order_id INT NOT NULL, 
    name text NOT NULL,
    cost int ,
    weight float ,
    quantity int NOT NULL,
    data json,
    PRIMARY KEY (id),
    FOREIGN KEY (order_id) REFERENCES orders(ID) ON DELETE CASCADE ON UPDATE CASCADE
);

