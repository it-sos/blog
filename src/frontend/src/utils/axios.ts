import axios, {AxiosRequestConfig} from 'axios'
import {ElMessage} from 'element-plus'
import {localGet} from './index'
// @ts-ignore
import config from '~/config'
import router from '../routes';
import {sign} from './sign'

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
    // @ts-ignore
    config.headers.ts = parseInt(Date.now()/1000)
    if (localGet('token') != null) {
        config.headers.token = localGet('token') || ""
    }
    var mixedData = {}
    Object.assign(mixedData, config.params, config.data);
    var s = sign(config.headers.ts, mixedData, config.headers.token)
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
    return res
}, res => {
    if (res.hasOwnProperty('response')) {
        // 需要登录授权，跳转登录页
        if (res.response.status == 401) {
            router.push({path: '/login'})
            return
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
