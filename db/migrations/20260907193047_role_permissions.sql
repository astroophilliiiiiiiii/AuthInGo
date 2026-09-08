-- +goose Up

CREATE TABLE IF NOT EXISTS role_permissions(
    id INT AUTO_INCREMENT PRIMARY KEY,

    role_id INT NOT NULL,

    permission_id INT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,

    FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

-- Admin (role_ID = 1) has all the permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT 1, id
FROM permissions;

-- User (role_ID = 2) has only user:read permission
INSERT INTO role_permissions (role_id, permission_id)
SELECT 2, id
FROM permissions
WHERE name IN ('user:read');

-- +goose Down

DROP TABLE IF EXISTS role_permissions;