-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_animes" table
CREATE TABLE `new_animes` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `title` text NOT NULL,
  `description` text NULL,
  `poster` text NULL,
  `episodes` integer NULL,
  `episodes_released` integer NOT NULL,
  `mal_id` integer NULL,
  `shiki_id` integer NULL
);
-- Copy rows from old table "animes" to new temporary table "new_animes"
INSERT INTO `new_animes` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `description`, `poster`, `episodes`, `episodes_released`, `mal_id`, `shiki_id`) SELECT `id`, `created_at`, `updated_at`, `deleted_at`, `title`, `description`, `poster`, `episodes`, `episodes_released`, `mal_id`, `shiki_id` FROM `animes`;
-- Drop "animes" table after copying rows
DROP TABLE `animes`;
-- Rename temporary table "new_animes" to "animes"
ALTER TABLE `new_animes` RENAME TO `animes`;
-- Create index "animes_mal_id" to table: "animes"
CREATE UNIQUE INDEX `animes_mal_id` ON `animes` (`mal_id`);
-- Create index "animes_shiki_id" to table: "animes"
CREATE UNIQUE INDEX `animes_shiki_id` ON `animes` (`shiki_id`);
-- Create index "idx_animes_deleted_at" to table: "animes"
CREATE INDEX `idx_animes_deleted_at` ON `animes` (`deleted_at`);
-- Create "new_collections" table
CREATE TABLE `new_collections` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `user_id` integer NULL,
  `name` text NOT NULL,
  `description` text NULL,
  `is_public` numeric NOT NULL DEFAULT true,
  CONSTRAINT `fk_users_collections` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Copy rows from old table "collections" to new temporary table "new_collections"
INSERT INTO `new_collections` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `name`, `description`, `is_public`) SELECT `id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `name`, `description`, `is_public` FROM `collections`;
-- Drop "collections" table after copying rows
DROP TABLE `collections`;
-- Rename temporary table "new_collections" to "collections"
ALTER TABLE `new_collections` RENAME TO `collections`;
-- Create index "idx_collections_deleted_at" to table: "collections"
CREATE INDEX `idx_collections_deleted_at` ON `collections` (`deleted_at`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
