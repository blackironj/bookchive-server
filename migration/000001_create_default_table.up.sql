-- users Table Create SQL
CREATE TABLE bookchive.users
(
    `uk`         INT            NOT NULL    AUTO_INCREMENT, 
    `email`      VARCHAR(45)    NOT NULL, 
    `name`       VARCHAR(15)    NULL, 
    `signin_dt`  INT            NULL, 
    PRIMARY KEY (uk)
);


-- users Table Create SQL
CREATE TABLE bookchive.books
(
    `uk`          INT                   NOT NULL    AUTO_INCREMENT, 
    `id`          VARCHAR(20) BINARY    NOT NULL, 
    `title`       VARCHAR(45)           NOT NULL, 
    `subtitle`    VARCHAR(100)          NULL, 
    `authors`     VARCHAR(50)           NOT NULL, 
    `publisher`   VARCHAR(25)           NOT NULL, 
    `categories`  VARCHAR(45)           NULL, 
    `thumbnail`   VARCHAR(250)          NULL, 
    `pages`       INT                   NULL, 
    PRIMARY KEY (uk)
);


-- users Table Create SQL
CREATE TABLE bookchive.libraries
(
    `uk`        INT    NOT NULL    AUTO_INCREMENT, 
    `user_uk`   INT    NOT NULL, 
    `book_uk`   INT    NOT NULL, 
    `added_dt`  INT    NULL, 
    PRIMARY KEY (uk)
);

ALTER TABLE bookchive.libraries
    ADD CONSTRAINT FK_libraries_user_uk_users_uk FOREIGN KEY (user_uk)
        REFERENCES bookchive.users (uk) ON DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE bookchive.libraries
    ADD CONSTRAINT FK_libraries_book_uk_books_uk FOREIGN KEY (book_uk)
        REFERENCES bookchive.books (uk) ON DELETE RESTRICT ON UPDATE RESTRICT;


-- users Table Create SQL
CREATE TABLE bookchive.goals
(
    `uk`            INT    NOT NULL    AUTO_INCREMENT, 
    `libraries_uk`  INT    NOT NULL, 
    `start_dt`      INT    NULL, 
    `end_dt`        INT    NULL, 
    PRIMARY KEY (uk)
);

ALTER TABLE bookchive.goals
    ADD CONSTRAINT FK_goals_libraries_uk_libraries_uk FOREIGN KEY (libraries_uk)
        REFERENCES bookchive.libraries (uk) ON DELETE RESTRICT ON UPDATE RESTRICT;


-- users Table Create SQL
CREATE TABLE bookchive.diaries
(
    `uk`            INT             NOT NULL    AUTO_INCREMENT, 
    `libraries_uk`  INT             NOT NULL, 
    `title`         VARCHAR(150)    NOT NULL, 
    `contents`      MEDIUMTEXT      NULL, 
    `pos_page`      INT             NULL, 
    `added_dt`      INT             NOT NULL, 
    `updated_dt`    INT             NOT NULL, 
    PRIMARY KEY (uk)
);

ALTER TABLE bookchive.diaries
    ADD CONSTRAINT FK_diaries_libraries_uk_libraries_uk FOREIGN KEY (libraries_uk)
        REFERENCES bookchive.libraries (uk) ON DELETE RESTRICT ON UPDATE RESTRICT;


