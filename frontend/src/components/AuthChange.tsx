import { Component, Show, createSignal, createEffect } from "solid-js";
import { getAuthSettings, changeAuth } from "../Utils";
import { useGlobalContext } from "./GlobalContext";

const AuthChange: Component = () => {
  const { token, setToken } = useGlobalContext();

  const [login, setLogin] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [passwordRepeated, setPasswordRepeated] = createSignal("");
  const [isError, setIsError] = createSignal(false);
  const [errorMessage, setErrorMessage] = createSignal("");

  // get auth settings
  createEffect(async () => {
    const response = await getAuthSettings(token());
    setLogin(response.login);
  });

  const submit = (event?: Event) => {
    if (event) {
      event.preventDefault();
    }

    if (login() === "" || password() === "" || passwordRepeated() === "") {
      setIsError(true);
      setErrorMessage("Invalid login or password");
      return;
    }

    changeAuth(login(), password(), token()).then((authResult) => {
      if (authResult.isError) {
        setIsError(true);
        setErrorMessage(authResult.ErrorMessage);
      } else {
        clearForm();
        setToken(authResult.Token);
      }
    });
  };

  const onLoginChange = (event: InputEvent) => {
    setLogin((event.target as HTMLInputElement).value);
  };

  const onPasswordChange = (event: InputEvent) => {
    setPassword((event.target as HTMLInputElement).value);
  };

  const onPasswordRepeatedChange = (event: InputEvent) => {
    setPasswordRepeated((event.target as HTMLInputElement).value);
  };

  const clearForm = () => {
    setIsError(false);
    setErrorMessage("");
    setLogin("");
    setPassword("");
    setPasswordRepeated("");
  };

  return (
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
      <div class="mb-3">
        <input
          type="password"
          class="form-control"
          id="password"
          placeholder="Repeat password"
          value={passwordRepeated()}
          oninput={onPasswordRepeatedChange}
        />
      </div>
      <Show when={isError()}>
        <div class="d-flex justify-content-center mb-3 text-danger">
          {errorMessage()}
        </div>
      </Show>
      <button class="btn btn-outline-secondary me-3" onclick={submit}>
        Save
      </button>
    </form>
  );
};

export default AuthChange;
