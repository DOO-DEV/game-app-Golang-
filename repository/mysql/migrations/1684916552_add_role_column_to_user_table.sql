-- +migrate Up
-- in mysql v8 default is user
ALTER TABLE `users` add column `role` ENUM('user', 'admin') NOT NULL;

-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;