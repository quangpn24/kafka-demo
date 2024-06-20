CREATE TABLE users (
   id int AUTO_INCREMENT PRIMARY KEY,
   first_name varchar(256),
   last_name varchar(256),
   gender varchar(256),
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
   id int AUTO_INCREMENT PRIMARY KEY,
   user_id int,
   title varchar(256),
   status varchar(256),
   description varchar(256),
   created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);