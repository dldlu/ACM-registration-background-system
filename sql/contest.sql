/*
 **************************************
  MySQL - 8.0.23 : Database - contest
 **************************************
 */

CREATE DATABASE `contest`;

USE `contest`;

CREATE TABLE admin_info (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `role_id` INT NOT NULL COMMENT '角色id',
  `username` VARCHAR(64) NOT NULL COMMENT '管理员名字',
  `password` VARCHAR(64) NOT NULL COMMENT '管理员密码',
  `email` VARCHAR(64) NOT NULL COMMENT '管理员邮箱',
  `status` TINYINT NOT NULL DEFAULT '1' COMMENT '状态位',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 26 DEFAULT CHARSET = utf8;

CREATE TABLE information (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `student_number` VARCHAR(12) NOT NULL COMMENT '学号',
  `student_name` VARCHAR(12) DEFAULT NULL COMMENT '姓名',
  `major_class` VARCHAR(12) DEFAULT NULL COMMENT '专业班级',
  `sex` VARCHAR(2) NOT NULL COMMENT '性别',
  `qq` VARCHAR(12) NOT NULL COMMENT 'qq号',
  `phone` VARCHAR(11) NOT NULL COMMENT '手机号',
  `email` VARCHAR(64) NOT NULL COMMENT '邮箱',
  `year` INT NOT NULL COMMENT '报名年份',
  `last_ip` VARCHAR(24) NOT NULL COMMENT '提交ip',
  `sign_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '提交时间',
  `status` TINYINT NOT NULL DEFAULT '1' COMMENT '状态位',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 26 DEFAULT CHARSET = utf8;

CREATE TABLE log (
  `id` INT NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `state` TINYINT NOT NULL COMMENT '日志状态',
  `method` TINYINT NOT NULL COMMENT '操作方法',
  `content` VARCHAR(128) NOT NULL COMMENT '内容',
  `operator_id` INT NOT NULL COMMENT '操作者id',
  `operator_ip` VARCHAR(16) NOT NULL COMMENT '操作者ip',
  `created_at` datetime NOT NULL COMMENT '生成时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 26 DEFAULT CHARSET = utf8;

INSERT INTO `contest`.`admin_info` (`id`, `role_id`, `username`, `password`, `email`, `status`) VALUES ('1', '1', '12345678', '$2a$12$hPxBJf0qPAiXgpTaYCEsUeBKabNoWt2j.qFTNu7mVknzRN7Zdwo6W', '', '1');
