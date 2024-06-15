CREATE TABLE users (
   id int AUTO_INCREMENT PRIMARY KEY,
   username varchar(256),
   age integer,
   gender varchar(256),
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
