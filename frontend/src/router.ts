import Vue from "vue";
import Router from "vue-router";
import LoginView from "@/views/Login.vue";
import MainMenuView from "@/views/MainMenu.vue";
import CategoryView from "@/views/Category.vue";
import TablesView from "@/views/Tables.vue";
import InputOrderView from "@/views/InputOrder.vue";

Vue.use(Router);

export default new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes: [
    {
      path: "/",
      name: "mainMenu",
      component: InputOrderView
    },
    {
      path: "/tables",
      name: "tables",
      component: TablesView
    },
    {
      path: "/login",
      name: "loginScreen",
      component: LoginView,
      meta: { layout: "login" }
    },
    {
      path: "/cat/:id",
      name: "category",
      component: InputOrderView
    }
  ]
});
