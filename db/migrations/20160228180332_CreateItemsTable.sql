
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE items (
  id SERIAL PRIMARY KEY,
  created_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR (100) NOT NULL,
  description TEXT NOT NULL,
  source VARCHAR (100) NOT NULL,
  url_image VARCHAR (2083) NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
