<template>
  <div class="ticketOrder">
    <h3>Table Order</h3>

    <section>
      <section v-for="item in currentOrder.items" v-bind:key="item.id">
        <section class="order-item-product">
          <strong>{{ item.products[0].name }}</strong>
          <span>{{ formatPrice(item.products[0].price) }}</span>
        </section>

        <section
          class="order-item-sub-product"
          v-for="product in item.products.slice(1)"
          v-bind:key="product.id"
        >
          <span>{{ product.name }}</span>
        </section>
      </section>
    </section>

    <section class="order-total-section" v-if="currentOrder.payments.length > 0">
      <div v-for="payment in currentOrder.payments" v-bind:key="payment.id">
        <strong>Payment</strong>
        <span>{{ formatPrice(-payment.amount) }}</span>
      </div>
    </section>

    <section class="order-total-section">
      <div>
        <strong>Subtotal</strong>
        <span class="total-cost">{{ formatPrice(subTotal) }}</span>
      </div>
      <div>
        <strong>Tax</strong>
        <span class="total-cost">{{ formatPrice(taxAmount) }}</span>
      </div>
      <div>
        <strong>Total</strong>
        <span class="total-cost">{{ formatPrice(totalAmount) }}</span>
      </div>
    </section>

    <section class="order-total-section">
      <div>
        <strong>Remaining</strong>
        <span class="total-cost">{{ formatPrice(totalAmount - appliedPayments) }}</span>
      </div>
    </section>

    <button-styled value="Make Payment" :clickHandler="() => setDialogState(true)" />

    <dialog-box
      v-if="dialogIsOpen"
      prompt="Payment Amount"
      :okHandler="makePayment"
      :cancelHandler="() => setDialogState(false)" />
  </div>
</template>

<script lang="ts">
import { Vue, Prop, Component } from "vue-property-decorator";
import ButtonStyled from "@/primatives/Button.vue";
import Dialog from "@/components/Dialog.vue";
import { APPLY_PAYMENT } from "@/graphql/queries/orderQueries";

@Component({
  components: {
    ButtonStyled,
    DialogBox: Dialog
  }
})
export default class TableOrder extends Vue {
  @Prop() private readonly currentOrder: any;
  @Prop() private readonly refetchFunc: any;

  private dialogIsOpen = false;

  private async makePayment(amount) {
    this.setDialogState(false);
    
    await this.$apollo.mutate({
      mutation: APPLY_PAYMENT,
      variables: {
        order: this.currentOrder.id,
        amount: amount * 100
      }
    });

    this.refetchFunc();
  }

  private setDialogState(open: boolean){
    this.dialogIsOpen = open;
  }

  private formatPrice(cents: number): string {
    return `\$${(cents / 100).toFixed(2)}`;
  }

  private get subTotal(): number {
    return this.currentOrder.items.reduce(
      (acc: number, item: any) => acc + item.products[0].price,
      0
    );
  }

  private get taxAmount(): number {
    return this.subTotal * 0.07;
  }

  private get totalAmount(): number {
    return this.subTotal + this.taxAmount;
  }

  private get appliedPayments(): number {
    return this.currentOrder.payments.reduce(
      (acc: number, payment: any) => acc + payment.amount, 0
    );
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
  background-color: white;
}

.order-item-product {
  display: flex;
  justify-content: space-between;
}

.order-item-sub-product {
  display: flex;
  justify-content: space-between;
  margin-left: 10px;
  /* font-weight: normal; */
}

.order-total-section {
  border-top: dotted black 2px;
}

.order-total-section > div {
  display: flex;
  justify-content: space-between;
}

.total-cost {
  font-weight: bolder;
  text-align: right;
}
</style>
