-- Create index "animes_mal_id" to table: "animes"
CREATE UNIQUE INDEX `animes_mal_id` ON `animes` (`mal_id`);
-- Create index "animes_shiki_id" to table: "animes"
CREATE UNIQUE INDEX `animes_shiki_id` ON `animes` (`shiki_id`);
-- Add column "last_online" to table: "users"
ALTER TABLE `users` ADD COLUMN `last_online` datetime NULL;
