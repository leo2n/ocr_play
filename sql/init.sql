USE db_play;

DROP TABLE ocr_history;

CREATE TABLE IF NOT EXISTS ocr_history (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL ,
    user VARCHAR(128) DEFAULT '-1' NOT NULL , /* 用户名默认设置为-1吧*/
    result TEXT , /*16383是utf8mb4类型的字符可以存储的最大值, 因为65535bytes/4 ~ 16383 */
    time timestamp DEFAULT CURRENT_TIMESTAMP
)
