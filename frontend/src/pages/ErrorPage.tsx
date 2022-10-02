import { useRouteError } from "react-router-dom";

export default function ErrorPage() {
  const error: any = useRouteError();
  console.error(error);

  return (
    <header className="App-header"> 
      <div id="error-page" className="text-center">
        <h1>Oops!</h1>
        <p className="py-6">Sorry, an unexpected error has occurred.</p>
        <p className="text-slate-400">
          <i>{error.statusText || error.message}</i>
        </p>
      </div>
    </header>
  );
}
