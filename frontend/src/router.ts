import Vue from 'vue';
import Router from 'vue-router';
import loginScreen from './components/loginScreen.vue';
import mainMenu from './components/mainMenu.vue';
import chickenMenu from './components/entreeItems/chickenMenu.vue';
import sideMenu from './components/sideMenu.vue';
import apatizerMenu from './components/entreeItems/apatizerMenu.vue';
import burgerMenu from './components/entreeItems/burgerMenu.vue';
import dessertMenu from './components/entreeItems/dessertMenu.vue';
import fishMenu from './components/entreeItems/fishMenu.vue';
import steakMenu from './components/entreeItems/steakMenu.vue';

Vue.use(Router);

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'loginScreen',
      component: loginScreen,
    },
    {
      path: '/mainMenu',
      name: 'mainMenu',
      component: mainMenu,
    },
    {
      path: '/chickenMenu',
      name: 'chickenMenu',
      component: chickenMenu,
    },
    {
      path: '/sideMenu',
      name: 'sideMenu',
      component: sideMenu,
    },
    {
      path: '/apatizerMenu',
      name: 'apatizerMenu',
      component: apatizerMenu,
    },
    {
      path: '/burgerMenu',
      name: 'burgerMenu',
      component: burgerMenu,
    },
    {
      path: '/dessertMenu',
      name: 'dessertMenu',
      component: dessertMenu,
    },
    {
      path: '/fishMenu',
      name: 'fishMenu',
      component: fishMenu,
    },
    {
      path: '/steakMenu',
      name: 'steakMenu',
      component: steakMenu,
    },
  ],
});
