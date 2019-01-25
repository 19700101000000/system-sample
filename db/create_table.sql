/* Create user */
CREATE TABLE `user` (
    `id`        INTEGER,
    `name`      VARCHAR(32) UNIQUE NOT NULL,
    `show_name` VARCHAR(32) NOT NULL DEFAULT '',
    `password`  CHAR(64)    NOT NULL,
    `creator`   BOOLEAN     NOT NULL DEFAULT FALSE,
    `alive`     BOOLEAN     NOT NULL DEFAULT TRUE,

    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DELIMITER $$
CREATE DEFINER='server'@'%' TRIGGER `insert_user`
    BEFORE INSERT ON `user` FOR EACH ROW
        BEGIN
            DECLARE `id` TYPE OF `user`.`id`;
            SELECT COUNT(*) FROM `user` INTO `id`;
            SET NEW.`id`        = `id` + 1,
                NEW.`password`  = SHA2(NEW.`password`, 0);
        END$$
DELIMITER ;

/* Create monitor */
CREATE TABLE `monitor` (
    `observer`  INTEGER,
    `target`    INTEGER,
    `rate`      INTEGER         NOT NULL DEFAULT 0,
    `review`    VARCHAR(1024)   NOT NULL DEFAULT '',

    PRIMARY KEY(`observer`, `target`),
    FOREIGN KEY(`observer`) REFERENCES `user`(`id`),
    FOREIGN KEY(`target`)   REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* Create gallery */
CREATE TABLE `gallery` (
    `user`      INTEGER,
    `id`        INTEGER,
    `image`     VARCHAR(128)    NOT NULL,
    `timestamp` TIMESTAMP,
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

/* Create gallery_favorite */
CREATE TABLE `gallery_favorite` (
    `creator`   INTEGER,
    `gallery`   INTEGER,
    `favorer`   INTEGER,
    PRIMARY KEY(`creator`, `gallery`, `favorer`),

    FOREIGN KEY(`creator`, `gallery`)   REFERENCES `gallery`(`user`, `id`),
    FOREIGN KEY(`favorer`)              REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* Create gallery_comment */
CREATE TABLE `gallery_comment` (
    `creator`   INTEGER,
    `gallery`   INTEGER,
    `commenter` INTEGER,
    `comment`   VARCHAR(128) NOT NULL,
    PRIMARY KEY(`creator`, `gallery`, `commenter`),

    FOREIGN KEY(`creator`, `gallery`)   REFERENCES `gallery`(`user`, `id`),
    FOREIGN KEY(`commenter`)            REFERENCES `user`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* Create gallery_category */
CREATE TABLE `category` (
    `id`    INTEGER PRIMARY KEY AUTO_INCREMENT,
    `name`  VARCHAR(128) UNIQUE NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `gallery_join_category` (
    `creator`   INTEGER,
    `gallery`   INTEGER,
    `category`  INTEGER,
    PRIMARY KEY(`creator`, `gallery`, `category`),

    FOREIGN KEY(`creator`, `gallery`)   REFERENCES `gallery`(`user`, `id`),
    FOREIGN KEY(`category`)             REFERENCES `category`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
