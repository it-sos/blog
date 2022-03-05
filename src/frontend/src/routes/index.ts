import {createRouter, createWebHistory} from "vue-router";
import global from "./global";

const Home = () => import('../views/Home.vue')
const Article = () => import('../views/Article.vue')
const Story = () => import('../views/Story.vue')
const Login = () => import('../views/Login.vue')

const routes = [
    { path: '/', component: Home, meta: {requiresAuth: false} },
    { path: '/login', component: Login, meta: {requiresAuth: false} },
    { path: '/:keyword', component: Home, meta: {requiresAuth: false} },
    { path: '/a/:title', component: Article, meta: {requiresAuth: false} },
    { path: '/e/', component: Story, meta: {requiresAuth: true} },
    { path: '/e/:id', component: Story, meta: {requiresAuth: true} },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

declare module 'vue-router' {
    interface RouteMeta {
        // 是可选的
        isAdmin?: boolean
        // 每个路由都必须声明
        requiresAuth: boolean
    }
}

global(router)

export default router
