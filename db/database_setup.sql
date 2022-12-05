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
INSERT INTO blog VALUES (2,
						 'First blog test',
						 1,
						 'blog testblog testblog testblog testblog test',
						 '<p>blog test</p><p>blog test</p><p>blog test</p><p>blog test</p>',
						 108,
						 4,
						 '2022-12-03 10:19:30','2022-12-03 10:19:30'),
						 (3,
						 'Second blog test',
						 1,
						 'blog testblog testblog testblog testblog test',
						 '<p>blog test</p><p>blog test</p><p>blog test</p><p>blog test</p>',
						 108,
						 4,
						 '2022-12-04 10:19:30','2022-12-04 10:19:30');


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
  ip varchar(50) DEFAULT '',
  content varchar(2000) DEFAULT '',
  blog_id int DEFAULT NULL,
  status smallint DEFAULT '0',
  add_time varchar(19) DEFAULT ''
);


-- insert test comment 
INSERT INTO comment VALUES (15,
							'127.0.0.1',
							'test comment test comment',
							2,
							0,
							'2022-12-05 15:12:35'),
							(16,
							 '127.0.0.1',
							 '兄弟你挺牛逼的嘛',
							 2,
							 0,
							 '2022-12-05 20:06:25');