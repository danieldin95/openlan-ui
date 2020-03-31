import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/vswitch',
    name: 'VSwitch',
    component: () => import('../views/VSwitch.vue')
  },
  {
    path: '/point',
    name: 'Point',
    component: () => import('../views/Point.vue')
  },
  {
    path: '/graph',
    name: 'Graph',
    component: () => import('../views/Graph.vue')
  }
];

const router = new VueRouter({
  routes
});

export default router;
