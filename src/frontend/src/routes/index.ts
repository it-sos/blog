import {createRouter, createWebHistory} from "vue-router";

const Home = () => import(/* webpackChunkName: "Home" */ '../views/Home.vue')
const Article = () => import(/* webpackChunkName: "Article" */ '../views/Article.vue')
const Story = () => import(/* webpackChunkName: "Story" */ '../views/Story.vue')

const routes = [
    { path: '/', component: Home },
    { path: '/:keyword', component: Home },
    { path: '/a/:title', component: Article },
    { path: '/e/', component: Story },
    { path: '/e/:id', component: Story },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router