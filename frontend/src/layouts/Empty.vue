<template>
  <div class="home">
    <header>
      <img class="logoImage" src="@/assets/biglogo.png" alt="big team logo" />
      <div class="logout-btn">
        <button-styled :clickHandler="logout" value="Logout"></button-styled>
      </div>
    </header>

    <div class="format">
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
import ButtonStyled from "@/primatives/Button.vue";

@Component({
  components: {
    ButtonStyled
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

.format {
  width: 100%;
  margin-left: auto;
  margin-right: auto;
}
</style>
