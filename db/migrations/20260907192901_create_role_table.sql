-- +goose Up
CREATE TABLE IF NOT EXISTS roles(
    id INT AUTO_INCREMENT PRIMARY KEY , 
    name VARCHAR(100) NOT NULL UNIQUE , 
    description TEXT , 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , 
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

INSERT INTO roles ( name , description ) VALUES 
( 'admin' , 'Administrator with full access') , 
('user' , 'Regular user with limited access') , 
('moderator' , 'Moderator with elevated priveledges');

-- +goose Down
DROP TABLE IF EXISTS roles ; 
