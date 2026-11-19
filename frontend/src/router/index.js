import Vue from 'vue'
import VueRouter from 'vue-router'
import authRoutes from '@/router/routes/auth';
import AuthService from "@/services/auth.service";
import store from "@/store";

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/home/Home.vue'),
      meta: {
        requiresAuth: true,
        showHeader: true,
      }
    },
    ...authRoutes,
    {
      path: '/404',
      name: 'error-404-page',
      component: () => import('@/views/error/Error404.vue'),
      meta: {
        requiresAuth: true,
        showHeader: true,
      }
    }
  ]
})

router.beforeEach(async(to, from, next) => {
  try {
    if (AuthService.isAuthenticated()) {
      await store.dispatch('auth/fetchUser');
    }
  } catch (error) {
    throw error;
  }

  if (to.matched.length === 0) {
    return next(AuthService.isAuthenticated() ? '/404' : '/login');
  }

  if (to.matched.some(record => record.meta.requiresAdmin)) {
    if (!AuthService.isAuthenticated() || !AuthService.isAdmin()) {
      return next('/404');
    }
  }

  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!AuthService.isAuthenticated()) {
      return next('/login');
    }
  }

  if (to.matched.some(record => record.meta.guestOnly)) {
    if (AuthService.isAuthenticated()) {
      return next('/');
    }
  }

  next();
});

export default router
