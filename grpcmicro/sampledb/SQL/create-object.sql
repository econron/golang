CREATE TABLE sample_db.category (
    id INT NOT NULL AUTO_INCREMENT,
    obj_id VARCHAR(36) NOT NULL,
    name VARCHAR(20) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_obj_id (obj_id)
)

CREATE TABLE sample_db.product (
    id INT NOT NULL AUTO_INCREMENT,
    obj_id VARCHAR(36) NOT NULL,
    name VARCHAR(20) NOT NULL,
    category_id INT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_obj_id (obj_id),
    FOREIGN KEY (category_id) REFERENCES category(id)
)