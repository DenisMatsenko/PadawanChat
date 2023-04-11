-- +goose Up
-- +goose StatementBegin
CREATE TABLE "Authors" (
    "username" character varying(50) NOT NULL,
    "id" serial primary key
);
CREATE TABLE "Messages" (
    "id" serial primary key,
    "content" text,
    "authorId" integer NOT NULL
);
INSERT INTO "Authors" ("username") VALUES ('denis');
INSERT INTO "Authors" ("username") VALUES ('pavel');
INSERT INTO "Messages" ("content", "authorId") VALUES ('Hello, my brother name is Petr Adamov', 1);
INSERT INTO "Messages" ("content", "authorId") VALUES ('and i often confuse u and ur names...', 1);
INSERT INTO "Messages" ("content", "authorId") VALUES ('most of Czech names are hard for me ;(', 1);
INSERT INTO "Messages" ("content", "authorId") VALUES ('Sup, its ok', 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Authors";
DROP TABLE "Messages";
-- +goose StatementEnd
