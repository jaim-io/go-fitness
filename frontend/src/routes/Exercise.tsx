import { useParams } from "react-router-dom";
import ExerciseCard from "../components/ExerciseCard";
import useExercise from "../hooks/useExercise";
import ErrorPage from "../pages/ErrorPage";

const Exercise = () => {
  let { id } = useParams();
  const { result, error } = useExercise(id ?? "");
  console.log(error);
  
  if (error !== undefined) return <ErrorPage message={error.message}/>;
  return <>{result && <ExerciseCard exercise={result} />}</>;
};

export default Exercise;
