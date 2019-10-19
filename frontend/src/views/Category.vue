<template>
  <div v-if="$apollo.loading">Loading...</div>
  <div v-else class="container text-center">
    <h1>{{ categoryData.name }} Menu</h1>

    <button type="button" @click="goBack" class="btn btn-secondary mx-1 my-1">&lt;- Back</button>

    <section class="products">
      <button
        v-for="product in categoryData.products"
        v-bind:key="product.id"
        class="btn btn-secondary mx-1 my-1"
        type="button"
        @click="addProductToOrder(product.id)"
      >{{ product.name }}</button>
    </section>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { GET_PRODUCTS_IN_CATEGORY_QUERIES } from "@/graphql/queries/categoryQueries";

@Component({
  apollo: {
    categoryData: {
      query: GET_PRODUCTS_IN_CATEGORY_QUERIES,
      update: data => data.category,
      variables() {
        return {
          id: this.$route.params.id
        };
      }
    }
  }
})
export default class CategoryView extends Vue {
  private addProductToOrder(productId: number) {
    console.log(productId);
  }

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
