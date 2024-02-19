import {
  Component,
  Show,
  createEffect,
  createSignal,
  onCleanup,
} from "solid-js";

import { Portal } from "solid-js/web";
import { authorize } from "../Utils";
import { useGlobalContext } from "./GlobalContext";

const Auth: Component = () => {
  const {
    authVisible,
    setAuthVisible,
    isLoggedIn,
    setIsLoggedIn,
    token,
    setToken,
  } = useGlobalContext();

  const [login, setLogin] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [isError, setIsError] = createSignal(false);
  const [errorMessage, setErrorMessage] = createSignal("");

  createEffect(() => {
    function handleKeyDown(event: KeyboardEvent) {
      if (event.key === "Escape") {
        setAuthVisible(false);
      } else if (event.key === "Enter") {
        submit();
      }
    }

    if (authVisible()) {
      window.addEventListener("keydown", handleKeyDown);
    } else {
      window.removeEventListener("keydown", handleKeyDown);
    }

    onCleanup(() => {
      window.removeEventListener("keydown", handleKeyDown);
    });
  });

  const clearForm = () => {
    setIsError(false);
    setErrorMessage("");
    setLogin("");
    setPassword("");
  };

  const submit = (event?: Event) => {
    if (event) {
      event.preventDefault();
    }
    authorize(login(), password()).then((authResult) => {
      if (authResult.isError) {
        setIsError(true);
        setErrorMessage(authResult.ErrorMessage);
        setIsLoggedIn(false);
        setToken("");
      } else {
        clearForm();
        setAuthVisible(false);
        setIsLoggedIn(true);
        setToken(authResult.Token);
      }
    });
  };

  const cancel = () => {
    clearForm();
    setAuthVisible(false);
  };

  const onLoginChange = (event: InputEvent) => {
    setLogin((event.target as HTMLInputElement).value);
  };

  const onPasswordChange = (event: InputEvent) => {
    setPassword((event.target as HTMLInputElement).value);
  };

  return (
    <Show when={authVisible()}>
      <Portal mount={document.body}>
        <div class="modal show" style="display: block;">
          <div class="modal-dialog">
            <div class="container modal-content p-2">
              <form>
                <div class="mb-3">
                  <input
                    type="text"
                    class="form-control"
                    id="login"
                    placeholder="Login"
                    value={login()}
                    oninput={onLoginChange}
                  />
                </div>
                <div class="mb-3">
                  <input
                    type="password"
                    class="form-control"
                    id="password"
                    placeholder="Password"
                    value={password()}
                    oninput={onPasswordChange}
                  />
                </div>
                <Show when={isError()}>
                  <div class="d-flex justify-content-center mb-3 text-danger">
                    {errorMessage()}
                  </div>
                </Show>
                <div class="d-flex justify-content-center">
                  <button
                    class="btn btn-secondary align-center me-3"
                    onclick={submit}
                  >
                    Submit
                  </button>
                  <button
                    class="btn btn-outline-secondary align-center"
                    onclick={cancel}
                  >
                    Cancel
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </Portal>
    </Show>
  );
};

export default Auth;
