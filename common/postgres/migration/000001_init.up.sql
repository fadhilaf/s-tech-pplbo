CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  password_hash varchar(255) NOT NULL,
  address varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE IF NOT EXISTS products (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  stock integer NOT NULL,
  description varchar(255) NOT NULL,
  price numeric(10,2),
  is_service boolean NOT NULL DEFAULT false,
  created_at timestamp NOT NULL DEFAULT (NOW())
);
