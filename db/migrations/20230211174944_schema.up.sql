CREATE TABLE "blogger" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "nickname" varchar NOT NULL,
  "sign" varchar NOT NULL,
  "profile" varchar NOT NULL,
  "img" varchar NOT NULL
);

CREATE TABLE "blog" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL DEFAULT '',
  "type_id" bigint NOT NULL DEFAULT 0,
  "summary" text,
  "content" text NOT NULL,
  "click_hit" int NOT NULL DEFAULT 0,
  "replay_hit" int NOT NULL DEFAULT 0,
  "add_time" timestamptz NOT NULL DEFAULT (now()),
  "update_time" timestamptz
);

CREATE TABLE "blog_type" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(200) NOT NULL DEFAULT '',
  "sort" int DEFAULT 0
);

CREATE TABLE "comment" (
  "id" bigserial PRIMARY KEY,
  "nick_name" varchar(50) NOT NULL DEFAULT '',
  "ip" varchar(50) NOT NULL DEFAULT '',
  "content" varchar(2000) NOT NULL DEFAULT '',
  "blog_id" bigint NOT NULL DEFAULT 0,
  "status" smallint NOT NULL DEFAULT 0,
  "add_time" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tag" (
  "id" bigserial PRIMARY KEY,
  "tag_name" varchar(50) NOT NULL DEFAULT '',
  "sort" int DEFAULT 0
);

CREATE TABLE "blog_tag" (
  "id" bigserial PRIMARY KEY,
  "blog_id" bigint NOT NULL,
  "tag_id" bigint NOT NULL
);

ALTER TABLE "blog" ADD FOREIGN KEY ("type_id") REFERENCES "blog_type" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "blog_tag" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "blog_tag" ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id");
