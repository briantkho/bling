CREATE TABLE IF NOT EXISTS account (
    account_id INT AUTO_INCREMENT PRIMARY KEY,
    account_name VARCHAR(50) NOT NULL,
    account_type VARCHAR(50) NOT NULL,
    current_balance DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    user_id CHAR(36),
    FOREIGN KEY (user_id) references user(user_id)
);
