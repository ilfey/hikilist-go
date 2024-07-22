-- Create "collections" table
CREATE TABLE `collections` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `user_id` integer NULL,
  `name` text NULL,
  `description` text NULL,
  `is_public` numeric NOT NULL DEFAULT true,
  CONSTRAINT `fk_users_collections` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_collections_deleted_at" to table: "collections"
CREATE INDEX `idx_collections_deleted_at` ON `collections` (`deleted_at`);
-- Create "animes_lists" table
CREATE TABLE `animes_lists` (
  `collection_id` integer NULL,
  `anime_id` integer NULL,
  PRIMARY KEY (`collection_id`, `anime_id`),
  CONSTRAINT `fk_animes_lists_collection` FOREIGN KEY (`collection_id`) REFERENCES `collections` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_animes_lists_anime` FOREIGN KEY (`anime_id`) REFERENCES `animes` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
