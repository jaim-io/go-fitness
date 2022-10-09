CREATE TABLE exercises (
  id            SERIAL  NOT NULL PRIMARY KEY,
  name          VARCHAR(60)         NOT NULL,
  description   VARCHAR(512)        NOT NULL, 
  image_path    VARCHAR(1024)       NOT NULL,
  video_link    VARCHAR(1024)       NOT NULL
);

CREATE TABLE muscle_groups(
  id          SERIAL  NOT NULL PRIMARY KEY,
  name        VARCHAR(60)         NOT NULL,
  description   VARCHAR(512)      NOT NULL, 
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
  (name, description, image_path)
VALUES
  (
    'Chest', 
    'pectoralis muscle, any of the muscles that connect the front walls of the chest with the bones of the upper arm and shoulder. There are two such muscles on each side of the sternum (breastbone) in the human body: pectoralis major and pectoralis minor.',
    '/assets/images/muscle_groups/chest'
  ),(
    'Tricep', 
    'The triceps brachii is a large, thick muscle on the dorsal part of the upper arm. It often appears as the shape of a horseshoe on the posterior aspect of the arm. The main function of the triceps is the extension of the elbow joint.',
    '/assets/images/muscle_groups/tricep'
  ),(
    'Shoulder', 
    'The primary muscle group that supports the shoulder joint is the rotator cuff muscles. The four rotator cuff muscles are the supraspinatus, infraspinatus, teres minor, and subscapularis. Together the rotator cuff muscles form a musculotendinous cuff as they insert on the proximal humerus.',
    '/assets/images/muscle_groups/shoulder'
  ),(
    'Bicep', 
    'The biceps is a large muscle situated on the front of the upper arm between the shoulder and the elbow. Also known by the Latin name biceps brachii (meaning "two-headed muscle of the arm"), the muscle"s primary function is to flex the elbow and rotate the forearm.',
    '/assets/images/muscle_groups/bicep'
  );


INSERT INTO exercise_muscle_groups VALUES (1,1), (1,2), (1,3), (2,4);

