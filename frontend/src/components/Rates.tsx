import { Component, For, createEffect, createSignal } from 'solid-js';
import { Rate, filterRates, getRates, formatDate, findRate, calculateAmount, Currency, getCurrencies } from '../Utils';
import RateList from './RateList';
import Converter from './Converter';

const [currencies, setCurrencies] = createSignal([] as Currency[])
const [rateDate, setRateDate] = createSignal(new Date())
const [rates, setRates] = createSignal([] as Rate[])

const Rates: Component = () => {

    // get currencies
    createEffect(async ()=>{        
        const response= await getCurrencies()
        setCurrencies(response)        
    })

    // get rates on date change
    createEffect(async ()=>{        
        const response= await getRates(rateDate())
        setRates(response)
    })

    return (
        <div>
            <Converter 
                currencies={currencies}
                rateDate={rateDate}
                setRateDate={setRateDate}
                rates={rates}
            />
            <hr/>
            <RateList rates={rates}/>
        </div>
    )
}

export default Rates;