CREATE TABLE `scene` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `video_project_id` int(11) NOT NULL,
                         `sequence` int(11) NOT NULL,
                         `title` text NOT NULL,
                         `narration` text NOT NULL,
                         `illustration` text NOT NULL,
                         `illustration_url` varchar(255) DEFAULT NULL,
                         `voice_url` varchar(255) DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `scene_video_project_id_fk` (`video_project_id`),
                         CONSTRAINT `scene_video_project_id_fk` FOREIGN KEY (`video_project_id`) REFERENCES `video_project` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci

CREATE TABLE `user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `email` varchar(255) NOT NULL,
                        `fullname` varchar(255) NOT NULL,
                        `passwd` varchar(255) DEFAULT NULL,
                        `passwdSalt` varchar(30) DEFAULT NULL,
                        `provider` varchar(30) DEFAULT NULL,
                        `provider_id` varchar(30) DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `user_pk2` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci

CREATE TABLE `video_project` (
                                 `id` int(11) NOT NULL AUTO_INCREMENT,
                                 `user_id` int(11) NOT NULL,
                                 `product_title` varchar(255) NOT NULL,
                                 `brand_name` varchar(255) NOT NULL,
                                 `product_type` varchar(255) NOT NULL,
                                 `market_target` varchar(255) NOT NULL,
                                 `superiority` text NOT NULL,
                                 `duration` int(11) NOT NULL,
                                 PRIMARY KEY (`id`),
                                 KEY `video_project_user_id_fk` (`user_id`),
                                 CONSTRAINT `video_project_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci

