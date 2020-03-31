import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'),
    meta: {
      title: 'Dashboard - OpenLAN'
    },
  },
  {
    path: '/vswitch',
    name: 'VSwitch',
    component: () => import('../views/VSwitch.vue'),
    meta: {
      title: 'Virtual Switch - OpenLAN'
    },
  },
  {
    path: '/point',
    name: 'Point',
    component: () => import('../views/Point.vue'),
    meta: {
      title: 'Access Point - OpenLAN'
    },
  },
  {
    path: '/topo',
    name: 'Topo',
    component: () => import('../views/Graph.vue'),
    meta: {
      title: 'Topology - OpenLAN'
    },
  }
];

const router = new VueRouter({
  routes
});

router.afterEach((to, from) => {
  window.document.title = to.meta.title || to.name;
});

export default router;
