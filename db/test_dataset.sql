INSERT INTO blogger VALUES (1,'user','abc123','WeiTim','哈摟','test test 123 abc','/static/upload/20221204/wei.png');


INSERT INTO blog_type (
	name,
	sort
) 
VALUES ('AWS', 1),('Golang', 2),('Data Structure', 3);

INSERT INTO blog (
	title,
	type_id,
	summary,
	content
) 
VALUES ('First blog test',
	2,
	'blog testblog testblog testblog testblog test',
	'## title1'),
	('Second blog test',
	1,
	'blog testblog testblog testblog testblog test',
	'## title1'),
	('測試部落格',
	3,
	'blog testblog testblog testblog testblog test',
	'## title1');



INSERT INTO comment (
	nick_name,
	ip,
	content,
	blog_id
)
VALUES (
		'testMan',
		'127.0.0.1',
		'test comment test comment',
		3
		),
		(
		'testMan2',
		'127.0.0.1',
		'test comment2 test comment2',
		1
		),
		(
		'testMan3',
		'127.0.0.1',
		'test comment3 test comment3',
		2
		);

INSERT INTO tag VALUES	(1, 'blog', '1'),
						(2, 'Algorithm', '2'),
						(3, 'Data Structure', '3');

INSERT INTO blog_tag VALUES	(1, 1, 1),
							(2, 1, 2),
							(3, 1, 3),
							(4, 2, 1),
							(5, 3, 3);