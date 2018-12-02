CREATE TABLE `user` (
    `id`        INTEGER     AUTO_INCREMENT,
    `name`      VARCHAR(32) UNIQUE NOT NULL,
    `show_name` VARCHAR(32) NOT NULL DEFAULT '',
    `password`  CHAR(64)    NOT NULL,
    `alive`     BOOLEAN     NOT NULL DEFAULT TRUE,
    `creator`   BOOLEAN     NOT NULL DEFAULT FALSE,

    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE DEFINER='server'@'%' TRIGGER `insert_user`
    BEFORE INSERT ON `user` FOR EACH ROW
        SET NEW.`password` = SHA2(NEW.`password`, 0);

CREATE TABLE `monitor` (
    `observer`  INTEGER,
    `target`    INTEGER,
    `rate`      INTEGER         NOT NULL DEFAULT 0,
    `review`    VARCHAR(1024)   NOT NULL DEFAULT '',

    PRIMARY KEY(`observer`, `target`),
    FOREIGN KEY(`observer`) REFERENCES `user`(`id`),
    FOREIGN KEY(`target`)   REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `gallery` (
    `user`      INTEGER,
    `id`        INTEGER,
    `image`     VARCHAR(128)    NOT NULL,
    `favorite`  INTEGER         NOT NULL DEFAULT 0,
    PRIMARY KEY(`user`, `id`),

    FOREIGN KEY(`user`) REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DELIMITER $$
CREATE DEFINER='server'@'%' TRIGGER `insert_gallery`
    BEFORE INSERT ON `gallery` FOR EACH ROW
        BEGIN
            DECLARE `id` TYPE OF `gallery`.`id`;
            SELECT COUNT(*) FROM `gallery` WHERE `user` = NEW.`user` INTO `id`;
            SET NEW.`id` = `id` + 1;
        END$$
DELIMITER ;

