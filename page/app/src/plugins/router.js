import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

const Home = () => import('../views/home.vue')
const Auth = () => import('../views/auth.vue')
const List = () => import('../views/list.vue')
const Info = () => import('../views/info.vue')

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            name: 'Home',
            path: '/',
            component: Home,
            meta: {
                title: '首页'
            }
        },
        {
            name: 'Auth',
            path: '/auth',
            component: Auth,
            meta: {
                title: '登录'
            }
        }, 
        {
            name: 'List',
            path: '/list',
            component: List,
            meta: {
                title: '文章列表'
            }
        }, 
        {
            name: 'Info',
            path: '/info/:id',
            component: Info,
            meta: {
                title: '文章详情'
            }
        }
    ],
})

export default router