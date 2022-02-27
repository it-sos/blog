import axios from 'axios'
import {ElMessage} from 'element-plus'
// @ts-ignore
import config from '~/config'
import {sign} from './sign'
import {store} from "../store/store";
import {useRouter} from "vue-router";

const router = useRouter()

// 这边由于后端没有区分测试和正式，姑且都写成一个接口。
axios.defaults.baseURL = config[import.meta.env.MODE].baseUrl
// 携带 cookie，对目前的项目没有什么作用，因为我们是 token 鉴权
axios.defaults.withCredentials = true
// 请求头，headers 信息
// @ts-ignore
axios.defaults.headers['X-Requested-With'] = 'XMLHttpRequest'
// @ts-ignore
axios.defaults.headers['token'] = ""
// 默认 post 请求，使用 application/json 形式
axios.defaults.headers.post['Content-Type'] = 'application/json'

axios.interceptors.request.use((config: any) => {
    let ts = store.getters.getTs()
    if (ts > 0) {
        config.headers.ts = ts
    } else {
        // @ts-ignore
        config.headers.ts = parseInt(Date.now() / 1000)
    }
    config.headers.token = store.getters.getToken()
    let mixedData = {}
    Object.assign(mixedData, config.params, config.data);
    let s = sign(config.headers.ts, mixedData, config.headers.token)
    config.headers.nonce = s[0]
    config.headers.sign = s[1]

    return config
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
})

// 请求拦截器，内部根据返回值，重新组装，统一管理。
// @ts-ignore
axios.interceptors.response.use((res: HttpResult) => {
    // 表示均从服务端获取时间戳
    if (store.getters.getTs() > 0) {
        store.commit('setTs', res.response.headers['x-time'])
    }
    return res
}, res => {
    if (res.hasOwnProperty('response')) {
        // 需要登录授权，跳转登录页
        if (res.response.status == 401) {
            store.commit("logout")
            router.push({path: '/login'})
            return Promise.reject(res)
        } else if (res.response.status == 403) {
            // @ts-ignore
            if (parseInt(Date.now() / 1000) - res.response.headers['x-time'] < -290) {
                store.commit('setTs', res.response.headers['x-time'])
                location.reload()
                return Promise.reject(res)
            }
        }
        if (res.response.hasOwnProperty('data')) {
            if (res.response.data.message != null) {
                ElMessage.error(res.response.data.message)
            } else {
                ElMessage.error(res.response.statusText)
            }
        } else {
            ElMessage.error(res.response.statusText)
        }
    } else {
        ElMessage.error('Network failure.')
    }
    return Promise.reject(res)
})

export default axios
