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
import ButtonStyled from "@/primatives/Button.vue";
import RouterLinkStyled from "@/primatives/RouterLink.vue";

@Component({
  components: {
    ButtonStyled,
    RouterLinkStyled
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
export default class MainMenu extends Vue {
  private categories: any = null;
}
</script>

<style scoped>
</style>
