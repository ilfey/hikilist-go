-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_user_actions" table
CREATE TABLE `new_user_actions` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `user_id` integer NULL,
  `title` text NOT NULL,
  `description` text NOT NULL,
  CONSTRAINT `fk_user_actions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Copy rows from old table "user_actions" to new temporary table "new_user_actions"
INSERT INTO `new_user_actions` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `title`, `description`) SELECT `id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `title`, `description` FROM `user_actions`;
-- Drop "user_actions" table after copying rows
DROP TABLE `user_actions`;
-- Rename temporary table "new_user_actions" to "user_actions"
ALTER TABLE `new_user_actions` RENAME TO `user_actions`;
-- Create index "idx_user_actions_deleted_at" to table: "user_actions"
CREATE INDEX `idx_user_actions_deleted_at` ON `user_actions` (`deleted_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
