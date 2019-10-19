<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>

    <span v-if="!current_table_id">No open order for table</span>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
import { GET_CURRENT_TABLE } from "@/graphql/queries/tableQueries";
import { GET_LATEST_ORDER_QUERY } from "@/graphql/queries/orderQueries";

@Component({
  apollo: {
    current_order: {
      query: GET_LATEST_ORDER_QUERY,
      variables() {
        return {
          table: this.current_table_id
        };
      },
      update: data => {
        if (data.table.orders.length > 0) {
          return data.table.orders[0];
        }
        return null;
      },
      pollInterval: 500
    }
  }
})
export default class TableOrder extends Vue {
  private current_table_id: number | null = null;

  async mounted() {
    const resp = await this.$apollo.query({
      query: GET_CURRENT_TABLE
    });

    this.current_table_id = resp.data.currentTable;
    console.log(this.current_table_id);
    if (!this.current_table_id) {
      this.$apollo.queries.current_order.stop();
    }
  }
}
</script>

<style scoped>
h3 {
  text-align: center;
}

.ticketOrder {
  border-radius: 25px;
  border: 2px solid;
}
</style>
