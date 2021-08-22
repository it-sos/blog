import {createApp} from 'vue'
import RightMenu from './RightMenu.vue'

interface Option {
    title?: string;
    content?: string;
}

export interface Position {
    x: number
    y: number
}

function mountContent(option = {} as Option) {
    const app = createApp(RightMenu, {
        close: () => { // @ts-ignore
            app.unmount('body')
        },
        ...option
    })

    app.directive('rightClick', {
        mounted(el, binding) {
            el.addEventListener("contextmenu", function (ev: any) {
                ev.preventDefault();
                const position: Position = {
                    x: ev.x,
                    y: ev.y,
                }
                binding.value(position)
                return false;
            }, false)
        }
    })

    app.mount('body')
}

export default mountContent
