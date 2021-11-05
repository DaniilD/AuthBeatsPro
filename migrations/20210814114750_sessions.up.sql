CREATE TABLE `Sessions`
(
    `refreshToken` VARCHAR(255) NOT NULL,
    `expiresAt` DATETIME NOT NULL,
    `userId` INT UNSIGNED NOT NULL
    PRIMARY KEY (refreshToken)
);