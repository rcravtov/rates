import { Component, For, createEffect, createSignal } from "solid-js";
import { Rate, filterRates } from "../Utils";

interface Props {
  rates: Function;
}

const [filteredRates, setFilteredRates] = createSignal([] as Rate[]);
const [filter, setFilter] = createSignal("");

const RateList: Component<Props> = (props) => {
  // filter rates on filter change
  createEffect(async () => {
    const response = filterRates(props.rates(), filter());
    setFilteredRates(response);
  });

  const onFilterChange = (filter: string) => {
    setFilter(filter);
  };

  const getChangeColor = (change: number) => {
    if (change === 0) {
      return "";
    } else if (change > 0) {
      return "text-success";
    } else {
      return "text-danger";
    }
  };

  return (
    <>
      <div class="row g-3 align-items-center mb-3">
        <div class="col-auto">
          <label for="filter" class="col-form-label">
            Filter
          </label>
        </div>
        <div class="col">
          <input
            type="text"
            id="filter"
            class="form-control"
            onInput={(event) => onFilterChange(event.target.value)}
          />
        </div>
      </div>

      <table class="table">
        <thead>
          <tr class="table-secondary">
            <th scope="col">Code</th>
            <th scope="col">Char code</th>
            <th scope="col">Name</th>
            <th scope="col">Multiplier</th>
            <th scope="col">Value</th>
            <th scope="col">Change</th>
          </tr>
        </thead>
        <tbody>
          <For each={filteredRates()}>
            {(rate: Rate) => (
              <tr>
                <td>{rate.Currency.Code}</td>
                <td>{rate.Currency.CharCode}</td>
                <td>{rate.Currency.Name}</td>
                <td>{rate.Multiplier}</td>
                <td>{rate.Value}</td>
                <td class={getChangeColor(rate.Change)}>
                  {rate.Change.toFixed(4)}
                </td>
              </tr>
            )}
          </For>
        </tbody>
      </table>
    </>
  );
};

export default RateList;
