<template>
  <div class="home">
    <div class="container">
      <div class="row">
        <div class="col-4 px-4">
          <table-order
            v-if="currentOrder"
            :currentOrder="currentOrder"
            :refetchFunc="refetchOrder"
          />
        </div>
        <div class="col-8">
          <category
            v-if="$route.params.id"
            :catId="$route.params.id"
            :addProductToOrder="addProductToOrder"
          />
          <main-menu v-else />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import TableOrder from "@/components/TableOrder.vue";
import Category from "@/views/Category.vue";
import MainMenu from "@/views/MainMenu.vue";
import { GET_CURRENT_TABLE } from "@/graphql/queries/tableQueries";
import {
  GET_LATEST_ORDER_QUERY,
  ADD_ITEMS_TO_ORDER_MUTATION,
  START_NEW_ORDER_MUTATION,
  CLOSE_ORDER_MUTATION
} from "@/graphql/queries/orderQueries";
import { GET_CURRENT_CUST_ORDER } from "@/graphql/queries/custCodeQueries";
import ButtonStyled from "@/primatives/Button.vue";
import { GET_CURRENT_TABLE_NOW } from "@/graphql/queries/tableQueries";

@Component({
  components: {
    TableOrder,
    Category,
    MainMenu,
    ButtonStyled
  },
  apollo: {
    currentOrder: {
      query: GET_CURRENT_CUST_ORDER,
      variables() {
        return {
          code: localStorage.getItem("customer-code")
        };
      },
      update: data => data.custcode.order
    }
  }
})
export default class InputOrder extends Vue {
  private currentTableId: number | null = null;
  private currentOrder: any = null;
  private currentOrderItem: number[] = [];
  private currentOrderItemCount: number = 0;

  private async addProductToOrder(productId: number, extraCount: number = 0) {
    this.currentOrderItem.push(productId);

    if (extraCount > 0) {
      this.currentOrderItemCount = extraCount;
    }

    if (this.currentOrderItem.length !== this.currentOrderItemCount + 1) {
      return;
    }

    await this.$apollo.mutate({
      mutation: ADD_ITEMS_TO_ORDER_MUTATION,
      variables: {
        order: this.currentOrder.id,
        products: this.currentOrderItem
      }
    });

    this.$apollo.queries.currentOrder.refetch();

    this.currentOrderItem = [];
    this.currentOrderItemCount = 0;
    this.$router.push({
      path: "/"
    });
  }

  private refetchOrder() {
    this.$apollo.queries.currentOrder.refetch();
  }
}
</script>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logoImage {
  margin-left: 5rem;
}

.logout-btn {
  margin-right: 5rem;
}
</style>
