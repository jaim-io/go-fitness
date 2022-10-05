import useExercises from "../hooks/useExercises";
import ErrorPage from "../pages/ErrorPage";

export const Exercises = () => {
  const { result, error } = useExercises();
  if (error !== undefined) return <ErrorPage message={error.message} />;
  return (
    <div id="exercises">
      {result &&
        result.map((exercise) => {
          if (exercise) {
            return <p>{exercise.name}</p>;
          }
          throw new Error("Internal error: exercise is of type null | undefined")
        })}
    </div>
  );
};

export default Exercises;
