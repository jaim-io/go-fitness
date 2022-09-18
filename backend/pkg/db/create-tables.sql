CREATE TABLE exercises (
  id    INT AUTO_INCREMENT NOT NULL,
  name  VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO exercises
  (name)
VALUES
  ("Dumbell bench press");