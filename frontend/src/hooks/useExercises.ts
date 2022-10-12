import { useEffect, useState } from "react";
import { IExercise } from "../types/IExercise";
import { IAPIResult } from "../types/IAPIResult";

const endpoint = `${process.env.REACT_APP_API_BASE}/api/v1/exercise`

const useExercises = (): IAPIResult<IExercise[]> => {
  const [exercises, setExercises] = useState<IExercise[] | undefined>([]);
  const [error, setError] = useState<Error | undefined>();

  useEffect(() => {
    const request = fetch(endpoint);
    request
      .then((res) => {
        if (!res.ok) {
          setError(new Error(`${res.status}: ${res.statusText}`));
          return;
        }
        return res.json();
      })
      .then((res) => {
        setExercises(res);
      })
  }, []);
  
  if (error !== undefined) {
    console.error(error)
  }
  return { result: exercises, error: error };
}

export default useExercises;