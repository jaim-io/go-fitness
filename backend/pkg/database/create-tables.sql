DROP TABLE IF EXISTS exercise;
CREATE TABLE exercise (
  id    INT AUTO_INCREMENT NOT NULL,
  name  VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO exercise
  (name)
VALUES
  ("Dumbell bench press");