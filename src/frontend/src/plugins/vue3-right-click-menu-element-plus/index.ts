import {App} from "vue";

import RightMenu from './RightMenu.vue'

export default {
    install: (app: App, opt: any) => {
        app.component('right-menu', RightMenu)
        app.directive('rightClick', {
            mounted(el :any, binding: any) {
                el.addEventListener("contextmenu", function (ev: any) {
                    ev.preventDefault();
                    const _value = binding.value()
                    _value.show = true
                    _value.position.x = ev.x
                    _value.position.y = ev.y
                    document.onclick = null
                    document.onclick = () => {
                        _value.show = false
                        document.onclick = null
                    }
                    return false;
                }, false)
            }
        })
    }
}

