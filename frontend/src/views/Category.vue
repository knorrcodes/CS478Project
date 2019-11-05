<template>
  <div v-if="$apollo.loading">Loading...</div>
  <div v-else class="container text-center">
    <h1>{{ categoryData.name }} Menu</h1>

    <button-c :clickHandler="() => goBack()" value="&lt;- Back"></button-c>

    <section class="products">
      <router-c
        v-for="product in categoryData.products"
        v-bind:key="product.id"
        :clickHandler="() => addProductToOrder(product.id)"
        :value="product.name"
        :to="{path: '/cat/' + 7}"
      ></router-c>
    </section>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { GET_PRODUCTS_IN_CATEGORY_QUERIES } from "@/graphql/queries/categoryQueries";
import { GET_CURRENT_TABLE } from "@/graphql/queries/tableQueries";
import { GET_ALL_CATEGORIES_QUERIES } from "@/graphql/queries/categoryQueries";

import RouterC from "@/primatives/RouterLink.vue";
import ButtonC from "@/primatives/Button.vue";

@Component({
  components: {
    RouterC,
    ButtonC
  },
  apollo: {
    categoryData: {
      query: GET_PRODUCTS_IN_CATEGORY_QUERIES,
      update: data => data.category,
      variables() {
        return {
          id: this.catId
        };
      }
    },
    categories: GET_ALL_CATEGORIES_QUERIES
  }
})
export default class CategoryView extends Vue {
  @Prop() private addProductToOrder: any;
  @Prop() private readonly catId: any;
  private categoryData: any = null;

  private goBack() {
    this.$router.back();
  }
}
</script>

<style scoped>
h1 {
  text-align: center;
  display: inline-block;
  margin-right: 1rem;
}
</style>
