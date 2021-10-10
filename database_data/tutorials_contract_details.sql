-- MySQL dump 10.13  Distrib 8.0.25, for Win64 (x86_64)
--
-- Host: localhost    Database: tutorials
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `contract_details`
--

DROP TABLE IF EXISTS `contract_details`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `contract_details` (
  `cont_id` varchar(50) DEFAULT NULL,
  `full_name` varchar(50) DEFAULT NULL,
  `contract_id` varchar(50) DEFAULT NULL,
  `contract_name` varchar(50) DEFAULT NULL,
  `component_id` varchar(50) DEFAULT NULL,
  `component_name` varchar(50) DEFAULT NULL,
  `party_role_id` varchar(50) DEFAULT NULL,
  `party_role_name` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contract_details`
--

LOCK TABLES `contract_details` WRITE;
/*!40000 ALTER TABLE `contract_details` DISABLE KEYS */;
INSERT INTO `contract_details` VALUES ('196-ukd-224','Suwanto','26H59R94','adreinan5','2D19LE934BM58709','Seamless','394KGYT15911','DR.G WHITENING BOOSTER AMPLE'),('196-ukd-224','Suwanto','26H59R94','adreinan5','2D19LE934BM58709','Seamless','175VGVO39188','ropinirole hydrochloride'),('196-ukd-224','Suwanto','26H59R94','adreinan5','2D19LE934BM58709','Seamless','639ZFZQ81718','METOPROLOL SUCCINATE'),('196-ukd-224','Suwanto','26H59R94','adreinan5','9M17FE108UX17089','Fundamental','733AMOG05740','Dentastat'),('196-ukd-224','Suwanto','26H59R94','adreinan5','9M17FE108UX17089','Fundamental','574MSNN26003','Methylphenidate Hydrochloride'),('196-ukd-224','Suwanto','26H59R94','adreinan5','9M35YS274NT91347','real-time','538EZQZ47562','ACYCLOVIR'),('196-ukd-224','Suwanto','21J82N85','jibbettg','1T98DX206YS86646','multi-tasking','091GFVN11369','ISENTRESS'),('196-ukd-224','Suwanto','21J82N85','jibbettg','1T98DX206YS86646','multi-tasking','409JPKQ49517','ARNICAID'),('196-ukd-224','Suwanto','21J82N85','jibbettg','9V84VC793HW11896','productivity','625HAWJ84355','METHADOSE'),('496-wbk-743','Sugianto','57E28L58','dlongstaffej','6P49AZ288MX69305','incremental','',''),('496-wbk-743','Sugianto','25G00P10','fsharplessk','0V90RE664OY43466','Reverse-engineered','372TTFW58921','NITROGEN'),('496-wbk-743','Sugianto','25G00P10','fsharplessk','0V90RE664OY43466','Reverse-engineered','419GOBG67519','Vida Mia Hand Sanitizer'),('116-jwo-611','Sumanto','36X50T08','mtruslerr','','','',''),('116-jwo-611','Sumanto','25N66Q39','lcanets','7D42ST109WB21925','application','176YJXY70556','Neutrogena Oil Free A'),('116-jwo-611','Sumanto','25N66Q39','lcanets','7D42ST109WB21925','application','647PUIS82355','HEADACHE/MIGRAINE RELIEF'),('116-jwo-611','Sumanto','25N66Q39','lcanets','7D42ST109WB21925','application','756AOUB29039','Mint'),('693-gux-907','Suwamto','','','','','','');
/*!40000 ALTER TABLE `contract_details` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-10-10 19:36:19
