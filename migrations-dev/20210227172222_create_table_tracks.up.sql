CREATE TABLE `Tracks`
(
    `id`               INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `userId`           INT UNSIGNED                NOT NULL,
    `name`             VARCHAR(255)                NOT NULL,
    `description`      TEXT,
    `dateTimeCreation` DATETIME                    NOT NULL DEFAULT (NOW()),
    `duration`         INT                         NOT NULL,
    `bpm`              SMALLINT UNSIGNED,
    `isDeleted`        BOOLEAN,
    PRIMARY KEY (id)
);