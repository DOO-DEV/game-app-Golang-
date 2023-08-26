-- +migrate Up
CREATE TABLE `games` (
                             `id` INT PRIMARY KEY AUTO_INCREMENT,
                             `start_time` TIMESTAMP,
                             `category_id` INT NOT NULL,
                             `question_id` INT NOT NULL,
                             `user_id` INT NOT NULL,
                             FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`),
                             FOREIGN KEY (`question_id`) REFERENCES `questions`(`id`),
                             FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- +migrate Down
DROP TABLE `games`;