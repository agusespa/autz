CREATE DATABASE auth;

USE auth;

CREATE TABLE user_auth (
    user_id INT AUTO_INCREMENT PRIMARY KEY, 
    user_uuid CHAR(36) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    refresh_token VARCHAR(512),
);

CREATE TABLE revoked_tokens (
    token_id INT AUTO_INCREMENT PRIMARY KEY,
    token_value VARCHAR(255) NOT NULL UNIQUE,
    revoked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_auth(user_id) ON DELETE CASCADE
);
