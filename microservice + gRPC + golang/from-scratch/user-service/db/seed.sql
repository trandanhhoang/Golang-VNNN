use user;

CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    token VARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (token) VALUES ('abc');
