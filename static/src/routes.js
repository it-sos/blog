import {createRouter, createWebHistory} from "vue-router";

const Home = () => import(/* webpackChunkName: "group-front" */ './components/Home.vue')
const Article = () => import(/* webpackChunkName: "group-front" */ './components/Article.vue')

const routes = [
    { path: '/', component: Home },
    { path: '/:keyword', component: Home },
    { path: '/a/:title', component: Article },
]

export const router = createRouter({
    // 4. 内部提供了 history 模式的实现
    history: createWebHistory(),
    routes, // `routes: routes` 的缩写
})