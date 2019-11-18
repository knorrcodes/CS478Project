<template>
  <div class="home">
    <header>
      <img class="logoImage" src="@/assets/biglogo.png" alt="big team logo" />
      <div class="logout-btn">
        <button-styled :clickHandler="logout" value="Logout"></button-styled>
      </div>
    </header>

    <div class="container format">
      <div class="row">
        <div class="col-12">
          <h1>Customer Mode</h1>
          <hr class="mb-4" />
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import ButtonStyled from "@/primatives/Button.vue";

@Component({
  components: {
    ButtonStyled
  }
})
export default class Main extends Vue {
  private logout() {
    localStorage.removeItem("customer-code");
    this.$router.push({ path: "/login" });
  }
}
</script>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgb(170, 170, 170);
  margin-bottom: 2rem;
}

.logoImage {
  margin-left: 2rem;
  height: 75px;
}

.logout-btn {
  margin-right: 2rem;
}

.format {
  width: 100%;
  margin-left: auto;
  margin-right: auto;
}
hr {
  width: 80%;
}

h1 {
  text-align: center;
  font-weight: bold;
}
</style>
