import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import VueAxios from 'vue-axios'
import router from './routes'
import RightClick from './plugins/vue3-right-click-menu-element-plus'


import App from './App.vue'
import axios from './utils/axios'

const app = createApp(App)
app.use(ElementPlus)
    .use(RightClick)
    .use(router)
    .use(VueAxios, axios)

app.config.globalProperties.axios=axios
app.config.unwrapInjectedRef = true

app.mount('#app')