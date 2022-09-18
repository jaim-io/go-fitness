-- CREATE USER 'jamey'@'localhost' IDENTIFIED BY 'password';
-- FLUSH PRIVILEGES; 
-- GRANT ALL PRIVILEGES ON *.* TO 'jamey'@'localhost' WITH GRANT OPTION;
-- FLUSH PRIVILEGES; 

CREATE DATABASE jaimio;
USE jaimio;


CREATE TABLE exercises (
  id    INT AUTO_INCREMENT NOT NULL,
  name  VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO exercises
  (name)
VALUES
  ("Dumbell bench press");