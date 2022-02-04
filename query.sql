SELECT * FROM customer;

describe customer;
describe comments;


-- create table user
-- (
--     id int AUTO_INCREMENT,
--     username varchar(255),
--     password varchar(255),
--     primary key (id)
-- ) ENGINE=InnoDB;

create table comments
(
    id int AUTO_INCREMENT PRIMARY KEY,
    email varchar(255),
    comment TEXT
   
) ENGINE=InnoDB;


insert into user(username,password) values('andi','admin123');

select * from user;