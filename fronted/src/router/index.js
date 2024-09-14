import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/lines',
    component: () => import('../components/LineCreate.vue')
  },
  {
    path: '/info',
    component: () => import('../components/InfoForm.vue')
  },
  {
    path: '/schedule',
    component: () => import('../components/ScheduleTable.vue')
  },
  
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;