-- Create "animes" table
CREATE TABLE `animes` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `title` text NULL
);
-- Create index "idx_animes_deleted_at" to table: "animes"
CREATE INDEX `idx_animes_deleted_at` ON `animes` (`deleted_at`);
-- Create "users" table
CREATE TABLE `users` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `username` text NULL,
  `password` text NULL
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX `idx_users_deleted_at` ON `users` (`deleted_at`);
