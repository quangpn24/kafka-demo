CREATE TABLE todo (
  id int AUTO_INCREMENT PRIMARY KEY,
  task varchar(256),
  status varchar(256),
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP
);
