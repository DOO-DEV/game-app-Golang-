-- +migrate Up
CREATE TABLE questions (
                         `id` INT PRIMARY KEY AUTO_INCREMENT,
                         `question` VARCHAR(255) NOT NULL ,
                         `correct_answer` INT NULL,
                         `difficulty` ENUM('1', '2', '3') DEFAULT '1',
                         `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         `answer_id` INT NOT NULL,
                         FOREIGN KEY (answer_id) REFERENCES answers(id)
);

-- +migrate Down
DROP TABLE questions;