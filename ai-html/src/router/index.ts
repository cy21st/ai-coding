import { createRouter, createWebHashHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/login/index.vue')
  },
  {
    path: '/',
    name: 'Layout',
    component: () => import('../layout/index.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('../views/dashboard/index.vue'),
        meta: { title: '仪表盘', icon: 'dashboard' }
      },
      {
        path: 'metadata',
        name: 'Metadata',
        component: () => import('../views/metadata/index.vue'),
        meta: { title: '元数据管理', icon: 'database' },
        redirect: '/metadata/events',
        children: [
          {
            path: 'events',
            name: 'Events',
            component: () => import('../views/metadata/events/index.vue'),
            meta: { title: '事件管理' }
          },
          {
            path: 'attributes',
            name: 'Attributes',
            component: () => import('../views/metadata/attributes/index.vue'),
            meta: { title: '属性管理' }
          },
          {
            path: 'relations',
            name: 'Relations',
            component: () => import('../views/metadata/relations/index.vue'),
            meta: { title: '关联管理' }
          }
        ]
      },
      {
        path: 'user',
        name: 'User',
        component: () => import('../views/user/index.vue'),
        meta: { title: '用户管理', icon: 'user' }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, _, next) => {
  const token = localStorage.getItem('token')
  if (to.path === '/login') {
    next()
  } else {
    if (!token) {
      next('/login')
    } else {
      next()
    }
  }
})

export default router 