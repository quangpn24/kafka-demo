CREATE TABLE users (
   id int AUTO_INCREMENT PRIMARY KEY,
   username varchar(256),
   age integer,
   email varchar(256),
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
   id int AUTO_INCREMENT PRIMARY KEY,
   title varchar(256),
   status varchar(256),
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tags (
   id int AUTO_INCREMENT PRIMARY KEY,
   tag_name varchar(256),
   task_id int,
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
