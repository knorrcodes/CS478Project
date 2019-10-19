<template>
  <div>
    <div class="container text-center">
      <h1>Categories</h1>
      <div class="row">
        <div class="col" align="center">
          <router-link
            v-for="category in categories"
            v-bind:key="category.id"
            class="btn btn-secondary mx-1 my-1"
            :to="{ path: '/cat/' + category.id }"
          >{{ category.name }}</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { GET_ALL_CATEGORIES_QUERIES } from "@/graphql/queries/categoryQueries";

@Component({
  apollo: {
    categories: GET_ALL_CATEGORIES_QUERIES
  }
})
export default class mainMenu extends Vue {}
</script>

<style scoped>
</style>
