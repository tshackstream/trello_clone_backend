
-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id` VARCHAR(32) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `name` VARCHAR(100) DEFAULT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME DEFAULT NULL,
    UNIQUE KEY uq_user_id_email (`user_id`, `email`)
);

-- +migrate Down
DROP TABLE IF EXISTS `users`;