import { useEffect, useState } from "react";
import { IExercise } from "../types/IExercise";
import { IAPIResult } from "../types/IAPIResult";

const endpoint = `${process.env.REACT_APP_API_BASE}/api/v1/exercise`;

const useExercise = (id: string): IAPIResult<IExercise> => {
  const [exercise, setExercise] = useState<IExercise | undefined>();
  const [error, setError] = useState<Error | undefined>();

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
  }, [id]);

  if (error !== undefined) {
    console.error(error)
  }
  return { result: exercise, error: error };
};

export default useExercise;
