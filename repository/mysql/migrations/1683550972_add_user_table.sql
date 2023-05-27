-- +migrate Up
CREATE TABLE `users` (
                       `id` int PRIMARY KEY AUTO_INCREMENT,
                       `name` VARCHAR(255) NOT NULL ,
                       `phone_number` VARCHAR(255) NOT NULL UNIQUE ,
                       `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE `users`;