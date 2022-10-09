ADD muscle group tests
ADD exercise tests

CHANGE image_path -> image [FRONTEND/IExercise-IMuscleGroup]
CHANGE image upload muscle group API
CHANGE image upload exercise API

Frontend error handling
Check if this:
  ```
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
  ```
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

PUT-POST exercise should check if musclegroup exists
  if TRUE
    - update the tables(exercises, exercise_muscle_groups)
  else if FALSE
    - reject request

DELETE exercise should remove relation from exercise_muscle_groups tables
  - update exercise_muscle_groups table

DELETE musclegroup should relation from exercise_muscle_groups tables
  - update exercise_muscle_groups table