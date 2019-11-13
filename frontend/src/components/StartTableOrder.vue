<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>

    <div>
      <button-styled :clickHandler="startOrder" value="Start Order"></button-styled>
    </div>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
import { START_NEW_ORDER_MUTATION } from "@/graphql/queries/orderQueries";
import ButtonStyled from "@/primatives/ButtonStyled.vue";

@Component({
  components: {
    ButtonStyled
  }
})
export default class TableOrder extends Vue {
  @Prop() private readonly startOrder: (() => void) | undefined;
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
