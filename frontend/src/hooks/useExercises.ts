import { useEffect, useState } from "react";

const endpoint = `${process.env.REACT_APP_API_BASE}/exercise`

const useExercises = () => {
  const [exercises, setExercises] = useState([]);

  useEffect(() => {
    const request = fetch(endpoint);

    request
      .then((apiResponse) => {
        if (!apiResponse.ok) {
          console.error(apiResponse.statusText);
          return;
        }
        return apiResponse.json();
      })
      .then((apiResult) => {
        setExercises(apiResult);
      })
  }, []);

  return exercises;
}

export default useExercises;