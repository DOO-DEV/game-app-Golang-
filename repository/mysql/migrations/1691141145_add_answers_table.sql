-- +migrate Up
CREATE TABLE answers (
                         `id` INT PRIMARY KEY AUTO_INCREMENT,
                         `text` VARCHAR(255) NOT NULL ,
                         `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         `question_id` INT NOT NULL,
                         FOREIGN KEY (question_id) REFERENCES questions(id)
);

-- +migrate Down
DROP TABLE answers;