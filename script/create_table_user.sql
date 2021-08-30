CREATE TABLE `user` (
  `id` int NOT NULL,
  `user_name` varchar(40) DEFAULT NULL,
  `parent` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
