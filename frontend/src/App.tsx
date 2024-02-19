import { Component, createSignal } from "solid-js";
import Nav from "./components/Nav";
import { Route, Routes } from "@solidjs/router";
import Rates from "./components/Rates";
import Settings from "./components/Settings";

const App: Component = () => {
  return (
    <div class="container">
      <Nav />
      <Routes>
        <Route path="/" component={Rates} />
        <Route path="/settings" component={Settings} />
      </Routes>
    </div>
  );
};

export default App;
