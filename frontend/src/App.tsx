import type { Component } from 'solid-js';
import Nav from './components/Nav';
import { Route, Routes } from '@solidjs/router';
import Rates from './components/Rates';

const App: Component = () => {
  return (
    <div class="container">
      <Nav/>
      <Routes>
        <Route path="/" component={Rates} />
      </Routes>
    </div>
  );
};

export default App;
