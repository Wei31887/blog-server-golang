CREATE TABLE blogger (
    id integer not null primary key,
    username varchar(40) not null,
    password varchar(100) not null,
    nickname varchar(20) not null,
    sign    varchar(20) not null,
    profile varchar(200) not null,
    img varchar(100) not null
);



DROP TABLE IF EXISTS blogger;
CREATE TABLE blogger (
  id integer not null primary key,
  username varchar(200) DEFAULT '',
  password varchar(200) DEFAULT '',
  nickname varchar(200) DEFAULT NULL,
  sign varchar(500) DEFAULT null,
  profile text,
  img varchar(500) DEFAULT NULL,
);

INSERT INTO blogger VALUES (1,'wei','tim870318','WeiTim','部落格初探','<p><img src=\"/api/static/upload/20221204/0511286.jpg\" style=\"max-width:100%;\"></p>','/static/upload/20221204/0511286.jpg');


DROP TABLE IF EXISTS blog;
CREATE TABLE blog (
  id serial primary key NOT NULL ,
  title varchar(200) DEFAULT '',
  typeId int DEFAULT '0',
  summary text,
  content text,
  click_hit int DEFAULT '0',
  replay_hit int DEFAULT '0',
  add_time varchar(19) DEFAULT '',
  update_time varchar(19) DEFAULT ''
);

-- insert test blog
INSERT INTO blog VALUES (1,
						 'First blog test',
						 1,
						 'blog testblog testblog testblog testblog test',
						 '<p>blog test</p><p>blog test</p><p>blog test</p><p>blog test</p>',
						 0,
						 4,
						 '2022-12-02 10:19:30','2022-12-02 10:19:30'),
						 
						 (2,
						 'Second blog test',
						 2,
						 'Nice one',
						 '<p>blog test2</p><p>blog test2</p><p>blog test2</p><p>blog test</p>',
						 10,
						 4,
						 '2022-12-04 22:19:30','2022-12-04 22:19:30'),
						 
						 (3,
						 'Third blog test',
						 2,
						 'Nice one',
						 '<p>blog test3</p><p>blog test3</p><p>blog test</p><p>blog test</p>',
						 10,
						 4,
						 '2022-12-04 22:29:30','2022-12-04 22:29:30'),
						 
						 (4,
						 'Continue blog test',
						 1,
						 'blog test blog test2 blog test2 blog test2 blog test',
						 '<p>blog test</p><p>blog test</p><p>blog test</p><p>blog test</p>',
						 10,
						 4,
						 '2022-12-04 23:19:30','2022-12-04 23:19:30');


DROP TABLE IF EXISTS blog_type;
CREATE TABLE blog_type (
  id serial primary key NOT NULL,
  name varchar(200) DEFAULT '',
  sort int DEFAULT '0'
);

-- insert test blog type
INSERT INTO blog_type VALUES (1,'AWS', 1),(2,'Golang', 6);


DROP TABLE IF EXISTS comment;
CREATE TABLE comment (
  id serial primary key NOT NULL,
  nick_name varchar(50) NOT NULL DEFAULT '',
  ip varchar(50) DEFAULT '',
  content varchar(2000) DEFAULT '',
  blog_id int DEFAULT NULL,
  status smallint DEFAULT '0',
  add_time varchar(19) DEFAULT ''
);


-- insert test comment 
INSERT INTO comment VALUES (15,
							'testMan',
							'127.0.0.1',
							'test comment test comment',
							3,
							0,
							'2022-12-05 15:12:35'),
							(16,
							'testMan2',
							 '127.0.0.1',
							 '兄弟你挺牛逼的嘛',
							 3,
							 0,
							 '2022-12-05 20:06:25');


DROP TABLE IF EXISTS tag;
CREATE TABLE tag (
  id serial primary key NOT NULL,
  tag_name varchar(50) NOT NULL DEFAULT '',
  sort int DEFAULT '0'
);

INSERT INTO tag VALUES	(1, 'blog', '1'),
						(2, 'Algorithm', '2'),
						(3, 'Data Structure', '3');

DROP TABLE IF EXISTS blog_tag;
CREATE TABLE blog_tag (
  id serial primary key NOT NULL,
  blog_id int DEFAULT NULL NOT NULL,
  tag_id int DEFAULT NULL NOT NULL
);

INSERT INTO blog_tag VALUES	(1, 3, 1),
							(2, 3, 2),
							(3, 3, 3),
							(4, 4, 1),
							(5, 5, 3);