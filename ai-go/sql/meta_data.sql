-- MySQL dump 10.13  Distrib 5.5.53, for Win32 (AMD64)
--
-- Host: localhost    Database: meta_data
-- ------------------------------------------------------
-- Server version	8.0.13

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `admin_user`
--

DROP TABLE IF EXISTS `admin_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(50) NOT NULL DEFAULT 'editor',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='后台用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin_user`
--

LOCK TABLES `admin_user` WRITE;
/*!40000 ALTER TABLE `admin_user` DISABLE KEYS */;
INSERT INTO `admin_user` VALUES (1,'admin','$2a$10$WjTaUW3lgFJWsT.0IvGBZeBQWfwNKjkrn8hI9rEypeeDWl4tr7vS6','admin',0,'2025-05-23 17:08:10','2025-05-23 17:08:10'),(2,'test','$2a$10$Ye0hFLPsyyYFo2jKuu5ZzuL86W50tA0.dqCnbMpjoztdcHZqw886a','user',0,'2025-05-28 11:50:43','2025-05-28 11:50:43');
/*!40000 ALTER TABLE `admin_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meta_attr`
--

DROP TABLE IF EXISTS `meta_attr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meta_attr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '属性ID',
  `attr_name` varchar(200) NOT NULL,
  `attr_type` enum('string','number','time','bool') NOT NULL DEFAULT 'string',
  `attr_desc` text,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='事件属性元数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meta_attr`
--

LOCK TABLES `meta_attr` WRITE;
/*!40000 ALTER TABLE `meta_attr` DISABLE KEYS */;
INSERT INTO `meta_attr` VALUES (1,'channel','string','来源渠道',0,'2025-05-23 17:10:02','2025-05-23 17:10:02'),(2,'register_time','time','注册时间',0,'2025-05-23 17:10:16','2025-05-26 14:28:19'),(3,'os_type','string','系统类型',1,'2025-05-23 17:10:47','2025-05-28 15:27:18'),(4,'max','number','最大值',0,'2025-05-23 17:11:18','2025-05-23 17:11:18'),(5,'min','number','最小值',0,'2025-05-23 17:11:27','2025-05-23 17:11:27'),(6,'rounds','number','对局回合数',0,'2025-05-23 17:12:18','2025-05-23 17:12:18'),(7,'uid','string','用户id',0,'2025-05-23 18:48:07','2025-05-23 18:48:07'),(8,'region','string','省市',0,'2025-05-23 18:52:58','2025-05-23 18:52:58'),(9,'recharge','number','充值金额',0,'2025-05-28 11:31:16','2025-05-28 11:31:16'),(10,'game_type','string','游戏玩法',0,'2025-05-28 11:32:45','2025-05-28 11:32:45'),(11,'ip','string','ip地址',0,'2025-05-28 11:35:47','2025-05-28 11:35:47'),(12,'event_time','time','打点时间',0,'2025-05-28 11:36:19','2025-05-28 14:27:27'),(16,'aaa','time','aaa',0,'2025-05-28 14:57:46','2025-05-28 15:19:19'),(17,'bbb','bool','bbb',0,'2025-05-28 15:08:58','2025-05-28 15:19:17');
/*!40000 ALTER TABLE `meta_attr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meta_event`
--

DROP TABLE IF EXISTS `meta_event`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meta_event` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '事件ID',
  `event_name` varchar(200) NOT NULL,
  `event_desc` text,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='埋点事件元数据表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meta_event`
--

LOCK TABLES `meta_event` WRITE;
/*!40000 ALTER TABLE `meta_event` DISABLE KEYS */;
INSERT INTO `meta_event` VALUES (1,'login','登录',1,'2025-05-23 17:08:23','2025-05-28 15:26:34'),(3,'dice','掷骰子',0,'2025-05-23 17:08:49','2025-05-26 14:26:46'),(4,'event_test1','test1',0,'2025-05-23 18:12:00','2025-05-23 18:12:00'),(5,'event_test1','test1',0,'2025-05-23 18:44:40','2025-05-23 18:44:40'),(6,'event_test2','test2',0,'2025-05-26 10:56:17','2025-05-26 10:56:17'),(7,'event_test3','test3',0,'2025-05-26 10:56:34','2025-05-26 10:56:34'),(8,'event_test4','test4',0,'2025-05-26 10:57:12','2025-05-26 10:57:12'),(9,'event_test5','test5',0,'2025-05-26 10:57:12','2025-05-26 10:57:12'),(10,'event_test6','test6',0,'2025-05-26 10:57:12','2025-05-26 10:57:12'),(11,'event_test7','test7',0,'2025-05-26 10:57:12','2025-05-26 10:57:12'),(12,'event_test8','test8',0,'2025-05-26 14:48:28','2025-05-26 14:48:28'),(13,'event_test9','test9',0,'2025-05-26 14:48:28','2025-05-26 14:48:28'),(14,'event_test10','test10',0,'2025-05-26 14:48:28','2025-05-26 14:48:28'),(15,'event_test11','test11',0,'2025-05-26 14:48:28','2025-05-28 14:28:02');
/*!40000 ALTER TABLE `meta_event` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `meta_relation`
--

DROP TABLE IF EXISTS `meta_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `meta_relation` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '关联ID',
  `event_id` bigint(20) unsigned NOT NULL COMMENT '事件ID（关联metadata_event.id）',
  `attr_id` bigint(20) unsigned NOT NULL COMMENT '属性ID（关联metadata_attribute.id）',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='事件与属性关联表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `meta_relation`
--

LOCK TABLES `meta_relation` WRITE;
/*!40000 ALTER TABLE `meta_relation` DISABLE KEYS */;
INSERT INTO `meta_relation` VALUES (4,3,4,'2025-05-23 17:13:29',0),(5,3,5,'2025-05-23 17:13:29',0),(6,3,6,'2025-05-23 17:13:29',0),(7,3,3,'2025-05-23 17:58:53',1),(8,4,8,'2025-05-23 18:52:58',0),(9,4,3,'2025-05-23 19:00:09',1),(10,6,3,'2025-05-26 11:02:05',1),(11,7,3,'2025-05-26 11:04:31',1),(12,8,3,'2025-05-26 11:04:34',1),(13,9,3,'2025-05-26 11:04:38',1),(14,1,1,'2025-05-26 11:19:38',1),(15,1,2,'2025-05-26 11:19:38',1),(16,1,3,'2025-05-26 11:19:38',1),(17,10,1,'2025-05-26 11:26:50',0),(18,10,2,'2025-05-26 11:26:50',0),(19,10,3,'2025-05-26 11:26:50',1),(20,11,1,'2025-05-26 11:26:50',0),(21,11,2,'2025-05-26 11:26:50',0),(22,11,3,'2025-05-26 11:26:50',1),(23,9,4,'2025-05-26 11:27:04',0),(24,9,5,'2025-05-26 11:27:04',0),(25,15,11,'2025-05-28 11:51:13',0),(26,14,6,'2025-05-28 11:52:30',0);
/*!40000 ALTER TABLE `meta_relation` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-28 15:28:50
