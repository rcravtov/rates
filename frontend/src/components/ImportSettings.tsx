import { Component, Show, createSignal, createEffect } from "solid-js";
import {
  getImportSettings,
  changeImportSettings,
  importNow,
  formatDate,
} from "../Utils";
import { useGlobalContext } from "./GlobalContext";

const ImportSettings: Component = () => {
  const { token } = useGlobalContext();

  const [autoImport, setAutoImport] = createSignal(false);
  const [importHours, setImportHours] = createSignal(0);
  const [importDate, setImportDate] = createSignal(new Date());
  const [importMinutes, setImportMinutes] = createSignal(0);
  const [isError, setIsError] = createSignal(false);
  const [errorMessage, setErrorMessage] = createSignal("");

  // get import settings
  createEffect(async () => {
    const response = await getImportSettings(token());
    setAutoImport(response.auto_import);
    setImportHours(response.import_hours);
    setImportMinutes(response.import_minutes);
  });

  const submit = (event?: Event) => {
    if (event) {
      event.preventDefault();
    }

    if (!isValidTime()) {
      setIsError(true);
      setErrorMessage("Invalid time");
      return;
    }

    changeImportSettings(
      autoImport(),
      importHours(),
      importMinutes(),
      token()
    ).then((genericResult) => {
      if (genericResult.isError) {
        setIsError(true);
        setErrorMessage(genericResult.ErrorMessage);
      } else {
        setIsError(false);
        setErrorMessage("");
      }
    });
  };

  const importNowClick = (event?: Event) => {
    if (event) {
      event.preventDefault();
    }
    importNow(importDate(), token()).then((genericResult) => {
      if (genericResult.isError) {
        setIsError(true);
        setErrorMessage(genericResult.ErrorMessage);
      } else {
        setIsError(false);
        setErrorMessage("");
      }
    });
  };

  const onAutoImportChange = (event: InputEvent) => {
    setAutoImport((event.target as HTMLInputElement).checked);
  };

  const onHoursChange = (event: InputEvent) => {
    setImportHours(Number((event.target as HTMLInputElement).value));
  };

  const onMinutesChange = (event: InputEvent) => {
    setImportMinutes(Number((event.target as HTMLInputElement).value));
  };

  const onDateChange = (event: InputEvent) => {
    const date = new Date((event.target as HTMLInputElement).value);
    setImportDate(date);
  };

  const isValidTime = (): boolean => {
    if (importHours() < 0 || importHours() > 23) {
      return false;
    }
    if (importMinutes() < 0 || importMinutes() > 59) {
      return false;
    }
    return true;
  };

  return (
    <form>
      <div class="form-check mb-2">
        <input
          class="form-check-input bg-secondary"
          type="checkbox"
          checked={autoImport()}
          id="flexCheckDefault"
          oninput={onAutoImportChange}
        />
        <label class="form-check-label" for="flexCheckDefault">
          Auto import
        </label>
      </div>
      <div class="container mb-3 p-0">
        <div class="form-group row">
          <label class="col col-form-lable">Time</label>
          <div class="col">
            <input
              type="text"
              class="form-control"
              id="hours"
              placeholder=""
              value={importHours()}
              oninput={onHoursChange}
            />
          </div>

          <div class="col">
            <input
              type="text"
              class="form-control"
              id="hours"
              placeholder=""
              value={importMinutes()}
              oninput={onMinutesChange}
            />
          </div>
        </div>
      </div>

      <button class="btn btn-outline-secondary me-3" onclick={submit}>
        Save
      </button>

      <hr />

      <div class="form-group row ps-2">
        <button
          class="col btn btn-outline-secondary me-3"
          onclick={importNowClick}
        >
          Import now
        </button>

        <label class="col col-form-lable pt-1">Date</label>
        <div class="col">
          <input
            type="date"
            class="form-control"
            id="date"
            placeholder=""
            value={formatDate(importDate())}
            oninput={onDateChange}
          />
        </div>
      </div>

      <Show when={isError()}>
        <div class="d-flex justify-content-center mb-3 text-danger">
          {errorMessage()}
        </div>
      </Show>
    </form>
  );
};

export default ImportSettings;
