import {createLogger, createStore} from "vuex";

// @ts-ignore
const debug = process.env.NODE_ENV !== 'production'

const store = createStore({
    state () {
        return {
            count: 0
        }
    },
    mutations: {
        increment (state: any) {
            state.count++
        }
    },
    strict: debug,
    plugins: debug ? [createLogger()] : []
})

export default store