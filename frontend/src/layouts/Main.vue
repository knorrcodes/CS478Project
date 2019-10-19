<template>
  <div class="home">
    <header>
      <img class="logoImage" src="@/assets/biglogo.png" alt="big team logo" />
      <div>
        <button type="button" @click="logout" class="btn btn-secondary mx-1 my-1">Logout</button>
      </div>
    </header>

    <div class="container">
      <div class="row">
        <div class="col-4 px-4">
          <table-order />
        </div>

        <div class="col-8">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import TableOrder from "@/components/TableOrder.vue";

@Component({
  components: {
    TableOrder
  }
})
export default class Main extends Vue {
  private logout() {
    localStorage.removeItem("server-code");
    this.$router.push({ path: "/login" });
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
  margin-right: 5rem;
}
</style>
