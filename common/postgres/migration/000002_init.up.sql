CREATE TYPE order_status AS ENUM ('pending', 'processing', 'delivered');

CREATE TABLE IF NOT EXISTS orders (
  id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  product_id uuid NOT NULL,
  quantity integer NOT NULL,
  status order_status NOT NULL,
  description varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (NOW()),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
