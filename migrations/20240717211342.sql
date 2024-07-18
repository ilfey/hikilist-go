-- Create "user_actions" table
CREATE TABLE `user_actions` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `user_id` integer NULL,
  `title` text NOT NULL,
  `description` text NOT NULL,
  CONSTRAINT `fk_user_actions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_user_actions_deleted_at" to table: "user_actions"
CREATE INDEX `idx_user_actions_deleted_at` ON `user_actions` (`deleted_at`);
