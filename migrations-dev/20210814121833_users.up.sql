CREATE TABLE `Users`
(
    `id`               INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `login`            VARCHAR(255) NOT NULL,
    `password`         VARCHAR(255) NOT NULL,
    `email`            VARCHAR(255) NOT NULL,
    `name`             VARCHAR(255) NOT NULL,
    `lastName`         VARCHAR(255),
    `type`             TINYINT UNSIGNED NOT NULL DEFAULT (0),
    `dateTimeCreation` DATETIME     NOT NULL DEFAULT (NOW()),
    `dateOfBirth`      DATE,
    `isDeleted`        BOOLEAN      NOT NULL DEFAULT (false),
    `isBanned`         BOOLEAN      NOT NULL DEFAULT (false),
    `isConfirmed`      BOOLEAN      NOT NULL DEFAULT (false),
    PRIMARY KEY (id)
);