
drop database ikmandb;
create database if not exists ikmandb;

use ikmandb;

CREATE TABLE IF NOT EXISTS item (
  itemid INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(100) NULL,
  descr VARCHAR(450) NULL,
  price VARCHAR(450) NULL,
  location VARCHAR(30) NULL,
  category VARCHAR(30) NULL,
  
  CONSTRAINT PRIMARY KEY (itemid)
);