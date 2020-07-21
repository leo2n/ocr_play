CREATE DATABASE IF NOT EXISTS db_play;

USE db_play;

DROP TABLE IF EXISTS ocr_history;

CREATE TABLE IF NOT EXISTS ocr_history (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL , /* 必须和主键一起使用*/
    userId VARCHAR(128) , /* 用户名默认设置为-1吧*/
    result TEXT ,
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    path VARCHAR(128) /* 存储文件路径*/
)