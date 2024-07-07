-- Create "outstanding_tokens" table
CREATE TABLE `outstanding_tokens` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `token` text NOT NULL,
  `user_id` integer NULL,
  CONSTRAINT `fk_users_tokens` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_outstanding_tokens_deleted_at" to table: "outstanding_tokens"
CREATE INDEX `idx_outstanding_tokens_deleted_at` ON `outstanding_tokens` (`deleted_at`);
