import { InjectionKey } from 'vue'
import { createStore, useStore as baseUseStore, Store } from 'vuex'
import {localGet, localRemove, localSet} from "../utils";

export interface State {
    account: string
    token: string
}

export const key: InjectionKey<Store<State>> = Symbol()

export const store = createStore<State>({
    state: {
        account: 'account',
        token: 'token'
    },
    mutations: {
        logout(state: State): void {
            localRemove(state.account)
            localRemove(state.token)
        },
        login(state: State, payload: State): void {
            localSet(state.account, payload.account)
            localSet(state.token, payload.token)
        },
    },
    getters: {
        getAccount: (state: State) => (): string => {
            return localGet(state.account)
        },
        getToken: (state: State) => (): string => {
            return localGet(state.token)
        },
    },
})

// 定义自己的 `useStore` 组合式函数
export function useStore () {
    return baseUseStore(key)
}
