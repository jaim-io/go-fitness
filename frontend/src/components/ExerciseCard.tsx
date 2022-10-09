import { FC } from "react";
import { useNavigate } from "react-router-dom";
import { IExercise } from "../types/IExercise";

interface ExerciseProps {
  exercise: IExercise;
}

const ExerciseCard: FC<ExerciseProps> = ({ exercise }) => {
  const navigate = useNavigate();
  return (
    <div className="p-8 flex flex-col items-center">
      <figure
        className="md:flex bg-slate-100 rounded-xl p-8 md:p-0 dark:bg-slate-800 max-w-xl"
        onClick={() => navigate("/exercise/" + exercise.id)}
      >
        <img
          className="w-24 h-24 md:w-48 md:h-auto md:rounded-none rounded-full mx-auto"
          src="/leg-press.jpg"
          alt=""
          width="384"
          height="512"
        />
        <div className="pt-6 md:p-8 text-center md:text-left space-y-4">
          <blockquote>
            <figcaption className="text-xl font-medium">
              <div className="text-sky-500 dark:text-sky-400">
                {exercise.name}
              </div>
            </figcaption>
            <p className="text-base font-medium text-white">
              “{exercise.description}”
            </p>
          </blockquote>
        </div>
      </figure>
    </div>
  );
};

export default ExerciseCard;
