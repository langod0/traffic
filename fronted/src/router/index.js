import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    component: () => import('../components/HomePage.vue')
  },
  {
    path: '/login',
    component: () => import('../components/LoginPage.vue')
  },
  {
    path: '/register',
    component: () => import('../components/Register.vue')
  },
  {
    path: '/main',
    component: () => import('../components/Main.vue')
  },
  {
    path: '/main2',
    component: () => import('../components/Main2.vue')
  },
  {
    path: '/lines',
    component: () => import('../components/items/LineCreate.vue')
  },
  {
    path: '/lnfo',
    component: () => import('../components/items/InfoForm.vue')
  },
  {
    path: '/schedule',
    component: () => import('../components/items/ScheduleTable.vue')
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;