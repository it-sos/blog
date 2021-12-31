import {createRouter, createWebHistory} from "vue-router";

const Home = () => import('../views/Home.vue')
const Article = () => import('../views/Article.vue')
const Story = () => import('../views/Story.vue')

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
