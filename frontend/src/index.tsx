/* @refresh reload */
import { render } from "solid-js/web";
import { Router } from "@solidjs/router";
import "bootstrap/dist/css/bootstrap.min.css";
import { GlobalContextProvider } from "./components/GlobalContext";

import App from "./App";

const root = document.getElementById("root");

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    "Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got mispelled?"
  );
}

render(
  () => (
    <GlobalContextProvider>
      <Router>
        <App />
      </Router>
    </GlobalContextProvider>
  ),
  root!
);
