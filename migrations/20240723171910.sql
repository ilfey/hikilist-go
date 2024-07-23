-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "animes_lists" table
DROP TABLE `animes_lists`;
-- Create "new_collections" table
CREATE TABLE `new_collections` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `user_id` integer NULL,
  `name` text NOT NULL,
  `description` text NULL,
  `is_public` numeric NOT NULL,
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
-- Create "animes_collections" table
CREATE TABLE `animes_collections` (
  `collection_id` integer NULL,
  `anime_id` integer NULL,
  PRIMARY KEY (`collection_id`, `anime_id`),
  CONSTRAINT `fk_animes_collections_anime` FOREIGN KEY (`anime_id`) REFERENCES `animes` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_animes_collections_collection` FOREIGN KEY (`collection_id`) REFERENCES `collections` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
