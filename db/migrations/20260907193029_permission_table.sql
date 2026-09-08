-- +goose Up

CREATE TABLE IF NOT EXISTS permissions(
    id INT AUTO_INCREMENT PRIMARY KEY,

    name VARCHAR(100) NOT NULL UNIQUE,

    description TEXT,

    resource VARCHAR(100) NOT NULL,

    action VARCHAR(50) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO permissions (name, description, resource, action) VALUES

('user:read', 'Permission to read user data', 'user', 'read'),

('user:create', 'Permission to create a new user', 'user', 'create'),

('user:update', 'Permission to update user data', 'user', 'update'),

('user:delete', 'Permission to delete a user', 'user', 'delete'),

('role:read', 'Permission to read role data', 'role', 'read'),

('role:create', 'Permission to create a new role', 'role', 'create'),

('role:update', 'Permission to update role data', 'role', 'update'),

('permission:read', 'Permission to read permission data', 'permission', 'read'),

('permission:create', 'Permission to create a new permission', 'permission', 'create'),

('permission:delete', 'Permission to delete a permission', 'permission', 'delete');

-- +goose Down

DROP TABLE IF EXISTS permissions;