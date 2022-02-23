import {reactive} from "vue";

const store = {
    debug: true,

    state: reactive({
        message: 'Hello!'
    }),

    setMessageAction(newValue) {
        if (this.debug) {
            console.log('setMessageAction triggered with', newValue)
        }

        this.state.message = newValue
    },

    clearMessageAction() {
        if (this.debug) {
            console.log('clearMessageAction triggered')
        }

        this.state.message = ''
    }
}
export default store
