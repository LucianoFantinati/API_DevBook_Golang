CREATE DATABASE IF NOT EXISTS devebook;

USE devebook;

DROP TABLE IF EXISTS devebook;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) NOT null,
    nick varchar(50) NOT null unique,
    email varchar(50) NOT null unique,
    senha varchar(100) NOT null ,
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;