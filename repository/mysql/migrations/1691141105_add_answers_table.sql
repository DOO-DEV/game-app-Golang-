-- +migrate Up
CREATE TABLE answers (
                         `id` INT PRIMARY KEY AUTO_INCREMENT,
                         `text` VARCHAR(255) NOT NULL ,
                         `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE answers;