<template>
  <div class="dspl">
    <button class="btn bg mx-1 my-1" type="button" @click="clickHandler">{{value}}</button>
  </div>
</template>

<script lang="ts">
import { Prop, Component, Vue } from "vue-property-decorator";

@Component
export default class ButtonStyled extends Vue {
  @Prop() private clickHandler: any;
  @Prop() private value: any;
}
</script>

<style scoped>
.bg {
  background-color: #3678a8;
  color: black;
}

.dspl {
  display: inline;
}
</style>
