import { FC } from "react";
import { IExercise } from "../types/IExercise";

interface ExerciseProps {
  exercise: IExercise;
}

const ExerciseCard : FC<ExerciseProps> = ({exercise}) => {
  return (
    <>
      <p>{exercise.id}</p>
      <p>{exercise.name}</p>
    </>
  );
};

export default ExerciseCard;
