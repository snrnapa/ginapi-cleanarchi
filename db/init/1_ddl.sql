SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `user_manage` DEFAULT CHARACTER SET utf8mb4 ;
USE `user_manage` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `user_manage`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `user_manage`.`user` (
  `id` VARCHAR(128) NOT NULL COMMENT 'ユーザID',
  `firstname` VARCHAR(64) NOT NULL COMMENT '1stName',
  `lastname` VARCHAR(64) NOT NULL COMMENT '2ndName',

  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ユーザ';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
