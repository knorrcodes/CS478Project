import Vue from "vue";
import Router from "vue-router";
import LoginView from "@/views/Login.vue";
import MainMenu from "@/views/MainMenu.vue";
import CategoryView from "@/views/Category.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "mainMenu",
      component: MainMenu
    },
    {
      path: "/login",
      name: "loginScreen",
      component: LoginView,
      meta: { layout: "empty" }
    },
    {
      path: "/cat/:id",
      name: "category",
      component: CategoryView
    }
  ]
});
