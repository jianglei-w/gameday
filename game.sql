-- MySQL dump 10.13  Distrib 8.1.0, for macos13.3 (arm64)
--
-- Host: sh-cynosdbmysql-grp-r4j6c55k.sql.tencentcdb.com    Database: game
-- ------------------------------------------------------
-- Server version	5.7.18-cynos-log

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
SET @MYSQLDUMP_TEMP_LOG_BIN = @@SESSION.SQL_LOG_BIN;
SET @@SESSION.SQL_LOG_BIN= 0;

--
-- GTID state at the beginning of the backup 
--

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ '893e2bc5-5bb2-11ee-a1cc-b8cef6203cac:1-1036';

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `admin` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL,
  `password` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_admin_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,NULL,NULL,NULL,'jianglei','jianglei');
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `complete`
--

DROP TABLE IF EXISTS `complete`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `complete` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `question_id` bigint(20) unsigned DEFAULT NULL,
  `score` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_complete_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `complete`
--

LOCK TABLES `complete` WRITE;
/*!40000 ALTER TABLE `complete` DISABLE KEYS */;
INSERT INTO `complete` VALUES (8,'2023-10-23 16:00:18.058','2023-10-23 16:00:18.058',NULL,126,5,100),(9,'2023-10-23 16:02:37.650','2023-10-23 16:02:37.650',NULL,126,6,99),(10,'2023-10-23 16:02:58.089','2023-10-23 16:02:58.089',NULL,126,9,10);
/*!40000 ALTER TABLE `complete` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `game`
--

DROP TABLE IF EXISTS `game`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `game` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `people` bigint(20) unsigned DEFAULT NULL,
  `status` bigint(20) unsigned DEFAULT NULL,
  `group_id` bigint(20) unsigned DEFAULT NULL,
  `group_name` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_game_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `game`
--

LOCK TABLES `game` WRITE;
/*!40000 ALTER TABLE `game` DISABLE KEYS */;
INSERT INTO `game` VALUES (9,'2023-10-21 00:05:10.160','2023-10-22 22:04:01.117',NULL,'test',11,1,1,NULL);
/*!40000 ALTER TABLE `game` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `group` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_group_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;
INSERT INTO `group` VALUES (1,'2023-10-17 16:35:28.233','2023-10-17 16:35:28.233',NULL,'题目组1'),(2,'2023-10-17 16:37:32.557','2023-10-17 16:37:32.557',NULL,'云计算题目组');
/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `question`
--

DROP TABLE IF EXISTS `question`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `question` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(191) DEFAULT NULL,
  `text` varchar(191) DEFAULT NULL,
  `answer` varchar(191) DEFAULT NULL,
  `grade` bigint(20) DEFAULT NULL,
  `level` varchar(191) DEFAULT NULL,
  `group_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_question_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question`
--

LOCK TABLES `question` WRITE;
/*!40000 ALTER TABLE `question` DISABLE KEYS */;
INSERT INTO `question` VALUES (5,'2023-10-17 10:36:50.618','2023-10-17 10:36:50.618',NULL,'jianglei','jiangleiw','123',100,'1',0),(6,'2023-10-19 10:29:37.684','2023-10-20 22:46:29.391',NULL,'题目11','<p>题目2</p>','nice',99,'2',0),(8,'2023-10-19 12:19:49.982','2023-10-20 22:53:55.188','2023-10-20 22:54:18.793','hhahahah','<p>Cloud Computing</p>','nice',100,'3',0),(9,'2023-10-20 22:54:58.186','2023-10-20 22:54:58.186',NULL,'题目1','<h1 id=\"20\">20</h1>','nice',10,'1',0);
/*!40000 ALTER TABLE `question` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `question_groups`
--

DROP TABLE IF EXISTS `question_groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `question_groups` (
  `group_id` bigint(20) unsigned NOT NULL,
  `question_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`group_id`,`question_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `question_groups`
--

LOCK TABLES `question_groups` WRITE;
/*!40000 ALTER TABLE `question_groups` DISABLE KEYS */;
INSERT INTO `question_groups` VALUES (1,5),(1,6),(1,9),(2,6),(2,7);
/*!40000 ALTER TABLE `question_groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `hashcode` varchar(191) DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL,
  `status` bigint(20) unsigned DEFAULT NULL,
  `game_id` bigint(20) unsigned DEFAULT NULL,
  `score` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (126,'2023-10-21 00:05:10.205','2023-10-22 23:21:00.817',NULL,'9caa8dd2-542e-4f08-b3e9-c9f2d43c7bdd','jianglei',2,9,0),(127,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'85848004-b10b-47cb-abbb-5e8acdbd02c2','',0,9,0),(128,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'31e56382-a0fc-4581-8d64-a8249a6aed3e','',0,9,0),(129,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'858418d4-aab8-4569-ba3f-66b619ac37fb','',0,9,0),(130,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'a77ee287-4c5d-49d6-ae0f-8bbe0f7b8e8f','',0,9,0),(131,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'5d080593-b5cf-480e-a49b-4c31f877b03a','',0,9,0),(132,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'ebb7661a-e0f9-42b8-9c03-ccf4385607b5','',0,9,0),(133,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'6993ce02-8ec6-4319-aaf4-1f7aae191fec','',0,9,0),(134,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'6bf7d899-bc9c-41ff-ba2b-9c169953c04e','',0,9,0),(135,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'4fb44797-1d93-4bda-acb6-565c14486f54','',0,9,0),(136,'2023-10-21 00:05:10.205','2023-10-21 00:05:10.205',NULL,'3357cf69-ac85-4657-91e8-e4eba7d64946','',0,9,0);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
SET @@SESSION.SQL_LOG_BIN = @MYSQLDUMP_TEMP_LOG_BIN;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-10-24 14:11:13
