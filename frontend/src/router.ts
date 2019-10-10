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
      path: '/chickenmenu',
      name: 'chickenMenu',
      component: chickenMenu,
    },
    {
      path: '/sidemenu',
      name: 'sideMenu',
      component: sideMenu,
    },
    {
      path: '/apatizermenu',
      name: 'apatizerMenu',
      component: apatizerMenu,
    },
    {
      path: '/burgermenu',
      name: 'burgerMenu',
      component: burgerMenu,
    },
    {
      path: '/dessertmenu',
      name: 'dessertMenu',
      component: dessertMenu,
    },
    {
      path: '/fishmenu',
      name: 'fishMenu',
      component: fishMenu,
    },
    {
      path: '/steakmenu',
      name: 'steakMenu',
      component: steakMenu,
    },
  ],
});
