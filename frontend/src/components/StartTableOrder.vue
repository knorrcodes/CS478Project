<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>

    <div>
      <button type="button" class="btn btn-secondary mx-1 my-1" @click="startNewOrder">Start Order</button>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
import { START_NEW_ORDER_MUTATION } from "@/graphql/queries/orderQueries";

@Component
export default class TableOrder extends Vue {
  @Prop() private readonly tableId: any;

  public async startNewOrder() {
    await this.$apollo.mutate({
      mutation: START_NEW_ORDER_MUTATION,
      variables: {
        input: {
          table: this.tableId
        }
      }
    });
  }
}
</script>

<style scoped>
h3 {
  text-align: center;
}

.ticketOrder {
  padding: 1rem 2rem;
  border-radius: 10px;
  border: 2px solid;
}
</style>
