<template>
  <div>
    <div class="container text-center">
      <h2>Categories</h2>
      <div class="row">
        <div class="col" align="center">
          <router-link-styled
            v-for="category in categories"
            v-bind:key="category.id"
            :value="category.name"
            :to="{path: '/cat/' + category.id}"
          ></router-link-styled>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { GET_ALL_CATEGORIES_QUERIES } from "@/graphql/queries/categoryQueries";
import ButtonStyled from "@/primatives/ButtonStyled.vue";
import RouterLinkStyled from "@/primatives/RouterLinkStyled.vue";
import { Category } from "@/graphql/schema";

@Component({
  components: {
    ButtonStyled,
    RouterLinkStyled
  },
  apollo: {
    categories: GET_ALL_CATEGORIES_QUERIES
  }
})
export default class MainMenu extends Vue {
  private categories: Category[] = [];
}
</script>

<style scoped>
</style>
