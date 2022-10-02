import { useEffect, useState } from "react";

const endpoint = `${process.env.REACT_APP_API_BASE}/exercise`

const useExercise = (id: number) => {
  const [exercise, setExercise] = useState();

  useEffect(() => {
    const request = fetch(`${endpoint}/${id}`);
    request
      .then((apiResponse) => {
        if (!apiResponse.ok) {
          console.error(apiResponse.statusText)
          return;
        }
        return apiResponse.json();
      })
      .then((apiResult) => {
        setExercise(apiResult);
      })
  }, [id]);

  return exercise;
}

export default useExercise;