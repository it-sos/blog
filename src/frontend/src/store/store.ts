import { InjectionKey } from 'vue'
import { createStore, useStore as baseUseStore, Store } from 'vuex'
import {localGet, localRemove, localSessionGet, localSessionSet, localSet} from "../utils";

export interface State {
    account: string
    token: string
    articleId: number
    saved: boolean
    ts: number
}

export const key: InjectionKey<Store<State>> = Symbol()

const account = 'account'
const token = 'token'
const session_key_ts = 'ts'

export const store = createStore<State>({
    state: {
        account: '',
        token: '',
        articleId: 0,
        saved: true,
        ts: 0,
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
        setArticleId(state: State, articleId: number): void {
            state.articleId = articleId
        },
        setTs(state: State, ts: number): void {
            localSessionSet(session_key_ts, ts)
            state.ts = ts
        },
        setSaved(state: State, is: boolean): void {
           state.saved = is
        }
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
        getArticleId: (state: State) => (): number => {
            return state.articleId
        },
        getTs: (state: State) => (): number => {
            state.ts = localSessionGet(session_key_ts)
            return state.ts
        },
        getSaved: (state: State) => (): boolean => {
            return state.saved
        }
    },
})

// 定义自己的 `useStore` 组合式函数
export function useStore () {
    return baseUseStore(key)
}
