import {
    createRouter,
    createWebHashHistory
} from 'vue-router'

const Home = () => import('../views/home.vue')
const Install = () => import('../views/install.vue')
const Auth = () => import('../views/auth.vue')
const AnyList = () => import('../views/any.vue')
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
            name: 'Install',
            path: '/install',
            component: Install,
            meta: {
                title: '初始化配置'
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
            name: 'Any',
            path: '/any',
            component: AnyList,
            meta: {
                title: '随便看看'
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
                title: '加载中'
            }
        }
    ],
})

router.beforeEach((to, _, next) => {
    let suffix = 'Blog'
    let cache = localStorage.getItem('cache:user');
    if (cache != undefined) {
        cache = JSON.parse(cache)
        if (cache.nickname) suffix = cache.nickname + ' ' + suffix
        else suffix = 'BetaX ' + suffix
    } else suffix = 'BetaX ' + suffix
    document.title = (to.meta.title === undefined ? '未知页面 - ' : to.meta.title + ' - ') + suffix
    next()
})

export default router