import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import { createApp } from 'vue'
import VueAxios from 'vue-axios'
import RightClick from './plugins/vue3-right-click-menu-element-plus'
import router from './routes'
import { key, store } from './store/store'

import App from './App.vue'
import axios from './utils/axios'

const app = createApp(App)
app.use(ElementPlus)
    .use(RightClick)
    .use(router)
    .use(store, key)
    .use(VueAxios, axios)

app.provide('$axios', axios)
app.config.unwrapInjectedRef = true

app.mount('#app')
