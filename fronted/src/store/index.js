import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  { path: '/', component: () => import('@/components/ScheduleForm.vue') },
  { path: '/schedule', component: () => import('@/components/ScheduleTable.vue') },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;