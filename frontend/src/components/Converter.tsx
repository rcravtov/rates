import { Component, For, createEffect, createSignal } from 'solid-js';
import { Currency, Rate, calculateAmount, findRate, formatDate } from '../Utils';

interface Props {
    currencies: Function
    rateDate: Function
    setRateDate: Function
    rates: Function
}

const [sellAmount, setSellAmount] = createSignal(1)
const [sellCurrencyCode, setSellCurrencyCode] = createSignal(498)
const [buyAmount, setBuyAmount] = createSignal(2)
const [buyCurrencyCode, setBuyCurrencyCode] = createSignal(498)
const [reverseCalculation, setReverseCalculation] = createSignal(false)

const Converter: Component<Props> = (props) => {

    // calculate amount
    createEffect(()=>{        
        
        const sellRate = findRate(props.rates(),sellCurrencyCode())
        const buyRate = findRate(props.rates(),buyCurrencyCode())

        if(sellRate === undefined || buyRate === undefined) {
            if(reverseCalculation()) {
                setSellAmount(0)
            } else {
                setBuyAmount(0)
            }
            return
        }

        if(reverseCalculation()) {
            const amount = calculateAmount(buyAmount(), buyRate as Rate, sellRate as Rate)
            setSellAmount(amount)
        } else {
            const amount = calculateAmount(sellAmount(), sellRate as Rate, buyRate as Rate)
            setBuyAmount(amount)
        }

    })

    const onDateChange = (dateString: string) => {
        const date = new Date(dateString)
        props.setRateDate(date)
    }

    const onSellAmountChange = (value: string) => {
        if(isNaN(Number(value))) {
            return
        }
        setReverseCalculation(false)
        setSellAmount(Number(value))
    }

    const onSellCurrencyChange = (value: string) => {
        setReverseCalculation(false)
        setSellCurrencyCode(Number(value))
    }

    const onBuyAmountChange = (value: string) => {
        if(isNaN(Number(value))) {
            return
        }
        setReverseCalculation(true)
        setBuyAmount(Number(value))
    }

    const onBuyCurrencyChange = (value: string) => {
        setReverseCalculation(false)
        setBuyCurrencyCode(Number(value))
    }

    return (
        <div class='row g-3 mb-3'>
            <div class="col-auto">
                <label for="date" class="col-form-label">Date</label>
            </div>
            <div class="col-auto">
                <input type="date" 
                id="date" 
                class="form-control"
                value={formatDate(props.rateDate())}
                onChange={(event) => onDateChange(event.target.value)}/>
            </div>
            
            <div class="col-auto">
                <label for="sell-amount" class="col-form-label">Sell</label>
            </div>
            <div class="col">
                <input type="text" 
                id="sell-amount" 
                class="form-control"
                value={sellAmount()}
                onInput={(event) => onSellAmountChange(event.target.value)}/>
            </div>
            <div class="col-auto">
                <select 
                id="sell-currency" 
                class="form-control"
                value={sellCurrencyCode()}
                onChange={(event) => onSellCurrencyChange(event.target.value)}>
                    <For each={props.currencies()}>
                        {(currency: Currency) => <option value={currency.Code}>
                            {currency.CharCode}
                        </option>}
                    </For>
                </select>
            </div>

            <div class="col-auto">
                <label for="sell-amount" class="col-form-label">Buy</label>
            </div>
            <div class="col">
                <input type="text" 
                id="buy-amount" 
                class="form-control"
                value={buyAmount()}
                onInput={(event) => onBuyAmountChange(event.target.value)}/>
            </div>
            <div class="col-auto">
                <select 
                id="buy-currency" 
                class="form-control"
                value={buyCurrencyCode()}
                onChange={(event) => onBuyCurrencyChange(event.target.value)}>
                    <For each={props.currencies()}>
                        {(currency: Currency) => <option value={currency.Code}>
                            {currency.CharCode}
                        </option>}
                    </For>
                </select>
            </div>

        </div>
    )
}

export default Converter;