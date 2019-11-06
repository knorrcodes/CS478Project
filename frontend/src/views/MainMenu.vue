<template>
  <div>
    <div class="container text-center">
      <h1>Categories</h1>
      <div class="row">
        <div class="col" align="center">
          <router-c
            v-for="category in categories"
            v-bind:key="category.id"
            :value="category.name"
            :to="{path: '/cat/' + category.id}"
          ></router-c>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { GET_ALL_CATEGORIES_QUERIES } from "@/graphql/queries/categoryQueries";
import ButtonC from "@/primatives/Button.vue";
import RouterC from "@/primatives/RouterLink.vue";

@Component({
  components: {
    ButtonC,
    RouterC
  },
  apollo: {
    categories: GET_ALL_CATEGORIES_QUERIES
  },
  methods: {
    test() {
      console.log("hello");
    }
  }
})
export default class MainMenu extends Vue {}
</script>

<style scoped>
</style>
