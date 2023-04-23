import {createRouter, createWebHistory} from 'vue-router'

const routerHistory = createWebHistory()
import hw from '../views/hw.vue'
// 定义路由
const routes = [
    {
        path: '/',
        children: [
            {
                path: '',
                name: 'index',
                component: () => import('../components/MainPage.vue'),
            },
            {
                path: 'clusters',
                children: [
                    {
                        path: '',
                        name: 'clusters',
                        component: () => import('../components/MainPage.vue'),
                    },
                    {
                        path: 'detail',
                        name: 'detail',
                        component: () => import('../components/MainPage.vue')
                    },
                  ],
            },
          ],
    },
    {
        path: '/hw',
        name: 'hw',
        component: hw
    },
    {
        path: '/:catchAll(.*)',
        name: '/404',
        component: () => import('../views/404.vue')
    },
]

// 创建路由器
const router = createRouter({
    history: routerHistory,
    routes
})


export default router;
