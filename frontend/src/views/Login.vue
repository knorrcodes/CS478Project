<template>
  <div>
    <div class="id1 container">
      <div class="row justify-content-center">
        <form action>
          <p class="text-center textFormat">Server Login</p>
          <input
            id="serverCodeNumber"
            class="form-control input1 textFormat"
            v-model="serverCode"
            @keydown.enter.prevent="serverCodeCheck"
            type="number"
            autofocus
          />
          <div v-if="errorMsg">
            <h4 class="errorMessage">{{errorMsg}}</h4>
          </div>
          <br />
          <button-styled
            id="serverCodeButton"
            :clickHandler="() => serverCodeCheck()"
            value="Enter"
          ></button-styled>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import { GET_SERVER_QUERY } from "@/graphql/queries/serverQueries";
import ButtonStyled from "@/primatives/Button.vue";

const NEXT_PAGE_URL = "/tables";

@Component({
  components: {
    ButtonStyled
  }
})
export default class LoginView extends Vue {
  private serverCode: number = 0;
  private errorMsg = "";

  public beforeMount() {
    if (localStorage.getItem("server-code")) {
      this.$router.push({
        path: NEXT_PAGE_URL
      });
    }
  }

  private async serverCodeCheck() {
    let resp;
    localStorage.setItem("server-code", this.serverCode.toString());

    try {
      resp = await this.$apollo.query({
        query: GET_SERVER_QUERY,
        variables: {
          code: this.serverCode
        }
      });
    } catch (e) {
      localStorage.removeItem("server-code");

      if (e.message.includes("401")) {
        this.errorMsg = "Invalid server code";
      } else {
        this.errorMsg = "An error occured";
      }
      return;
    }

    this.$router.push({ path: NEXT_PAGE_URL });
  }
}
</script>

<style scoped>
.id1 {
  text-align: center;
  justify-content: center;
}
.input1 {
  -webkit-appearance: none;
  background: rgba(255, 255, 255, 0);
  text-align: center;
  border: none;
  display: block;
  margin-left: auto;
  margin-right: auto;
}

.textFormat {
  font-family: "Gill Sans", "Gill Sans MT", Calibri, "Trebuchet MS", sans-serif;
  font-size: 18pt;
}

input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  margin: 0;
}

.errorMessage {
  color: red;
  text-transform: uppercase;
}
</style>
