import axios from 'axios'
import {ElMessage} from 'element-plus'
import {localGet} from './index'
import config from '~/config'
import router from '../routes';

// 这边由于后端没有区分测试和正式，姑且都写成一个接口。
axios.defaults.baseURL = config[import.meta.env.MODE].baseUrl
// 携带 cookie，对目前的项目没有什么作用，因为我们是 token 鉴权
axios.defaults.withCredentials = true
// 请求头，headers 信息
// @ts-ignore
axios.defaults.headers['X-Requested-With'] = 'XMLHttpRequest'
// @ts-ignore
axios.defaults.headers['token'] = localGet('token') || ''
// 默认 post 请求，使用 application/json 形式
axios.defaults.headers.post['Content-Type'] = 'application/json'

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
            ElMessage.error(res.response.data.message)
        } else {
            ElMessage.error(res.response.statusText)
        }
    } else {
        ElMessage.error('Network failure.')
    }
    return Promise.reject(res)
})

export default axios