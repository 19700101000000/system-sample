CREATE TABLE `user` (
    `id`        INTEGER     AUTO_INCREMENT,
    `name`      VARCHAR(32) UNIQUE NOT NULL,
    `show_name` VARCHAR(32) NOT NULL DEFAULT '',
    `password`  CHAR(64)    NOT NULL,
    `alive`     BOOLEAN     NOT NULL DEFAULT TRUE,
    `creator`   BOOLEAN     NOT NULL DEFAULT FALSE,

    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE DEFINER='server'@'%' TRIGGER `user`
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

