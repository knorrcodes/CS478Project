<template>
  <div class="home">
    <header>
      <img class="logoImage" src="@/assets/biglogo.png" alt="big team logo" />
      <div class="logout-btn">
        <button-c :clickHandler="logout" value="Logout"></button-c>
      </div>
    </header>

    <div class="container">
      <div class="row">
        <div class="col-12">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import ButtonC from "@/primatives/Button.vue";

@Component({
  components: {
    ButtonC
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
}

.logout-btn {
  margin-right: 5rem;
}
</style>
