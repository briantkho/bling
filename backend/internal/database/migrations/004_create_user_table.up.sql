CREATE TABLE IF NOT EXISTS user (
    user_id CHAR(36) PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash CHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    current_level_points INT DEFAULT 0,
    total_lifetime_points INT DEFAULT 0,
    level_id INT NOT NULL,
    FOREIGN KEY (level_id) REFERENCES user_level(level_id)
);