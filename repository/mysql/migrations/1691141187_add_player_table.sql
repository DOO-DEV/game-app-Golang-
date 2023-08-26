-- +migrate Up
CREATE TABLE `players` (
                             `id` INT PRIMARY KEY AUTO_INCREMENT,
                             `score` INT DEFAULT 0,
                             `answer_id` INT NOT NULL,
                             `user_id` INT NOT NULL,
                             `game_id` INT NOT NULL,
                             FOREIGN KEY (`answer_id`) REFERENCES `answers`(`id`),
                             FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                             FOREIGN KEY (`game_id`) REFERENCES `games`(`id`)
);

-- +migrate Down
DROP TABLE `players`;