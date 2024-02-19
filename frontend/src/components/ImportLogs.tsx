import { Component, For, createEffect, createSignal } from "solid-js";
import {
  ImportLogPage,
  ImportLogRecord,
  getImportLogPage,
  formatDate,
} from "../Utils";
import { useGlobalContext } from "./GlobalContext";

const ImportLogs: Component = () => {
  const { token } = useGlobalContext();
  const [importLogPage, setImportLogPage] = createSignal({} as ImportLogPage);
  const [pageNumber, setPageNumber] = createSignal(0);

  createEffect(async () => {
    const response = await getImportLogPage(pageNumber(), token());
    setImportLogPage(response);
  });

  const onDisabledClick = (event: HTMLInputElement | MouseEvent) => {
    const ev = event as Event;
    const target = ev.target as HTMLInputElement;
    target.checked = !target.checked;
  };

  const pagesArray = (totalPages: number) => {
    const result:number[]=[];
    for(let i=0;i<totalPages;i++) {
      result.push(i+1);
    }
    return result;
  }

  const onPageSelect = (event: Event) => {
    const ev = event as Event;
    const target = ev.target as HTMLInputElement;
    setPageNumber(Number(target.value));
  }

  return (
    <>
      <form>
        <div class="form-group row mb-3">
          <label for="page" class="col-sm-2 col-form-label">
            Page
          </label>
          <div class="col-sm-10">
            <select
              class="form-control"
              id="page"
              onChange={onPageSelect}
            >
              <For each={pagesArray(importLogPage().totalPages)}>
                {(page: number) => <option>{page}</option>}
              </For>
            </select>
          </div>
        </div>
      </form>

      <table class="table">
        <thead>
          <tr class="table-secondary">
            <th scope="col">Date</th>
            <th scope="col">Rate date</th>
            <th scope="col">Auto</th>
            <th scope="col">Success</th>
            <th scope="col">Description</th>
          </tr>
        </thead>
        <tbody>
          <For each={importLogPage().data}>
            {(record: ImportLogRecord) => (
              <tr>
                <td>{formatDate(record.date, true)}</td>
                <td>{formatDate(record.data_date)}</td>
                <td>
                  <input
                    type="checkbox"
                    onClick={onDisabledClick}
                    checked={record.auto}
                  />
                </td>
                <td>
                  <input
                    type="checkbox"
                    onClick={onDisabledClick}
                    checked={record.success}
                  />
                </td>
                <td>{record.description}</td>
              </tr>
            )}
          </For>
        </tbody>
      </table>
    </>
  );
};

export default ImportLogs;
