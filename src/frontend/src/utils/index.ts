export function localGet (key: any) {
    const value = window.localStorage.getItem(key)
    try {
        return JSON.parse(<string>window.localStorage.getItem(key))
    } catch (error) {
        return value
    }
}

export function localSet (key: any, value: any) {
    window.localStorage.setItem(key, JSON.stringify(value))
}

export function localRemove (key: string) {
    window.localStorage.removeItem(key)
}

export function localSessionSet(key: any, value: any) {
    window.sessionStorage.setItem(key, JSON.stringify(value))
}

export function localSessionGet (key: any) {
    const value = window.sessionStorage.getItem(key)
    try {
        return JSON.parse(<string>window.sessionStorage.getItem(key))
    } catch (error) {
        return value
    }
}
