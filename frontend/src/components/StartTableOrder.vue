<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>

    <div>
      <button-c :clickHandler="() => startOrder()" value="Start Order"></button-c>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
import { START_NEW_ORDER_MUTATION } from "@/graphql/queries/orderQueries";
import ButtonC from "@/primatives/Button.vue";

@Component({
  components: {
    ButtonC
  }
})
export default class TableOrder extends Vue {
  @Prop() private readonly startOrder: any;
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
