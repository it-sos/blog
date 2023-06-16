interface State {
    data: string[]
}

export default {
    namespaced: true,
    state: () => ({
        data: []
    }),
    mutations: {
        send(state: State, data: string) {
            state.data.push(data)
        },
        remove(state: State, index: number) {
            state.data.splice(index, 1);
        },
        reset(state: State) {
            state.data = []
        },
    },
    getters: {
        list(state: State): string[] {
            return state.data
        }
    }
}
