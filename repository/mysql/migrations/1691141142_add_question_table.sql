-- +migrate Up
CREATE TABLE questions (
                         `id` INT PRIMARY KEY AUTO_INCREMENT,
                         `question` VARCHAR(255) NOT NULL ,
                         `difficulty` ENUM('1', '2', '3') DEFAULT '1',
                         `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         `answer_id` INT NOT NULL,
                         `category_id` INT NOT NULL,
                         FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- +migrate Down
DROP TABLE questions;