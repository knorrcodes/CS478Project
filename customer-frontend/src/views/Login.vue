<template>
  <div>
    <div class="id1 container">
      <div class="row justify-content-center">
        <form action>
          <p class="text-center textFormat">Customer Login</p>
          <input
            id="serverCodeNumber"
            class="form-control input1 textFormat"
            v-model="serverCode"
            @keydown.enter.prevent="serverCodeCheck"
            type="text"
            v-focus
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
import { CHECK_FOR_CUST_CODE } from "@/graphql/queries/custCodeQueries";
import ButtonStyled from "@/primatives/Button.vue";

const NEXT_PAGE_URL = "/";

@Component({
  components: {
    ButtonStyled
  }
})
export default class LoginView extends Vue {
  private serverCode: number | null = null;
  private errorMsg = "";

  public beforeMount() {
    if (localStorage.getItem("customer-code")) {
      this.$router.push({
        path: NEXT_PAGE_URL
      });
    }
  }

  private async serverCodeCheck() {
    if (!this.serverCode) {
      return;
    }

    let resp;
    localStorage.setItem("customer-code", this.serverCode.toString());

    try {
      resp = await this.$apollo.query({
        query: CHECK_FOR_CUST_CODE,
        variables: {
          code: this.serverCode
        }
      });
    } catch (e) {
      localStorage.removeItem("customer-code");

      if (e.message.includes("401")) {
        this.errorMsg = "Customer code is not active";
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
  background: rgba(255, 255, 255, 0.63);
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
