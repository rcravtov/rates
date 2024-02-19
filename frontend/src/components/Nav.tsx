import { Component, Switch, Match, Show, createEffect } from "solid-js";
import { A, useNavigate } from "@solidjs/router";
import { useGlobalContext } from "./GlobalContext";
import Auth from "./Auth";

const Nav: Component = () => {
  const navigate = useNavigate();

  const {
    authVisible,
    setAuthVisible,
    isLoggedIn,
    setIsLoggedIn,
    token,
    setToken,
  } = useGlobalContext();

  const loginClick = () => {
    setAuthVisible(true);
  };

  const logoutClick = () => {
    setIsLoggedIn(false);
    setToken("");
    navigate("/");
  };

  return (
    <>
      <nav class="navbar navbar-expand-lg bg-body-tertiary">
        <div class="container-fluid">
          <div class="navbar-brand">
            <A href="/" class="nav-link active">
              Rates
            </A>
          </div>

          <ul class="navbar-nav me-auto">
            <li class="nav-item">
              <Show when={isLoggedIn()}>
                <A href="/settings" class="nav-link active">
                  Settings
                </A>
              </Show>
            </li>
          </ul>

          <div class="d-flex">
            <Switch>
              <Match when={isLoggedIn()}>
                <button
                  class="btn btn-outline-success"
                  type="submit"
                  onclick={logoutClick}
                >
                  Logout
                </button>
              </Match>

              <Match when={!isLoggedIn()}>
                <button
                  class="btn btn-outline-secondary"
                  type="submit"
                  onclick={loginClick}
                >
                  Login
                </button>
              </Match>
            </Switch>
          </div>
        </div>
      </nav>
      <Auth />
    </>
  );
};

export default Nav;
