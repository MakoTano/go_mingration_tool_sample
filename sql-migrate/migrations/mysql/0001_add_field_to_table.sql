-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE book
(
id INT(11) NOT NULL AUTO_INCREMENT,
category_id INT(11),
title VARCHAR(64),
PRIMARY KEY (id)
);

INSERT INTO book (category_id,title)
VALUES (10,"よく分かるサル講座"),
(10,"歴史に学ぶ！サルの捉え方100選"),
(10,"おびき寄せて網をかける方法");

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE book;