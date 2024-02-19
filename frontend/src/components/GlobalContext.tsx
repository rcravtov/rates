import {
  Accessor,
  Setter,
  createContext,
  useContext,
  createSignal,
} from "solid-js";

interface ContextProps {
  authVisible: Accessor<boolean>;
  setAuthVisible: Setter<boolean>;
  isLoggedIn: Accessor<boolean>;
  setIsLoggedIn: Setter<boolean>;
  token: Accessor<string>;
  setToken: Setter<string>;
}

const GlobalContext = createContext<ContextProps>();

export function GlobalContextProvider(props: any) {
  const [authVisible, setAuthVisible] = createSignal(false);
  const [isLoggedIn, setIsLoggedIn] = createSignal(false);
  const [token, setToken] = createSignal("");

  return (
    <GlobalContext.Provider
      value={{
        authVisible,
        setAuthVisible,
        isLoggedIn,
        setIsLoggedIn,
        token,
        setToken,
      }}
    >
      {props.children}
    </GlobalContext.Provider>
  );
}

export const useGlobalContext = () => useContext(GlobalContext)!;
