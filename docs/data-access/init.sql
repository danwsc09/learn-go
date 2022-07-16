DROP TABLE IF EXISTS album;
/*
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);
*/
CREATE TABLE album (
  id         SERIAL PRIMARY KEY,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  region    VARCHAR(255)
);

INSERT INTO album
  (title, artist, price, region)
VALUES
  ('Blue Train', 'John Coltrane', 56.99, NULL),
  ('Giant Steps', 'John Coltrane', 63.99, NULL),
  ('Jeru', 'Gerry Mulligan', 17.99, NULL),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98, 'California');