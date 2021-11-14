-- +goose Up
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password_digest` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_users_on_email` (`email`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb4;
CREATE TABLE `agents` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `working_time` varchar(255) DEFAULT NULL,
  `telephone_number` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `photo_url` varchar(255) DEFAULT NULL,
  `slogan` varchar(255) DEFAULT NULL,
  `access` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_agents_on_name` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 150 DEFAULT CHARSET = utf8mb4;
CREATE TABLE `buildings` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `access` varchar(255) DEFAULT NULL,
  `year_built` date DEFAULT NULL,
  `building_type` varchar(255) DEFAULT NULL,
  `structure` varchar(255) DEFAULT NULL,
  `storeys` int(11) DEFAULT NULL,
  `underground_storeys` int(11) DEFAULT NULL,
  `photo_url` varchar(255) DEFAULT NULL,
  `longitude` decimal(20, 13) DEFAULT NULL,
  `latitude` decimal(20, 13) DEFAULT NULL,
  `average_size` float DEFAULT NULL,
  `average_fee` float DEFAULT NULL,
  `distance` decimal(20, 13) DEFAULT NULL,
  `condition_type` int(11) DEFAULT '0',
  `office_id` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_buildings_on_name` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 93 DEFAULT CHARSET = utf8mb4;
CREATE TABLE `rooms` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `suumo_id` varchar(255) DEFAULT NULL,
  `building_id` bigint(20) DEFAULT NULL,
  `agent_id` bigint(20) DEFAULT NULL,
  `rent_fee` decimal(10, 0) DEFAULT NULL,
  `reikin` decimal(10, 0) DEFAULT NULL,
  `shikikin` decimal(10, 0) DEFAULT NULL,
  `management_cost` decimal(10, 0) DEFAULT NULL,
  `caution_fee` decimal(10, 0) DEFAULT NULL,
  `layout_image_url` varchar(255) DEFAULT NULL,
  `size` int(11) DEFAULT NULL,
  `direction` varchar(255) DEFAULT NULL,
  `facilities` text,
  `floor` int(11) DEFAULT NULL,
  `car_park` varchar(255) DEFAULT NULL,
  `condition` varchar(255) DEFAULT NULL,
  `note` text,
  `layout` varchar(255) DEFAULT NULL,
  `layout_detail` varchar(255) DEFAULT NULL,
  `deal_type` varchar(255) DEFAULT NULL,
  `move_in_time` date DEFAULT NULL,
  `move_in` varchar(255) DEFAULT NULL,
  `damage_insurance` varchar(255) DEFAULT NULL,
  `guarantor` varchar(255) DEFAULT NULL,
  `other_fees` varchar(255) DEFAULT NULL,
  `other_initial_fees` varchar(255) DEFAULT NULL,
  `last_update` date DEFAULT NULL,
  `suumo_link` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`building_id`) REFERENCES `buildings`(`id`),
  FOREIGN KEY (`agent_id`) REFERENCES `agents`(`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 661 DEFAULT CHARSET = utf8mb4;
CREATE TABLE `images` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `url` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `room_id` bigint(20) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 11483 DEFAULT CHARSET = utf8mb4;
CREATE TABLE `user_rooms` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `room_id` bigint(20) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- SQL in this section is executed when the migration is applied.
-- +goose Down
DROP TABLE IF EXISTS `images`;
DROP TABLE IF EXISTS `user_rooms`;
DROP TABLE IF EXISTS `rooms`;
DROP TABLE IF EXISTS `buildings`;
DROP TABLE IF EXISTS `agents`;
DROP TABLE IF EXISTS `users`;
-- SQL in this section is executed when the migration is rolled back.
