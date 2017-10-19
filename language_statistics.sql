-- Dumping structure for table gitea.language_statistics
CREATE TABLE IF NOT EXISTS `language_statistics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `repository_id` int(10) unsigned NOT NULL,
  `repository_branch` varchar(255) NOT NULL,
  `language` varchar(100) NOT NULL,
  `percentage` float unsigned NOT NULL,
  `color` char(7) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
