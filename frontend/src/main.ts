import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import { apolloProvider } from "@/graphql/apollo";

import EmptyLayout from "@/layouts/Empty.vue";
import LoginLayout from "@/layouts/Login.vue";

import BootstrapVue from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";


Vue.component("login-layout", LoginLayout);
Vue.component("empty-layout", EmptyLayout);

Vue.config.productionTip = false;
Vue.use(BootstrapVue);

if (!localStorage.getItem("server-code")) {
  router.push({
    path: "/login"
  });
}

new Vue({
  apolloProvider,
  router,
  render: h => h(App)
}).$mount("#app");
