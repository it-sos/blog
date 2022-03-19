import {store} from "../store/store";

export default (router: any) => {
    router.beforeEach((to: any, from: any) => {
        // 而不是去检查每条路由记录
        // to.matched.some(record => record.meta.requiresAuth)
        if (to.meta.requiresAuth && store.getters.getToken() == '') {
            return {
                path: '/login',
                // 保存我们所在的位置，以便以后再来
                // query: { redirect: to.fullPath },
            }
        }
    })
}
