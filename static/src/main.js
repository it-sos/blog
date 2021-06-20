import { createApp } from 'vue';
import ElementPlus from 'element-plus';
import 'element-plus/lib/theme-chalk/index.css';
import 'element-plus/lib/theme-chalk/display.css';
import { router } from './routes'
import App from './App.vue';

const app = createApp(App)
app.use(ElementPlus)
app.use(router)

app.mount('#app')