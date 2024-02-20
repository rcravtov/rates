
export type Currency = {
    Name: string;
    Code: number;
    CharCode: string;
}

export type Rate = {
    Date: string;
    Multiplier: number;
    Value: number;
    Change: number;
    Currency: Currency;
}

export type AuthResult = {
    isError: boolean;
    ErrorMessage: string;
    Token: string;
}

export type AuthSettings = {
    login: string;
    token: string;
}

export type ImportSettings = {
    auto_import: boolean;
    import_hours: number;
    import_minutes: number;
}

export type GenericResult = {
    isError: boolean;
    ErrorMessage: string;
}

export type ImportLogRecord = {
    date: Date;
    data_date: Date;
    auto: boolean;
    success: boolean;
    description: string;
}

export type ImportLogPage = {
    totalPages: number;
    pageNumber: number;
    data: ImportLogRecord[];
}

export const apiURL = (): string => {
    return window.location.origin + "/api/"
}

export const getCurrencies = async (): Promise<Currency[]> => {

    const url = `${apiURL()}currencies`
    const res = await fetch(url)

    let data = await res.json() as Currency[]

    return data
}

export const formatDate = (date: Date, includeTime: boolean = false): string => {

    const secs = date.getSeconds().toString().padStart(2, "0");
    const mins = date.getMinutes().toString().padStart(2, "0");
    const hours = date.getHours().toString().padStart(2, "0");

    const day = date.getDate().toString().padStart(2, "0");
    const month = (date.getMonth() + 1).toString().padStart(2, "0");
    const year = date.getFullYear();

    if (includeTime) {
        return `${year}-${month}-${day} ${hours}:${mins}:${secs}`
    }

    return `${year}-${month}-${day}`

}

export const getRates = async (date: Date): Promise<Rate[]> => {

    const dateString = formatDate(date)
    const url = `${apiURL()}rates?date=${dateString}`
    const res = await fetch(url)

    let data = await res.json() as Rate[]

    return data
}

export const authorize = async (login: String, password: String): Promise<AuthResult> => {

    const url = `${apiURL()}auth`
    const data = {
        login: login,
        password: password
    }

    const res = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })

    const resJSON = await res.json();

    const authResult: AuthResult = {
        isError: false,
        ErrorMessage: "",
        Token: ""
    }

    if (res.status !== 200) {

        authResult.isError = true;
        authResult.ErrorMessage = "Bad login or password";
        return authResult;

    }

    authResult.Token = resJSON.token;

    return authResult;

}

export const getAuthSettings = async (token: string): Promise<AuthSettings> => {

    const url = `${apiURL()}admin/auth_settings`;

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    });

    let data = await res.json() as AuthSettings;

    return data;
}

export const changeAuth = async (login: String, password: String, token: string): Promise<AuthResult> => {

    const url = `${apiURL()}admin/auth_settings`
    const data = {
        login: login,
        password: password
    }

    const res = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(data)
    })

    const resJSON = await res.json();

    const authResult: AuthResult = {
        isError: false,
        ErrorMessage: "",
        Token: ""
    }

    if (res.status !== 200) {

        authResult.isError = true;
        authResult.ErrorMessage = "Error changing auth data";
        return authResult;

    }

    authResult.Token = resJSON.token;

    return authResult;

}

export const getImportSettings = async (token: string): Promise<ImportSettings> => {

    const url = `${apiURL()}admin/settings`;

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    });

    let data = await res.json() as ImportSettings;

    return data;
}

export const changeImportSettings = async (auto_import: Boolean, import_hours: Number, import_minutes: Number, token: string): Promise<GenericResult> => {

    const url = `${apiURL()}admin/settings`
    const data = {
        auto_import: auto_import,
        import_hours: import_hours,
        import_minutes: import_minutes,
    }

    const res = await fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(data)
    })

    const resJSON = await res.json();

    const result: GenericResult = {
        isError: false,
        ErrorMessage: "",
    }

    if (res.status !== 200) {

        result.isError = true;
        result.ErrorMessage = "Error changing settings";
        return result;

    }

    return result;

}

export const filterRates = (rates: Rate[], filter: string): Rate[] => {

    const filterToLower = filter.toLowerCase()

    if (filter.length > 0) {
        return rates?.filter(rate => {
            return (rate.Currency.Name.toLowerCase().includes(filterToLower)) ||
                (rate.Currency.Code.toString().includes(filterToLower)) ||
                (rate.Currency.CharCode.toLowerCase().includes(filterToLower))
        })
    }

    return rates

}

export const findRate = (rates: Rate[], currency_code: number): Rate | undefined => {
    return rates?.find(rate => rate.Currency.Code === currency_code)
}

export const calculateAmount = (initialAmount: number, rate1: Rate, rate2: Rate): number => {
    return Number((initialAmount * rate1.Value / rate1.Multiplier / rate2.Value * rate2.Multiplier).toFixed(2))
}

export const getImportLogPage = async (pageNumber: number, token: string): Promise<ImportLogPage> => {

    let limit = 20;
    let offset = (pageNumber - 1) * limit;

    const url = `${apiURL()}admin/import_logs?limit=${limit}&offset=${offset}`;

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    })

    const resJSON = await res.json();
    const data = resJSON.data as ImportLogRecord[];

    if (data != null) {
        for (let r of data) {
            r.date = new Date(String(r.date));
            r.data_date = new Date(String(r.data_date));
        }
    }
    const result: ImportLogPage = {
        totalPages: Math.ceil(resJSON.total / limit),
        pageNumber: pageNumber,
        data: data
    }

    return result;

}

export const importNow = async (date: Date, token: string): Promise<GenericResult> => {

    const url = `${apiURL()}admin/import?date=${formatDate(date)}`

    const res = await fetch(url, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`,
        }
    })

    //const resJSON = await res.json();

    const result: GenericResult = {
        isError: false,
        ErrorMessage: "",
    }

    // if (res.status !== 200) {

    //     result.isError = true;
    //     result.ErrorMessage = "Error importing";
    //     return result;

    // }

    return result;

}