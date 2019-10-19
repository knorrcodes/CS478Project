<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>

    <section>
      <section class="order-item" v-for="item in currentOrder.items" v-bind:key="item.id">
        <section
          class="order-item-product"
          v-for="product in item.products"
          v-bind:key="product.id"
        >
          <strong>{{ product.name }}</strong>
          <span>{{ formatPrice(product.price) }}</span>
        </section>
      </section>
    </section>

    <section class="order-total-section">
      <strong>Total</strong>
      <span class="total-cost">{{ totalCost }}</span>
    </section>
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";

@Component
export default class TableOrder extends Vue {
  @Prop() private readonly currentOrder: any;

  private formatPrice(cents: number): string {
    return `\$${(cents / 100).toFixed(2)}`;
  }

  private get totalCost(): string {
    const cost = this.currentOrder.items.reduce(
      (acc: number, item: any) =>
        acc +
        item.products.reduce(
          (acc: number, product: any) => acc + product.price,
          0
        ),
      0
    );

    return this.formatPrice(cost);
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

.order-item-product {
  display: flex;
  justify-content: space-between;
}

.order-total-section {
  border-top: dotted black 2px;
  display: flex;
  justify-content: space-between;
}

.total-cost {
  font-weight: bolder;
  text-align: right;
}
</style>
