USE db_play;

DROP TABLE ocr_history;

CREATE TABLE IF NOT EXISTS ocr_history (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL , /* 必须和主键一起使用*/
    userId VARCHAR(128) , /* 用户名默认设置为-1吧*/
    result TEXT , /*16383是utf8mb4类型的字符可以存储的最大值, 因为65535bytes/4 ~ 16383 */
    time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
