import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import BootstrapVue from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import VueRouter from 'vue-router';
import mainMenu from './components/mainMenu.vue';
import loginScreen from './components/loginScreen.vue';
Vue.config.productionTip = false;

Vue.use(BootstrapVue, VueRouter)

const routes = [
  { path: '/', component: loginScreen },
  { path: '/mainMenu', component: mainMenu },
];

const router = new VueRouter({
  routes
})

new Vue({
  store,
  router,
  render: (h) => h(App),
}).$mount('#app');

