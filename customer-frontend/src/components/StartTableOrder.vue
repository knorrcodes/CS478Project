<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>

    <div v-if="cust_code == null">
      <p>Please ask your server to generate a code for you</p>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
import { START_NEW_ORDER_MUTATION } from "@/graphql/queries/orderQueries";
import { CHECK_FOR_CUST_CODE } from "@/graphql/queries/custCodeQueries";
import ButtonStyled from "@/primatives/Button.vue";

@Component
export default class TableOrder extends Vue {
  @Prop() private readonly startOrder: any;
  cust_code = localStorage.getItem("customer-code");
}
</script>

<style scoped>
h3 {
  text-align: center;
}

.ticketOrder {
  padding: 1rem 2rem;
  border-radius: 10px;
  background: white;
}
</style>
