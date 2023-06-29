-- +migrate Up
INSERT INTO permission (`id`, `title`) VALUES(1, 'user-list');
INSERT INTO permission (`id`, `title`) VALUES(2, 'user-delete');

INSERT INTO `access_control` (`actor_type`, `actor_id`, `permission_id`) VALUES ('role', 2, 1);
INSERT INTO `access_control` (`actor_type`, `actor_id`, `permission_id`) VALUES ('role', 2, 2);

-- +migrate Down
DELETE FROM `permission` WHERE id in(1,2);