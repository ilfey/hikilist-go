-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_outstanding_tokens" table
CREATE TABLE `new_outstanding_tokens` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `token` text NOT NULL
);
-- Copy rows from old table "outstanding_tokens" to new temporary table "new_outstanding_tokens"
INSERT INTO `new_outstanding_tokens` (`id`, `created_at`, `updated_at`, `deleted_at`, `token`) SELECT `id`, `created_at`, `updated_at`, `deleted_at`, `token` FROM `outstanding_tokens`;
-- Drop "outstanding_tokens" table after copying rows
DROP TABLE `outstanding_tokens`;
-- Rename temporary table "new_outstanding_tokens" to "outstanding_tokens"
ALTER TABLE `new_outstanding_tokens` RENAME TO `outstanding_tokens`;
-- Create index "idx_outstanding_tokens_deleted_at" to table: "outstanding_tokens"
CREATE INDEX `idx_outstanding_tokens_deleted_at` ON `outstanding_tokens` (`deleted_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
