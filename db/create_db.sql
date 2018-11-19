/* データベース削除 */
DROP DATABASE IF EXISTS `ih2018_db`;

/* データベース生成 */
CREATE DATABASE `ih2018_db` CHARACTER SET utf8;

/* 権限付与 */
GRANT ALL PRIVILEGES ON `ih2018_db`.* TO 'server'@'%';

