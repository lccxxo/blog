import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Home.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/Register.vue')
  },
  {
    path: '/article/:id',
    name: 'ArticleDetail',
    component: () => import('@/views/ArticleDetail.vue')
  },
  {
    path: '/create-article',
    name: 'CreateArticle',
    component: () => import('@/views/CreateArticle.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/edit-article/:id',
    name: 'EditArticle',
    component: () => import('@/views/EditArticle.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/my-articles',
    name: 'MyArticles',
    component: () => import('@/views/MyArticles.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/categories',
    name: 'Categories',
    component: () => import('@/views/Categories.vue')
  },
  {
    path: '/tags',
    name: 'Tags',
    component: () => import('@/views/Tags.vue')
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/Profile.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/my-comments',
    name: 'MyComments',
    component: () => import('@/views/MyComments.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  // 如果访问需要认证的页面但未登录，重定向到登录页
  if (to.meta.requiresAuth && !userStore.isLoggedIn) {
    next('/login')
  } 
  // 如果已登录用户访问登录/注册页，重定向到首页
  else if ((to.path === '/login' || to.path === '/register') && userStore.isLoggedIn) {
    next('/')
  } 
  else {
    next()
  }
})

export default router

