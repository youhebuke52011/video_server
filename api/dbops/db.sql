create database video_server;

use video_server;

create table users(
  id int primary key auto_increment,
  login_name varchar(64) unique ,
  pwd varchar(64)
);

create table video_info(
  id varchar(64) primary key not null ,
  author_id int,
  name text,
  display_ctime TEXT,
  create_time datetime
);

create table comments(
  id varchar(64) primary key not null ,
  vedio_id varchar(64),
  author_id int,
  content text,
  time DATETIME
);

create table sessions(
  session_id varchar(64) primary key not null,
  TTL TINYTEXT,
  login_name varchar(64)
);



