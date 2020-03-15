-- users Table Create SQL
CREATE TABLE bookchive.users
(
    `uuid`       VARCHAR(36)    NOT NULL, 
    `email`      VARCHAR(45)    NOT NULL, 
    `name`       VARCHAR(15)    NULL, 
    `signin_dt`  INT            NULL, 
    PRIMARY KEY (uuid)
);


-- users Table Create SQL
CREATE TABLE bookchive.books
(
    `id`          VARCHAR(12) BINARY    NOT NULL, 
    `title`       VARCHAR(45)           NOT NULL, 
    `subtitle`    VARCHAR(100)          NULL, 
    `authors`     VARCHAR(50)           NOT NULL, 
    `publisher`   VARCHAR(25)           NOT NULL, 
    `categories`  VARCHAR(45)           NULL, 
    `thumbnail`   VARCHAR(250)          NULL, 
    `pages`       INT                   NULL, 
    PRIMARY KEY (id)
);


-- users Table Create SQL
CREATE TABLE bookchive.libraries
(
    `uk`         INT                   NOT NULL    AUTO_INCREMENT, 
    `user_uuid`  VARCHAR(36)           NOT NULL, 
    `book_id`    VARCHAR(12) BINARY    NOT NULL, 
    `added_dt`   INT                   NULL, 
    PRIMARY KEY (uk)
);

ALTER TABLE bookchive.libraries
    ADD CONSTRAINT FK_libraries_user_uuid_users_uuid FOREIGN KEY (user_uuid)
        REFERENCES bookchive.users (uuid) ON DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE bookchive.libraries
    ADD CONSTRAINT FK_libraries_book_id_books_id FOREIGN KEY (book_id)
        REFERENCES bookchive.books (id) ON DELETE RESTRICT ON UPDATE RESTRICT;


-- users Table Create SQL
CREATE TABLE bookchive.goals
(
    `uk`            INT           NOT NULL    AUTO_INCREMENT, 
    `libraries_uk`  INT           NOT NULL, 
    `start_dt`      INT           NULL, 
    `end_dt`        INT           NULL, 
    `goal_dt`       INT           NULL, 
    `memo`          MEDIUMTEXT    NULL, 
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


