
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO items (name, description, source, url_image) VALUES (
  'The Persistence of Memory',
  'By Salvador Dali (1931)',
  'custom',
  'http://uploads5.wikiart.org/images/salvador-dali/the-persistence-of-memory-1931.jpg'
);

INSERT INTO items (name, description, source, url_image) VALUES (
  'Ballerina in a Death''s Head',
  'By Salvador Dali (1939)',
  'custom',
  'http://uploads2.wikiart.org/images/salvador-dali/ballerina-in-a-death-s-head.jpg'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
