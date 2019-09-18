-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema sample_db
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `sample_db` ;

-- -----------------------------------------------------
-- Schema sample_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `sample_db` DEFAULT CHARACTER SET utf8 ;
USE `sample_db` ;

-- -----------------------------------------------------
-- Table `sample_db`.`user_statuses`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sample_db`.`user_statuses` ;

CREATE TABLE IF NOT EXISTS `sample_db`.`user_statuses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `status` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sample_db`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sample_db`.`users` ;

CREATE TABLE IF NOT EXISTS `sample_db`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(16) NOT NULL,
  `password` VARCHAR(100) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `status_id_idx` (`status_id` ASC),
  UNIQUE INDEX `username_UNIQUE` (`username` ASC),
  CONSTRAINT `status_id`
    FOREIGN KEY (`status_id`)
    REFERENCES `sample_db`.`user_statuses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `sample_db`.`posts`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sample_db`.`posts` ;

CREATE TABLE IF NOT EXISTS `sample_db`.`posts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `posted_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `user_id_idx` (`user_id` ASC),
  CONSTRAINT `user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `sample_db`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sample_db`.`article_statuses`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sample_db`.`article_statuses` ;

CREATE TABLE IF NOT EXISTS `sample_db`.`article_statuses` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `status` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `sample_db`.`articles`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `sample_db`.`articles` ;

CREATE TABLE IF NOT EXISTS `sample_db`.`articles` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(45) NOT NULL,
  `content` VARCHAR(1000) NOT NULL,
  `post_id` INT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `user_id` INT NOT NULL,
  `article_statuse_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `post_id_idx` (`post_id` ASC),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  INDEX `fk_articles_users1_idx` (`user_id` ASC),
  INDEX `fk_articles_article_statuses1_idx` (`article_statuse_id` ASC),
  CONSTRAINT `post_id`
    FOREIGN KEY (`post_id`)
    REFERENCES `sample_db`.`posts` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_articles_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `sample_db`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_articles_article_statuses1`
    FOREIGN KEY (`article_statuse_id`)
    REFERENCES `sample_db`.`article_statuses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
