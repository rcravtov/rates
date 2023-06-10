
export type Currency = {
    Name: string
    Code: number
    CharCode: string
}

export type Rate = {
    Date: string
    Multiplier: number
    Value: number
    Change: number
    Currency: Currency
}

export const getCurrencies = async ():Promise<Currency[]> => {
    
    const res = await fetch(`http://localhost:8080/currencies`)
    
    let data = await res.json() as Currency[]

    return data
}

export const formatDate = (date: Date):string => {
    
    const day = date.getDate().toString().padStart(2, "0")
    const month = (date.getMonth()+1).toString().padStart(2, "0")
    const year = date.getFullYear()
    
    return `${year}-${month}-${day}`

}

export const getRates = async (date: Date):Promise<Rate[]> => {
    
    const dateString = formatDate(date)
    const res = await fetch(`http://localhost:8080/rates?date=${dateString}`)
    
    let data = await res.json() as Rate[]

    return data
}

export const filterRates = (rates: Rate[], filter: string): Rate[] => {
    
    const filterToLower = filter.toLowerCase()    

    if(filter.length > 0) {
        return rates?.filter(rate => { return (rate.Currency.Name.toLowerCase().includes(filterToLower)) ||
            (rate.Currency.Code.toString().includes(filterToLower)) || 
            (rate.Currency.CharCode.toLowerCase().includes(filterToLower))})
    }

    return rates

}

export const findRate = (rates: Rate[], currency_code: number):Rate|undefined => {
    return rates?.find(rate => rate.Currency.Code === currency_code)
}

export const calculateAmount = (initialAmount:number, rate1:Rate, rate2:Rate):number => {
    return Number((initialAmount * rate1.Value / rate1.Multiplier / rate2.Value * rate2.Multiplier).toFixed(2))
}