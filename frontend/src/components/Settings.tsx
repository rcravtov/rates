import { Component, createEffect } from "solid-js";
import { useNavigate } from "@solidjs/router";
import { useGlobalContext } from "./GlobalContext";
import AuthChange from "./AuthChange";
import ImportSettings from "./ImportSettings";
import ImportLogs from "./ImportLogs";

const Settings: Component = () => {
  const { isLoggedIn } = useGlobalContext();
  const navigate = useNavigate();

  createEffect(() => {
    if (!isLoggedIn()) {
      navigate("/");
    }
  });

  return (
    <div>
      <div class="container">
        <div class="row">
          <div class="col d-flex">
            <div class="card w-100 mb-2">
              <div class="card-header">Administrator account</div>
              <div class="card-body d-flex flex-column flex-grow-1">
                <AuthChange />
              </div>
            </div>
          </div>

          <div class="col d-flex">
            <div class="card w-100 mb-2">
              <div class="card-header">Import settings</div>
              <div class="card-body d-flex flex-column flex-grow-1">
                <ImportSettings />
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="row p-2">
        <div class="col d-flex">
          <div class="card w-100 mb-2">
            <div class="card-header">Import logs</div>
            <div class="card-body d-flex flex-column flex-grow-1">
              <ImportLogs />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Settings;
