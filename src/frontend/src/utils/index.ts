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