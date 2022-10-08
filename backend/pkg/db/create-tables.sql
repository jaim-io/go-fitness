CREATE TABLE exercises (
  id            SERIAL  NOT NULL PRIMARY KEY,
  name          VARCHAR(60)         NOT NULL,
  description   VARCHAR(255)        NOT NULL, 
  image_path    VARCHAR(1024)       NOT NULL,
  video_link    VARCHAR(1024)       NOT NULL
);

CREATE TABLE muscle_groups(
  id          SERIAL  NOT NULL PRIMARY KEY,
  name        VARCHAR(60)         NOT NULL,
  image_path  VARCHAR(1024)       NOT NULL
);


CREATE TABLE exercise_muscle_groups(
  exercise_id INT NOT NULL,
  muscle_group_id INT NOT NULL,
  FOREIGN KEY (exercise_id)     REFERENCES exercises(id),
  FOREIGN KEY (muscle_group_id) REFERENCES muscle_groups(id));

INSERT INTO exercises
  (name, description, image_path, video_link)
VALUES
  (
    'Barbell bench press', 
    'The barbell bench press exercise is an upper body pressing movement. Commonly used to build muscle size and body strength; this exercise targets the upper body muscles including: chest, triceps, and shoulders.',
    '/assets/images/exercises/barbell_bench_press',
    'https://www.youtube.com/watch?v=vcBig73ojpE'
  ),(
    'Preacher curl', 
    'The preacher curl is a variation of the traditional biceps curl. It is an isolation bicep exercise that allows you to practice your lifting form with a controlled movement supported by a preacher bench.',
    '/assets/images/exercises/preacher_curl',
    'https://www.youtube.com/watch?v=fIWP-FRFNU0'
  )
  ;

INSERT INTO muscle_groups
  (name, image_path)
VALUES
  ('Chest', '/assets/images/muscle_groups/chest'),
  ('Tricep', '/assets/images/muscle_groups/tricep'),
  ('Shoulder', '/assets/images/muscle_groups/shoulder'),
  ('Bicep', '/assets/images/muscle_groups/bicep');


INSERT INTO exercise_muscle_groups VALUES (1,1), (1,2), (1,3), (2,4);

