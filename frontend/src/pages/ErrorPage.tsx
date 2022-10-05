import { FC } from "react";

interface ErrorProps {
  message: string
}
const ErrorPage: FC<ErrorProps> = (props) => {
  return (
    <header className="App-header">
      <div id="error-page" className="text-center">
        <h1>Oops!</h1>
        <p className="py-6">Sorry, an unexpected error has occurred.</p>
        <p className="text-slate-400">
          <i>{props.message}</i>
        </p>
      </div>
    </header>
  );
};

export default ErrorPage;
