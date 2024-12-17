CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    active BOOLEAN NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users (name, age, active) VALUES
('Alice', 25, true),
('Bob', 30, true),
('Charlie', 22, false),
('David', 28, true),
('Eve', 35, false),
('Frank', 26, true),
('Grace', 24, false),
('Hannah', 29, true),
('Ivy', 31, false),
('Jack', 27, true);
