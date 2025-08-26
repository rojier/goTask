CREATE table learn_student(
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(50) NOT NULL,
 age  TINYINT NOT NULL,
 grade VARCHAR(50) NOT NULL
 
);

CREATE table account(
 id INT PRIMARY KEY AUTO_INCREMENT,
 balance FLOAT NULL
);
CREATE table transaction(
 id INT PRIMARY KEY AUTO_INCREMENT,
 from_balance_id INT,
 to_balance_id INT,
 amount FLOAT NULL
);

CREATE table employee(
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(50) NOT NULL,
 department  VARCHAR(50) NOT NULL,
 salary int NULL
);
CREATE table user(
 id INT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(50) NOT NULL
);
alter table user add COLUMN post_num INT;

INSERT into  user(name) VALUES ('张三');
INSERT into  user(name) VALUES ('李四');


CREATE table post(
 id INT PRIMARY KEY AUTO_INCREMENT,
 title VARCHAR(50) NOT NULL,
 user_id INT  NOT NULL
);
INSERT into  post(title,user_id) VALUES ('张三文章一',1); 
INSERT into  post(title,user_id) VALUES ('张三文章二',1); 
INSERT into  post(title,user_id) VALUES ('李四文章一',2); 
INSERT into  post(title,user_id) VALUES ('李四文章二',2); 

CREATE table comment(
 id INT PRIMARY KEY AUTO_INCREMENT,
 content VARCHAR(50) NOT NULL,
 post_id INT  NOT NULL
);
INSERT into comment(content,post_id) values ('张三文章1的评论1',1);
INSERT into comment(content,post_id) values ('张三文章1的评论2',1);
 
INSERT into comment(content,post_id) values ('李四文章1的评论1',3);
INSERT into comment(content,post_id) values ('李四文章1的评论2',3);
INSERT into comment(content,post_id) values ('李四文章2的评论1',4);


INSERT into employee(name,department,salary) values ('张三','技术部',15000);
INSERT into employee(name,department,salary) values ('李四','技术部',12000);
INSERT into employee(name,department,salary) values ('王五','销售',12000);

INSERT into account(balance) values (10000)


SELECT * from USER;

SELECT * from post;

SELECT * from comment;

select * from post where id in (
	select post_id from (
		SELECT  post_id,count(*) as count from comment GROUP BY post_id 
		HAVING count(*) = (select MAX(count) from (
		SELECT count(*) as count from comment  GROUP BY post_id ) as counts) 
		) as a 
)







SELECT * from learn_student;
SELECT * from account;
SELECT * from transaction;
SELECT * from employee;






