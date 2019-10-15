<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>
    <dl>
      <dt>main entree</dt>
      <dd>side1</dd>
      <dd>side2</dd>
      <dd>drink</dd>
    </dl>
    <h4>test data</h4>
    <h4>test data</h4>
    <h4>test data</h4>
    <h4>test data</h4>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";

@Component
export default class tableOrder extends Vue {}
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