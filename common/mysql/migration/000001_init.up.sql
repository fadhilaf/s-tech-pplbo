CREATE TABLE IF NOT EXISTS users (
  id int(10) unsigned auto_increment PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  password_hash varchar(255) NOT NULL,
  address varchar(255),
  phone_number varchar(255),
  created_at timestamp NOT NULL DEFAULT (NOW())
);
