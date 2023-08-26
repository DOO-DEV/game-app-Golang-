-- +migrate Up
CREATE TABLE `categories` (
                           `id` INT PRIMARY KEY AUTO_INCREMENT,
                           `name` VARCHAR(255) NOT NULL

);

-- +migrate Down
DROP TABLE `categories`;