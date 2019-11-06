<template>
  <div class="home">
    <div class="container">
      <div class="row">
        <div class="col-4 px-4">
          <button-styled :clickHandler="() => changeTables()" value="Change Table"></button-styled>
          <button-styled :clickHandler="() =>closeOrder()" value="Close Order"></button-styled>
          <start-table-order v-if="!currentOrder" :startOrder="startNewOrder"></start-table-order>
          <table-order v-else :currentOrder="currentOrder" />
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
import StartTableOrder from "@/components/StartTableOrder.vue";
import Category from "@/views/Category.vue";
import MainMenu from "@/views/MainMenu.vue";
import { GET_CURRENT_TABLE } from "@/graphql/queries/tableQueries";
import {
  GET_LATEST_ORDER_QUERY,
  ADD_ITEMS_TO_ORDER_MUTATION,
  START_NEW_ORDER_MUTATION,
  CLOSE_ORDER_MUTATION
} from "@/graphql/queries/orderQueries";
import ButtonStyled from "@/primatives/Button.vue";

@Component({
  components: {
    TableOrder,
    StartTableOrder,
    Category,
    MainMenu,
    ButtonStyled
  },
  apollo: {
    currentOrder: {
      query: GET_LATEST_ORDER_QUERY,
      variables() {
        return {
          table: this.currentTableId
        };
      },
      update: data => {
        if (data.table.orders.length > 0) {
          return data.table.orders[0];
        }
        return null;
      },
      skip() {
        return this.currentTableId === null;
      }
    }
  }
})
export default class InputOrder extends Vue {
  private currentTableId: number | null = null;
  private currentOrder: any = null;

  public async mounted() {
    const resp = await this.$apollo.query({
      query: GET_CURRENT_TABLE
    });

    this.currentTableId = resp.data.currentTable;

    if (!this.currentTableId) {
      this.$router.push({
        path: "/tables"
      });
      return;
    }
    console.log("refetch");
    this.$apollo.queries.currentOrder.refetch();
  }

  private async addProductToOrder(productId: number) {
    await this.$apollo.mutate({
      mutation: ADD_ITEMS_TO_ORDER_MUTATION,
      variables: {
        order: this.currentOrder.id,
        products: [productId]
      }
    });

    this.$apollo.queries.currentOrder.refetch();
  }

  private async startNewOrder() {
    await this.$apollo.mutate({
      mutation: START_NEW_ORDER_MUTATION,
      variables: {
        input: {
          table: this.currentTableId
        }
      }
    });

    this.$apollo.queries.currentOrder.refetch();
  }

  private async closeOrder() {
    await this.$apollo.mutate({
      mutation: CLOSE_ORDER_MUTATION,
      variables: {
        id: this.currentOrder.id
      }
    });

    this.$apollo.queries.currentOrder.refetch();
  }

  private changeTables() {
    this.$router.push({
      path: "/tables"
    });
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
