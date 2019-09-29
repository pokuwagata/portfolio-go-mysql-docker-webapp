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

INSERT INTO `sample_db`.`posts`
(`user_id`)
VALUES
(1),
(1),
(1),
(1),
(1),
(1);

INSERT INTO `sample_db`.`articles`
(`title`, `content`, `post_id`, `user_id`, `article_status_id`, `created_at`, `updated_at`)
VALUES
('タイトル', '内容', 1, 1, 1, '2019-09-20 14:23:51', '2019-09-20 14:23:51'),
('タイトル', '内容', 2, 1, 1, '2019-09-21 14:23:51', '2019-09-21 14:23:51'),
('タイトル', '内容', 3, 1, 1, '2019-09-22 14:23:51', '2019-09-22 14:23:51'),
('タイトル', '内容', 4, 1, 1, '2019-09-23 14:23:51', '2019-09-23 14:23:51'),
('タイトル', '内容', 5, 1, 1, '2019-09-24 14:23:51', '2019-09-24 14:23:51'),
('タイトル', '内容', 6, 1, 1, '2019-09-25 14:23:51', '2019-09-25 14:23:51');