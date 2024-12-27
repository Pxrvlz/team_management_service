import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Employees from '../views/Employees.vue';
import Teams from '../views/Teams.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'خانه',
      component: Home
    },
    {
      path: '/employees',
      name: 'کارمند ها',
      component: Employees
    },
    {
      path: '/teams',
      name: 'تیم ها',
      component: Teams
    }
  ]
});

export default router;
