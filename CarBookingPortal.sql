-- CarBookingPortal.sql
-- MySQL dump 10.13  Distrib 5.7.31, for Linux (x86_64)
--
-- Host: localhost    Database: car_book
-- ------------------------------------------------------
-- Server version	5.7.31-0ubuntu0.18.04.1-log

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
-- Table structure for table `cb_cabs`
--

DROP TABLE IF EXISTS `cb_cabs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cb_cabs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cab_no` varchar(20) NOT NULL,
  `status` varchar(10) NOT NULL,
  `location` varchar(45) NOT NULL,
  `lat` float NOT NULL,
  `lng` float NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=big5;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cb_cabs`
--

LOCK TABLES `cb_cabs` WRITE;
/*!40000 ALTER TABLE `cb_cabs` DISABLE KEYS */;
INSERT INTO `cb_cabs` VALUES (1,'MH12DS1211','Active','Ambegaon',19.1132,73.7327,'2020-09-08 15:54:21','2020-09-07 18:31:57',NULL),(2,'MH12DS1212','Active','Aundh',18.5602,73.8031,'2020-09-08 15:54:21','2020-09-07 18:31:57',NULL),(3,'MH12DS1213','Active','Baner',18.559,73.7868,'2020-09-08 15:54:21','2020-09-07 18:31:57',NULL),(4,'MH12DS1214','Active','Dhankawadi',18.4655,73.8547,'2020-09-08 15:54:21','2020-09-08 15:54:21',NULL),(5,'MH12DS1215','Active','Hadapsar',18.5089,73.9259,'2020-09-08 15:54:21','2020-09-08 15:54:21',NULL),(6,'MH12DS1216','Active','Katraj',18.4529,73.8652,'2020-09-08 15:59:21',NULL,NULL),(7,'MH12DS1217','Active','Koregaon Park',18.5362,73.894,'2020-09-08 16:00:29',NULL,NULL);
/*!40000 ALTER TABLE `cb_cabs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cb_drivers`
--

DROP TABLE IF EXISTS `cb_drivers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cb_drivers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `mobile` varchar(15) NOT NULL,
  `cab_id` int(11) NOT NULL,
  `status` varchar(10) NOT NULL DEFAULT 'Active',
  `license_no` varchar(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=big5;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cb_drivers`
--

LOCK TABLES `cb_drivers` WRITE;
/*!40000 ALTER TABLE `cb_drivers` DISABLE KEYS */;
INSERT INTO `cb_drivers` VALUES (1,'Heir Apparel','9843543399',1,'Active','212092091','2020-09-08 16:14:14',NULL,NULL),(2,'Teri Dactyl','9832121232',2,'Active','121','2020-09-08 16:15:47',NULL,NULL),(3,'Olive Yew','8432923238',3,'Active','2198291092','2020-09-08 16:16:19',NULL,NULL),(4,'Aida Bugg','8432923238',4,'Active','2198291092','2020-09-08 16:16:32',NULL,NULL),(5,'Allie Grater','9898987834',5,'Active','21819281982','2020-09-08 16:18:41',NULL,NULL),(6,'Joe V. Awl','7823828323',6,'Active','832989283','2020-09-08 16:18:41',NULL,NULL),(7,'Ben Dover','8734893289',7,'Active','832983928','2020-09-08 16:18:41',NULL,NULL);
/*!40000 ALTER TABLE `cb_drivers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `cb_rides`
--

DROP TABLE IF EXISTS `cb_rides`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `cb_rides` (
  `ride_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `cab_no` varchar(20) NOT NULL,
  `driver_id` int(11) NOT NULL,
  `booking_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `pickup_loc` varchar(60) NOT NULL,
  `drop_loc` varchar(60) NOT NULL,
  PRIMARY KEY (`ride_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=big5;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cb_rides`
--

LOCK TABLES `cb_rides` WRITE;
/*!40000 ALTER TABLE `cb_rides` DISABLE KEYS */;
INSERT INTO `cb_rides` VALUES (1,1,'MH12DS1211',1,3423,'2020-09-10 11:54:26',NULL,'Baner','Aundh'),(2,1,'MH12DS1214',4,7873,'2020-09-10 11:56:24',NULL,'Aundh','Katraj'),(3,1,'MH12DS1214',4,3234,'2020-09-10 11:57:04',NULL,'Katraj','Koregaon Park'),(4,2,'MH12DS1212',2,2342,'2020-09-10 11:58:30',NULL,'Dhankawadi','Hadapsar');
/*!40000 ALTER TABLE `cb_rides` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-09-10 19:59:09
