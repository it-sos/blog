import { InjectionKey } from 'vue'
import { createStore, useStore as baseUseStore, Store } from 'vuex'
import {localGet, localRemove, localSet} from "../utils";

export interface State {
    account: string
    token: string
}

export const key: InjectionKey<Store<State>> = Symbol()

const account = 'account'
const token = 'token'

export const store = createStore<State>({
    state: {
        account: '',
        token: ''
    },
    mutations: {
        logout(state: State): void {
            state.account = ""
            state.token = ""
            localRemove(account)
            localRemove(token)
        },
        login(state: State, payload: State): void {
            localSet(account, payload.account)
            localSet(token, payload.token)
            state.account = payload.account
            state.token = payload.token
        },
    },
    getters: {
        getAccount: (state: State) => (): string => {
            state.account = localGet(account)
            return state.account
        },
        getToken: (state: State) => (): string => {
            state.token = localGet(token) || ''
            return state.token
        },
    },
})

// 定义自己的 `useStore` 组合式函数
export function useStore () {
    return baseUseStore(key)
}
