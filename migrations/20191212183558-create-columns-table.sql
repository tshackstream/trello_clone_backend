
-- +migrate Up
CREATE TABLE IF NOT EXISTS `columns` (
    `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `board_id` INT(11) NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `disp_order` INT(11) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS `columns`;
