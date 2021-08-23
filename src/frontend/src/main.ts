import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import 'element-plus/lib/theme-chalk/display.css'
import axios from 'axios'
import VueAxios from 'vue-axios'
import { router } from './routes'
import RightClick from './plugins/vue3-right-click-menu-element-plus'


import App from './App.vue'

if (process.env.NODE_ENV === "production") {
    axios.defaults.baseURL = '/api';
} else {
    axios.defaults.baseURL = 'http://localhost:8090';
}

const app = createApp(App)
app.use(ElementPlus)
    .use(RightClick)
    .use(router)
    .use(VueAxios, axios)

app.config.unwrapInjectedRef = true

app.mount('#app')

