import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

const Home = () => import('../views/home.vue')
const List = () => import('../views/list.vue')
const Edit = () => import('../views/edit.vue')

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
            name: 'List',
            path: '/list',
            component: List,
            meta: {
                title: '文章列表'
            }
        }, 
        {
            name: 'Add',
            path: '/add',
            component: Edit,
            meta: {
                title: '撰写新文章'
            }
        }, 
        {
            name: 'Edit',
            path: '/edit/:id',
            component: Edit,
            meta: {
                title: '编辑文章'
            }
        }
    ],
})

export default router