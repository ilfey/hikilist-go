-- Add column "description" to table: "animes"
ALTER TABLE `animes` ADD COLUMN `description` text NULL;
-- Add column "poster" to table: "animes"
ALTER TABLE `animes` ADD COLUMN `poster` text NULL;
-- Add column "episodes" to table: "animes"
ALTER TABLE `animes` ADD COLUMN `episodes` integer NULL;
-- Add column "episodes_released" to table: "animes"
ALTER TABLE `animes` ADD COLUMN `episodes_released` integer NULL;
-- Add column "mal_id" to table: "animes"
ALTER TABLE `animes` ADD COLUMN `mal_id` integer NULL;
-- Create "animes_related" table
CREATE TABLE `animes_related` (
  `anime_id` integer NULL,
  `related_id` integer NULL,
  PRIMARY KEY (`anime_id`, `related_id`),
  CONSTRAINT `fk_animes_related_anime` FOREIGN KEY (`anime_id`) REFERENCES `animes` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT `fk_animes_related_related` FOREIGN KEY (`related_id`) REFERENCES `animes` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);
