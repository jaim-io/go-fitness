FEATURE Login
  DATABASE tables
  Backend 
    user model
    oauth API
    oauth-context mock
    oauth API tests
  Fronend login screen

CHANGE image_path -> image [FRONTEND/IExercise-IMuscleGroup] <br />
CHANGE image upload muscle group API <br />
CHANGE image upload exercise API <br />

Frontend error handling <br />
Check if this:
  ```tsx
  useEffect(() => {
    const request = fetch(`${endpoint}/${id}`);
    request
      .then((res) => {
        if (!res.ok) {
          setError(new Error(`${res.status}: ${res.statusText}`));
          return;
        }
        return res.json();
      })
      .then((res) => {
        setExercise(res);
      });
  ```
Should be changed to this:
  ```tsx
  useEffect(() => {
    const request = fetch(`${endpoint}/${id}`);
    request
      .then((res) => {
        if (!res.ok) {
          setError(new Error(`${res.status}: ${res.json()}`));
          return;
        }
        return res.json();
      })
  ```

DELETE && PUT exercise should remove relation from exercise_muscle_groups tables <br />
  - update exercise_muscle_groups table <br />

DELETE && PUT musclegroup should relation from exercise_muscle_groups tables <br />
  - update exercise_muscle_groups table <br />


When adding new exercise <br />
  option to choos from existing musclegroup  <br />
    or  <br />
  create new <br />