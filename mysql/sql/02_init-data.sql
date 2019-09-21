INSERT INTO `sample_db`.`article_statuses`
(`status`)
VALUES
("公開");

INSERT INTO `sample_db`.`article_statuses`
(`status`)
VALUES
('削除');


INSERT INTO `sample_db`.`user_statuses`
(`status`)
VALUES
('有効');

INSERT INTO `sample_db`.`user_statuses`
(`status`)
VALUES
('無効');

INSERT INTO `sample_db`.`users`
(`username`, `password`, `status_id`)
VALUES
('sample', '$2a$10$xR4efFuokmGrHXeffMCNou4nBM2QoZKiu3OWo1YCWSktVXuMnIe8u', 1);
