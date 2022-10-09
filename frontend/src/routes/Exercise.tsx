import { useParams } from "react-router-dom";
import ExerciseCard from "../components/ExerciseCard";
import useExercise from "../hooks/useExercise";
import ErrorPage from "../pages/ErrorPage";

const Exercise = () => {
  let { id } = useParams();
  const { result, error } = useExercise(id ?? "");
  if (error !== undefined) return <ErrorPage message={error.message} />;
  return (
    <div id="exercise">{result && <ExerciseCard exercise={result} />}</div>
  );
};

export default Exercise;
