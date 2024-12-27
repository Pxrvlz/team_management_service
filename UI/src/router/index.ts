import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Employees from '../views/Employees.vue';
import Teams from '../views/Teams.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/employees',
      name: 'employees',
      component: Employees
    },
    {
      path: '/teams',
      name: 'teams',
      component: Teams
    }
  ]
});

export default router;