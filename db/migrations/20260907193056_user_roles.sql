-- +goose Up

CREATE TABLE IF NOT EXISTS user_roles(
    id INT AUTO_INCREMENT PRIMARY KEY,

    user_id INT NOT NULL,

    role_id INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

-- No seeders data as we don't know what all users are there

-- +goose Down

DROP TABLE IF EXISTS user_roles;