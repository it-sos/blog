import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import 'element-plus/lib/theme-chalk/display.css'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VueUeditorWrap from 'vue-ueditor-wrap'
import { router } from './routes'
import App from './App.vue'

axios.defaults.baseURL = 'http://localhost:8080';

const app = createApp(App)
app.use(ElementPlus)
    .use(router)
    .use(VueAxios, axios)
    .use(VueUeditorWrap)

app.mount('#app')
