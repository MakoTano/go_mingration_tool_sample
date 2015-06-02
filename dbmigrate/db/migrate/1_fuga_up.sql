CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    category_id int,
    title char(50)
);

INSERT INTO book (category_id,title) VALUES (10,'三匹のサル');