<template>
  <div v-if="$apollo.loading || !categoryData">Loading...</div>
  <div v-else class="container text-center">
    <h2>{{ categoryData.name }} Menu</h2>

    <button-styled :clickHandler="() => goBack()" value="&lt;- Back"></button-styled>

    <section class="products">
      <button-styled
        v-for="product in categoryData.products"
        v-bind:key="product.id"
        :clickHandler="() => addProductItem(product.id)"
        :value="product.name"
      ></button-styled>
    </section>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { GET_PRODUCTS_IN_CATEGORY_QUERIES } from "@/graphql/queries/categoryQueries";
import { GET_CURRENT_TABLE } from "@/graphql/queries/tableQueries";
import { GET_ALL_CATEGORIES_QUERIES } from "@/graphql/queries/categoryQueries";
import { Category } from "@/graphql/schema";

import RouterLinkStyled from "@/primatives/RouterLinkStyled.vue";
import ButtonStyled from "@/primatives/ButtonStyled.vue";

@Component({
  components: {
    RouterLinkStyled,
    ButtonStyled
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
  @Prop() private addProductToOrder!: (
    productId: number,
    extraCount: number
  ) => void;
  @Prop() private readonly catId!: number;
  private categoryData: Category | null = null;

  private goBack() {
    this.$router.back();
  }

  private addProductItem(id: number) {
    if (!this.categoryData || !this.addProductToOrder) {
      return;
    }
    const product = this.categoryData.products.find(
      (item: any) => item.id === id
    );

    if (!product) {
      return;
    }

    this.addProductToOrder(id, product.num_of_sides);

    if (product.num_of_sides > 0) {
      this.$router.push({
        path: "/cat/7"
      });
    }
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
