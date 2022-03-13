import {localGet, localSessionGet, localSessionSet} from "../../utils";

export interface State {
    status: number
}

const session_is_first_loading = 'firstloading'

export default {
    namespaced: true,
    state: () => ({
        status: 0
    }),
    mutations: {
        setFirstLoading(state: State) {
            state.status = 1
            localSessionSet(session_is_first_loading, state.status)
        }
    },
    getters: {
        getIsFirstLoading (state: State) : boolean {
            state.status = localSessionGet(session_is_first_loading) || 0
            return state.status == 0
        }
    }
}
