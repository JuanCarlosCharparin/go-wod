-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: wod-db-go
-- ------------------------------------------------------
-- Server version	9.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `calendar`
--

DROP TABLE IF EXISTS `calendar`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `calendar` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `status` enum('inscripto','presente','ausente','cancelado') DEFAULT 'inscripto',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `calendar_user_id_foreign` (`user_id`),
  KEY `calendar_class_id_foreign` (`class_id`),
  CONSTRAINT `calendar_class_id_foreign` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `calendar_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `calendar`
--

LOCK TABLES `calendar` WRITE;
/*!40000 ALTER TABLE `calendar` DISABLE KEYS */;
INSERT INTO `calendar` VALUES (1,1,1,'inscripto','2025-06-09 23:42:25','2025-06-09 23:42:25'),(2,3,1,'inscripto','2025-06-17 01:10:55','2025-06-17 01:10:55'),(3,3,5,'inscripto','2025-06-17 21:43:47','2025-06-17 21:43:47'),(4,1,5,'inscripto','2025-06-21 22:17:35','2025-06-21 22:17:35'),(5,2,5,'inscripto','2025-06-21 22:17:42','2025-06-21 22:17:42'),(7,4,5,'inscripto','2025-06-22 04:04:30','2025-06-22 04:04:30'),(8,4,6,'inscripto','2025-06-22 04:10:47','2025-06-22 04:10:47'),(9,4,7,'inscripto','2025-06-22 15:20:11','2025-06-22 15:20:11'),(10,1,7,'inscripto','2025-06-22 15:20:20','2025-06-22 15:20:20'),(11,2,7,'inscripto','2025-06-22 15:20:24','2025-06-22 15:20:24'),(12,7,7,'inscripto','2025-06-22 21:37:27','2025-06-22 21:37:27'),(13,3,7,'inscripto','2025-06-22 21:39:15','2025-06-22 21:39:15'),(14,3,8,'inscripto','2025-06-22 21:42:56','2025-06-22 21:42:56'),(15,3,9,'inscripto','2025-06-22 21:43:20','2025-06-22 21:43:20'),(16,1,9,'inscripto','2025-06-22 21:43:30','2025-06-22 21:43:30'),(17,1,10,'inscripto','2025-06-22 21:43:42','2025-06-22 21:43:42'),(18,3,45,'ausente','2025-06-23 22:24:21','2025-06-25 21:46:18'),(19,3,41,'ausente','2025-06-25 21:50:16','2025-06-25 21:50:28'),(20,3,42,'ausente','2025-06-25 21:51:07','2025-06-25 21:51:24'),(21,3,43,'cancelado','2025-06-25 21:52:07','2025-06-25 21:52:16'),(22,3,43,'inscripto','2025-06-25 22:53:01','2025-06-25 22:53:01'),(23,1,83,'inscripto','2025-07-01 22:33:23','2025-07-01 22:33:23'),(24,2,83,'inscripto','2025-07-01 22:33:32','2025-07-01 22:33:32'),(25,4,83,'inscripto','2025-07-01 22:33:56','2025-07-01 22:33:56'),(26,7,83,'inscripto','2025-07-01 22:34:29','2025-07-01 22:34:29'),(27,3,83,'inscripto','2025-07-02 00:34:52','2025-07-02 00:34:52'),(28,8,83,'inscripto','2025-07-02 00:39:09','2025-07-02 00:39:09'),(29,9,83,'inscripto','2025-07-02 00:39:16','2025-07-02 00:39:16'),(30,10,83,'inscripto','2025-07-02 00:39:19','2025-07-02 00:39:19'),(31,11,83,'inscripto','2025-07-02 00:39:24','2025-07-02 00:39:24'),(32,12,83,'inscripto','2025-07-02 00:42:20','2025-07-02 00:42:20'),(33,1,86,'inscripto','2025-07-06 21:41:03','2025-07-06 21:41:03'),(34,2,86,'inscripto','2025-07-06 21:43:56','2025-07-06 21:43:56'),(35,3,86,'inscripto','2025-07-06 21:44:12','2025-07-06 21:44:12'),(36,1,88,'inscripto','2025-07-06 21:49:17','2025-07-06 21:49:17'),(37,2,88,'inscripto','2025-07-06 21:49:28','2025-07-06 21:49:28'),(38,3,88,'inscripto','2025-07-06 21:49:34','2025-07-06 21:49:34'),(39,4,88,'inscripto','2025-07-06 21:49:42','2025-07-06 21:49:42'),(40,7,88,'inscripto','2025-07-06 21:49:54','2025-07-06 21:49:54'),(41,8,88,'inscripto','2025-07-06 21:49:58','2025-07-06 21:49:58'),(42,9,88,'inscripto','2025-07-06 21:50:25','2025-07-06 21:50:25'),(43,10,88,'inscripto','2025-07-06 21:50:29','2025-07-06 21:50:29'),(44,11,88,'inscripto','2025-07-06 21:50:34','2025-07-06 21:50:34'),(45,12,88,'inscripto','2025-07-06 21:50:50','2025-07-06 21:50:50'),(46,1,89,'inscripto','2025-07-06 23:38:41','2025-07-06 23:38:41'),(47,7,86,'inscripto','2025-07-07 00:28:11','2025-07-07 00:28:11'),(48,7,87,'inscripto','2025-07-07 00:28:26','2025-07-07 00:28:26'),(49,7,89,'inscripto','2025-07-07 00:28:37','2025-07-07 00:28:37'),(50,7,90,'inscripto','2025-07-07 00:28:53','2025-07-07 00:28:53'),(51,7,91,'inscripto','2025-07-07 00:29:00','2025-07-07 00:29:00'),(52,2,92,'inscripto','2025-07-07 00:45:07','2025-07-07 00:45:07'),(53,3,92,'cancelado','2025-07-07 22:01:40','2025-07-07 22:08:51'),(54,3,92,'cancelado','2025-07-07 22:08:18','2025-07-07 22:10:15'),(55,4,125,'inscripto','2025-07-20 23:37:52','2025-07-20 23:37:52'),(56,4,124,'inscripto','2025-07-20 23:38:59','2025-07-20 23:38:59'),(57,4,123,'inscripto','2025-07-20 23:39:04','2025-07-20 23:39:04'),(58,4,122,'inscripto','2025-07-20 23:39:08','2025-07-20 23:39:08'),(59,4,121,'inscripto','2025-07-20 23:39:12','2025-07-20 23:39:12'),(60,4,120,'inscripto','2025-07-20 23:39:24','2025-07-20 23:39:24'),(61,4,119,'inscripto','2025-07-20 23:39:38','2025-07-20 23:39:38'),(62,4,118,'inscripto','2025-07-20 23:39:47','2025-07-20 23:39:47'),(63,3,117,'inscripto','2025-07-21 21:59:52','2025-07-21 21:59:52'),(64,2,117,'inscripto','2025-07-21 22:01:23','2025-07-21 22:01:23');
/*!40000 ALTER TABLE `calendar` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `classes`
--

DROP TABLE IF EXISTS `classes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `classes` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `time` time NOT NULL,
  `capacity` smallint unsigned NOT NULL,
  `gym_id` bigint unsigned NOT NULL,
  `discipline_id` bigint unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_class` (`gym_id`,`date`,`time`),
  KEY `classes_discipline_id_foreign` (`discipline_id`),
  CONSTRAINT `classes_discipline_id_foreign` FOREIGN KEY (`discipline_id`) REFERENCES `disciplines` (`id`) ON DELETE CASCADE,
  CONSTRAINT `classes_gym_id_foreign` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classes`
--

LOCK TABLES `classes` WRITE;
/*!40000 ALTER TABLE `classes` DISABLE KEYS */;
INSERT INTO `classes` VALUES (1,'2025-06-09','20:00:00',25,1,1,'2025-06-08 22:39:19','2025-06-09 00:05:24',NULL),(2,'2025-06-09','21:00:00',20,1,1,'2025-06-09 00:06:06','2025-06-09 00:06:06',NULL),(3,'2025-06-17','20:00:00',25,1,1,'2025-06-15 17:22:38','2025-06-15 17:22:38',NULL),(4,'2025-06-18','20:00:00',25,1,1,'2025-06-15 21:43:51','2025-06-15 21:43:51',NULL),(5,'2025-06-19','10:00:00',15,1,1,'2025-06-15 22:43:25','2025-06-15 22:43:25',NULL),(6,'2025-06-23','08:00:00',20,1,1,'2025-06-21 15:56:33','2025-06-21 15:56:33',NULL),(7,'2025-06-23','09:00:00',15,1,1,'2025-06-21 15:56:33','2025-06-21 15:56:33',NULL),(8,'2025-06-23','10:00:00',10,1,1,'2025-06-21 15:56:33','2025-06-21 15:56:33',NULL),(9,'2025-06-23','15:00:00',15,1,1,'2025-06-21 15:56:34','2025-06-21 15:56:34',NULL),(10,'2025-06-23','18:00:00',10,1,1,'2025-06-21 15:56:34','2025-06-21 15:56:34',NULL),(16,'2025-06-24','08:00:00',16,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(17,'2025-06-24','09:00:00',16,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(18,'2025-06-24','14:00:00',16,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(19,'2025-06-24','15:00:00',16,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(20,'2025-06-24','18:00:00',16,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(21,'2025-06-24','19:00:00',18,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(22,'2025-06-24','20:00:00',18,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(23,'2025-06-24','21:00:00',16,1,1,'2025-06-21 16:07:11','2025-06-21 16:07:11',NULL),(37,'2025-06-25','08:00:00',16,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(38,'2025-06-25','09:00:00',16,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(39,'2025-06-25','14:00:00',16,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(40,'2025-06-25','15:00:00',16,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(41,'2025-06-25','18:00:00',16,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(42,'2025-06-25','19:00:00',18,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(43,'2025-06-25','20:00:00',18,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(44,'2025-06-25','21:00:00',16,1,1,'2025-06-21 22:04:02','2025-06-21 22:04:02',NULL),(45,'2025-06-23','20:00:00',16,1,1,'2025-06-23 22:19:16','2025-06-23 22:19:16',NULL),(46,'2025-06-30','08:00:00',20,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(47,'2025-06-30','09:00:00',15,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(48,'2025-06-30','10:00:00',10,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(49,'2025-06-30','15:00:00',15,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(50,'2025-06-30','18:00:00',10,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(51,'2025-07-01','08:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(52,'2025-07-01','09:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(53,'2025-07-01','14:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(54,'2025-07-01','15:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(55,'2025-07-01','18:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(56,'2025-07-01','19:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(57,'2025-07-01','20:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(58,'2025-07-01','21:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(59,'2025-07-02','08:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(60,'2025-07-02','09:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(61,'2025-07-02','14:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(62,'2025-07-02','15:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(63,'2025-07-02','18:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(64,'2025-07-02','19:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(65,'2025-07-02','20:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(66,'2025-07-02','21:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(67,'2025-07-03','08:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(68,'2025-07-03','09:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(69,'2025-07-03','14:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(70,'2025-07-03','15:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(71,'2025-07-03','18:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(72,'2025-07-03','19:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(73,'2025-07-03','20:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(74,'2025-07-03','21:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(75,'2025-07-04','08:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(76,'2025-07-04','09:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(77,'2025-07-04','14:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(78,'2025-07-04','15:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(79,'2025-07-04','18:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(80,'2025-07-04','19:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(81,'2025-07-04','20:00:00',18,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(82,'2025-07-04','21:00:00',16,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(83,'2025-07-05','08:00:00',10,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(84,'2025-07-05','09:00:00',10,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(85,'2025-07-05','14:00:00',10,1,1,'2025-06-30 22:50:40','2025-06-30 22:50:40',NULL),(86,'2025-07-07','08:00:00',20,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(87,'2025-07-07','09:00:00',15,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(88,'2025-07-07','10:00:00',10,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(89,'2025-07-07','15:00:00',15,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(90,'2025-07-07','18:00:00',10,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(91,'2025-07-08','08:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(92,'2025-07-08','09:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(93,'2025-07-08','14:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(94,'2025-07-08','15:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(95,'2025-07-08','18:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(96,'2025-07-08','19:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(97,'2025-07-08','20:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(98,'2025-07-08','21:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(99,'2025-07-09','08:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(100,'2025-07-09','09:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(101,'2025-07-09','14:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(102,'2025-07-09','15:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(103,'2025-07-09','18:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(104,'2025-07-09','19:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(105,'2025-07-09','20:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(106,'2025-07-09','21:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(107,'2025-07-10','08:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(108,'2025-07-10','09:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(109,'2025-07-10','14:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(110,'2025-07-10','15:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(111,'2025-07-10','18:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(112,'2025-07-10','19:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(113,'2025-07-10','20:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(114,'2025-07-10','21:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(115,'2025-07-11','08:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(116,'2025-07-11','09:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(117,'2025-07-11','14:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(118,'2025-07-11','15:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(119,'2025-07-11','18:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(120,'2025-07-11','19:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(121,'2025-07-11','20:00:00',18,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(122,'2025-07-11','21:00:00',16,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(123,'2025-07-12','08:00:00',10,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(124,'2025-07-12','09:00:00',10,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(125,'2025-07-12','14:00:00',10,1,1,'2025-07-06 21:38:56','2025-07-06 21:38:56',NULL),(126,'2025-07-23','08:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(127,'2025-07-23','09:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(128,'2025-07-23','14:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(129,'2025-07-23','15:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(130,'2025-07-23','18:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(131,'2025-07-23','19:00:00',18,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(132,'2025-07-23','20:00:00',18,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(133,'2025-07-23','21:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(134,'2025-07-24','08:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(135,'2025-07-24','09:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(136,'2025-07-24','14:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(137,'2025-07-24','15:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(138,'2025-07-24','18:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(139,'2025-07-24','19:00:00',18,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(140,'2025-07-24','20:00:00',18,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(141,'2025-07-24','21:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(142,'2025-07-25','08:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(143,'2025-07-25','09:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(144,'2025-07-25','14:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(145,'2025-07-25','15:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(146,'2025-07-25','18:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(147,'2025-07-25','19:00:00',18,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(148,'2025-07-25','20:00:00',18,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(149,'2025-07-25','21:00:00',16,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(150,'2025-07-26','08:00:00',10,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(151,'2025-07-26','09:00:00',10,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL),(152,'2025-07-26','14:00:00',10,1,1,'2025-07-23 22:32:49','2025-07-23 22:32:49',NULL);
/*!40000 ALTER TABLE `classes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `countries`
--

DROP TABLE IF EXISTS `countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `countries` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext COLLATE utf8mb4_unicode_ci,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_countries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `countries`
--

LOCK TABLES `countries` WRITE;
/*!40000 ALTER TABLE `countries` DISABLE KEYS */;
INSERT INTO `countries` VALUES (1,'Argentina','2025-06-02 22:27:26.174','2025-06-11 18:51:56.144',NULL),(2,'Chile','2025-06-04 18:59:34.397','2025-06-04 19:04:51.316','2025-06-05 20:00:54.126');
/*!40000 ALTER TABLE `countries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `disciplines`
--

DROP TABLE IF EXISTS `disciplines`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `disciplines` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_disciplines_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `disciplines`
--

LOCK TABLES `disciplines` WRITE;
/*!40000 ALTER TABLE `disciplines` DISABLE KEYS */;
INSERT INTO `disciplines` VALUES (1,'Crossfit','2025-06-08 13:39:26.000','2025-06-08 13:43:39.000',NULL),(2,'Crossfit','2025-06-08 19:12:05.617','2025-06-08 19:12:05.617','2025-06-08 19:12:52.332'),(3,'Fitness','2025-06-13 21:19:39.729','2025-06-13 21:19:39.729',NULL);
/*!40000 ALTER TABLE `disciplines` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gym_settings`
--

DROP TABLE IF EXISTS `gym_settings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gym_settings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `gym_id` bigint unsigned NOT NULL,
  `cancel_time_limit_minutes` int DEFAULT '60',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `gym_id` (`gym_id`),
  CONSTRAINT `gym_settings_ibfk_1` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gym_settings`
--

LOCK TABLES `gym_settings` WRITE;
/*!40000 ALTER TABLE `gym_settings` DISABLE KEYS */;
INSERT INTO `gym_settings` VALUES (1,1,60,'2025-06-25 18:48:25','2025-06-25 18:48:25',NULL);
/*!40000 ALTER TABLE `gym_settings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gyms`
--

DROP TABLE IF EXISTS `gyms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gyms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `location` longtext,
  `phone` longtext,
  `email` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `country_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_gyms_country` (`country_id`),
  KEY `idx_gyms_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_gyms_country` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`),
  CONSTRAINT `gyms_country_id_foreign` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gyms`
--

LOCK TABLES `gyms` WRITE;
/*!40000 ALTER TABLE `gyms` DISABLE KEYS */;
INSERT INTO `gyms` VALUES (1,'Hook Fitness','Av Urquiza y Bruno Morón, Coquimbito','2615752185','hookfitness@gmail.com','2025-06-05 20:20:32.649','2025-06-05 20:20:32.649',NULL,1),(2,'Hook Fitness Buenos Aires','Av Mar del Plata','2615752189','hookfitness@gmail.ar','2025-06-05 20:28:05.859','2025-06-05 20:51:44.472',NULL,1),(3,'On Fitness','Av Sarmiento 1265 Luzuriaga','2615752185','onfintess@gmail.com.ar','2025-06-13 21:40:01.455','2025-06-13 21:42:12.314',NULL,1);
/*!40000 ALTER TABLE `gyms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `packs`
--

DROP TABLE IF EXISTS `packs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `packs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `pack_name` varchar(255) NOT NULL,
  `price` decimal(8,2) NOT NULL,
  `class_quantity` int NOT NULL,
  `months` int NOT NULL,
  `gym_id` bigint unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `packs_gym_id_foreign` (`gym_id`),
  CONSTRAINT `packs_gym_id_foreign` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `packs`
--

LOCK TABLES `packs` WRITE;
/*!40000 ALTER TABLE `packs` DISABLE KEYS */;
INSERT INTO `packs` VALUES (1,'Pack Mensual x8',20000.00,8,1,1,'2025-06-15 23:44:53','2025-06-15 23:44:53',NULL),(2,'Pack Mensual x12',30000.00,12,1,1,'2025-06-16 00:39:19','2025-06-16 00:40:46',NULL);
/*!40000 ALTER TABLE `packs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'superadmin','2025-06-19 20:17:58.000','2025-06-19 20:17:58.000',NULL),(2,'admin','2025-06-19 20:18:38.000','2025-06-19 20:18:38.000',NULL),(3,'user','2025-06-19 20:18:43.000','2025-06-19 20:18:43.000',NULL);
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedule_blocks`
--

DROP TABLE IF EXISTS `schedule_blocks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schedule_blocks` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `template_id` bigint unsigned NOT NULL,
  `start_time` time NOT NULL,
  `end_time` time NOT NULL,
  `capacity` int NOT NULL,
  `discipline_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `template_id` (`template_id`),
  KEY `discipline_id` (`discipline_id`),
  CONSTRAINT `schedule_blocks_ibfk_1` FOREIGN KEY (`template_id`) REFERENCES `schedule_templates` (`id`),
  CONSTRAINT `schedule_blocks_ibfk_2` FOREIGN KEY (`discipline_id`) REFERENCES `disciplines` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedule_blocks`
--

LOCK TABLES `schedule_blocks` WRITE;
/*!40000 ALTER TABLE `schedule_blocks` DISABLE KEYS */;
INSERT INTO `schedule_blocks` VALUES (1,1,'08:00:00','09:00:00',20,1),(2,1,'09:00:00','10:00:00',15,1),(3,1,'10:00:00','11:00:00',10,1),(4,1,'15:00:00','16:00:00',15,1),(6,2,'08:00:00','09:00:00',16,1),(7,2,'09:00:00','10:00:00',16,1),(8,2,'14:00:00','15:00:00',16,1),(9,2,'15:00:00','16:00:00',16,1),(10,2,'18:00:00','19:00:00',16,1),(11,2,'19:00:00','20:00:00',18,1),(12,2,'20:00:00','21:00:00',18,1),(13,2,'21:00:00','22:00:00',16,1),(14,3,'08:00:00','09:00:00',16,1),(15,3,'09:00:00','10:00:00',16,1),(16,3,'14:00:00','15:00:00',16,1),(17,3,'15:00:00','16:00:00',16,1),(18,3,'18:00:00','19:00:00',16,1),(19,3,'19:00:00','20:00:00',18,1),(20,3,'20:00:00','21:00:00',18,1),(21,3,'21:00:00','22:00:00',16,1),(22,4,'08:00:00','09:00:00',16,1),(23,4,'09:00:00','10:00:00',16,1),(24,4,'14:00:00','15:00:00',16,1),(25,4,'15:00:00','16:00:00',16,1),(26,4,'18:00:00','19:00:00',16,1),(27,4,'19:00:00','20:00:00',18,1),(28,4,'20:00:00','21:00:00',18,1),(29,4,'21:00:00','22:00:00',16,1),(30,5,'08:00:00','09:00:00',16,1),(31,5,'09:00:00','10:00:00',16,1),(32,5,'14:00:00','15:00:00',16,1),(33,5,'15:00:00','16:00:00',16,1),(34,5,'18:00:00','19:00:00',16,1),(35,5,'19:00:00','20:00:00',18,1),(36,5,'20:00:00','21:00:00',18,1),(37,5,'21:00:00','22:00:00',16,1),(38,6,'08:00:00','09:00:00',10,1),(39,6,'09:00:00','10:00:00',10,1),(40,6,'14:00:00','15:00:00',10,1),(41,1,'18:00:00','19:00:00',16,1),(42,1,'19:00:00','20:00:00',18,1);
/*!40000 ALTER TABLE `schedule_blocks` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schedule_templates`
--

DROP TABLE IF EXISTS `schedule_templates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schedule_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `gym_id` bigint unsigned NOT NULL,
  `day` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `gym_id` (`gym_id`),
  CONSTRAINT `schedule_templates_ibfk_1` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schedule_templates`
--

LOCK TABLES `schedule_templates` WRITE;
/*!40000 ALTER TABLE `schedule_templates` DISABLE KEYS */;
INSERT INTO `schedule_templates` VALUES (1,1,'Lunes'),(2,1,'Martes'),(3,1,'Miércoles'),(4,1,'Jueves'),(5,1,'Viernes'),(6,1,'Sábado');
/*!40000 ALTER TABLE `schedule_templates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `lastname` longtext,
  `gender` longtext,
  `phone` longtext,
  `email` varchar(191) DEFAULT NULL,
  `movil_phone` longtext,
  `birth_date` longtext,
  `dni` longtext,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` longtext,
  `remember_token` varchar(100) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `gym_id` bigint unsigned DEFAULT NULL,
  `role_id` bigint NOT NULL DEFAULT '3',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_email` (`email`),
  KEY `fk_users_gym` (`gym_id`),
  KEY `idx_users_deleted_at` (`deleted_at`),
  KEY `fk_role` (`role_id`),
  CONSTRAINT `fk_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `fk_users_gym` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`),
  CONSTRAINT `users_gym_id_foreign` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Juan','Charparin','Masculino','2615752185','juan.charpa@example.com','2615123456','2000-08-30','42862648',NULL,'123456',NULL,'2025-06-08 13:07:45.000','2025-06-08 13:22:16.000','2025-07-11 18:56:05.085',1,1),(2,'Rocio','Charparin','Femenino','2615752185','rocio.charpa@example.com','2615752185','1998-04-30','41083888',NULL,'123456',NULL,'2025-06-15 19:29:13.323','2025-06-15 19:35:43.595',NULL,1,3),(3,'Emiliano','Guzman','Masculino','2615752185','emiliano.guzman@example.com',NULL,'2000-07-02','41083888',NULL,'$2a$10$xfPRSOyB2L1T2Px6zbXgk.Rfr7gycIQC4lVsrNGygnZlbogoNG1aK',NULL,'2025-06-16 16:48:30.909','2025-07-11 20:44:18.233',NULL,1,3),(4,'Analia','Moreno','Femenino','2615752185','analia.moreno@example.com',NULL,'1970-11-22','42862649',NULL,'',NULL,'2025-06-18 22:23:40.043','2025-07-11 19:15:21.955',NULL,1,3),(7,'Agustin','Narvaez','Masculino','2613739846','agustin.narvaez@example.com',NULL,'2001-05-14','43305826',NULL,'$2a$10$91FzNAKgoz.VTBuFuaISDegJvaBUqlrZBYdigxOFsk03j9jx.YXvW',NULL,'2025-06-22 17:51:33.945','2025-06-22 17:51:33.945',NULL,1,3),(8,'Juan','Android','Masculino','2615752184','jccharparin@gmail.com',NULL,'2000-08-31','42862648',NULL,'$2a$10$FqGL2BLLkUeNbAtEnoFUFuV1ghOXcatrgyYptZN/hpw0ynuPEaf.O',NULL,'2025-06-30 20:59:57.279','2025-06-30 20:59:57.279',NULL,1,3),(9,'Lautaro','Sanchez','Masculino','2615752184','lautaromanuels@gmail.com',NULL,'2000-11-14','42862648',NULL,'$2a$10$XG/Z5iwGG.8874acVS1gKOmjatfMnYnxqFptWRMyz.ew4KQQ9jWaK',NULL,'2025-07-01 19:29:20.890','2025-07-01 19:29:20.890',NULL,1,3),(10,'Franco','Antunez','Masculino','2615752184','francoanutnez@gmail.com',NULL,'2000-10-02','42862648',NULL,'$2a$10$FJv647.BskAPJcMffIrsOe1.cH.20Jjs8m/GC5SmcfizIm8peXYKG',NULL,'2025-07-01 19:30:29.426','2025-07-01 19:30:29.426',NULL,1,3),(11,'Pelito','Gaboardi','Masculino','2615752184','pelitogaboardi@gmail.com',NULL,'1992-08-09','42862648',NULL,'$2a$10$ZsYjZH5dqZtnhdlJg3Yq.evpUvKbpdmlMixsIhw1qEhmAbhnWHS3m',NULL,'2025-07-01 19:32:14.861','2025-07-01 19:32:14.861',NULL,1,3),(12,'Gaston','Guevara','Masculino','2615752184','gaston_guevara@gmail.com',NULL,'1992-08-10','42862648',NULL,'$2a$10$0aJ.B5cfyi5p45w1l3OGs.USr4wleNq2D4H1ljTthLqPfaTCQurl2',NULL,'2025-07-01 21:40:42.419','2025-07-01 21:40:42.419',NULL,1,3),(13,'Esteban','Maldonado','Masculino','2615752184','esteban_maldonado@gmail.com',NULL,'1991-07-02','42862648',NULL,'$2a$10$DNkTTOkou/eJjmNgAHFK9eisAI6UUaZhIyCvWlmnEzmet9pGL2R3O',NULL,'2025-07-01 21:41:26.300','2025-07-01 21:41:26.300',NULL,1,3),(14,'Nahuel','Bottani','Masculino','2615752184','nahuelbotanni@gmail.com','2615752184','2000-08-31','42862648',NULL,'',NULL,'2025-07-07 20:10:31.544','2025-07-07 20:10:31.544',NULL,1,2),(15,'Nahuel','Bottani','Masculino','2615752184','nahuel_botanni@gmail.com',NULL,'2000-08-31','42862648',NULL,'$2a$10$9PNe7L6ki29.f5ax4Eidn.IR2n.LTJcNh0Plmw7mKGygxIEL0op8m',NULL,'2025-07-07 20:34:31.674','2025-07-07 20:34:31.674',NULL,1,2);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users_packs`
--

DROP TABLE IF EXISTS `users_packs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users_packs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `start_date` date NOT NULL,
  `expiration_date` date NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `gym_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `pack_id` bigint unsigned NOT NULL,
  `discipline_id` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `users_packs_gym_id_foreign` (`gym_id`),
  KEY `users_packs_user_id_foreign` (`user_id`),
  KEY `users_packs_pack_id_foreign` (`pack_id`),
  KEY `users_packs_discipline_id_foreign` (`discipline_id`),
  CONSTRAINT `users_packs_discipline_id_foreign` FOREIGN KEY (`discipline_id`) REFERENCES `disciplines` (`id`) ON DELETE CASCADE,
  CONSTRAINT `users_packs_gym_id_foreign` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`) ON DELETE CASCADE,
  CONSTRAINT `users_packs_pack_id_foreign` FOREIGN KEY (`pack_id`) REFERENCES `packs` (`id`) ON DELETE CASCADE,
  CONSTRAINT `users_packs_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users_packs`
--

LOCK TABLES `users_packs` WRITE;
/*!40000 ALTER TABLE `users_packs` DISABLE KEYS */;
INSERT INTO `users_packs` VALUES (1,'2025-06-15','2025-07-15',0,1,1,1,1,'2025-06-16 00:47:28','2025-07-07 00:18:41',NULL),(2,'2025-06-16','2025-07-16',0,1,2,2,1,'2025-06-16 01:20:31','2025-07-16 22:28:35',NULL),(3,'2025-06-21','2025-07-21',0,1,4,2,1,'2025-06-22 02:01:24','2025-07-20 23:40:15',NULL),(4,'2025-06-21','2025-07-21',0,1,7,1,1,'2025-06-22 21:36:46','2025-07-07 00:32:33',NULL),(5,'2025-06-21','2025-07-21',0,1,3,1,1,'2025-06-22 21:38:45','2025-06-22 21:38:45',NULL),(6,'2025-06-30','2025-07-31',1,1,3,2,1,'2025-07-02 00:33:24','2025-07-02 00:33:24',NULL),(7,'2025-07-01','2025-08-01',1,1,9,2,1,'2025-07-02 00:37:50','2025-07-02 00:37:50',NULL),(8,'2025-07-01','2025-08-01',1,1,8,2,1,'2025-07-02 00:38:27','2025-07-02 00:38:27',NULL),(9,'2025-07-01','2025-08-01',1,1,10,2,1,'2025-07-02 00:38:38','2025-07-02 00:38:38',NULL),(10,'2025-07-01','2025-08-01',1,1,11,2,1,'2025-07-02 00:38:42','2025-07-02 00:38:42',NULL),(11,'2025-07-01','2025-08-01',1,1,12,2,1,'2025-07-02 00:41:55','2025-07-02 00:41:55',NULL),(12,'2025-07-01','2025-08-01',1,1,13,2,1,'2025-07-02 00:42:01','2025-07-02 00:42:01',NULL),(13,'2025-07-16','2025-08-16',1,1,2,1,1,'2025-07-16 23:13:03','2025-07-16 23:13:03',NULL),(14,'2025-07-20','2025-08-20',1,1,4,1,3,'2025-07-20 23:43:21','2025-07-20 23:43:21',NULL);
/*!40000 ALTER TABLE `users_packs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `waitlists`
--

DROP TABLE IF EXISTS `waitlists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `waitlists` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `class_id` bigint unsigned NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `class_id` (`class_id`),
  CONSTRAINT `waitlists_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `waitlists_ibfk_2` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `waitlists`
--

LOCK TABLES `waitlists` WRITE;
/*!40000 ALTER TABLE `waitlists` DISABLE KEYS */;
INSERT INTO `waitlists` VALUES (1,13,83,'2025-07-01 21:42:32','2025-07-02 00:42:32',NULL),(2,13,88,'2025-07-06 18:51:06','2025-07-06 21:51:06',NULL);
/*!40000 ALTER TABLE `waitlists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wods`
--

DROP TABLE IF EXISTS `wods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wods` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext,
  `description` longtext,
  `type` longtext,
  `duration` bigint DEFAULT NULL,
  `level` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `gym_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_wods_gym` (`gym_id`),
  KEY `idx_wods_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_wods_gym` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`),
  CONSTRAINT `wods_gym_id_foreign` FOREIGN KEY (`gym_id`) REFERENCES `gyms` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wods`
--

LOCK TABLES `wods` WRITE;
/*!40000 ALTER TABLE `wods` DISABLE KEYS */;
INSERT INTO `wods` VALUES (1,'Crossfit','Cardio al 100','AMRAP',18,'Avanzado','2025-06-08 19:14:29.659','2025-06-08 19:15:22.996',NULL,1),(2,'Heavy','Heavy without air','For Time',15,'Avanzado','2025-06-15 12:52:56.235','2025-06-15 12:55:05.505',NULL,1);
/*!40000 ALTER TABLE `wods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'wod-db-go'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-24 19:10:47
