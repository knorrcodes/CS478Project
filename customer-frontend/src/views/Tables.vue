<template>
  <div>
    <div class="container text-center">
      <h1>Tables</h1>
      <div class="row">
        <div class="col" align="center">
          <button-styled
            v-for="table in tables"
            v-bind:key="table.id"
            :clickHandler="() => setCurrentTable(table.id)"
            :value="table.num"
          ></button-styled>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
import {
  GET_ALL_TABLES_QUERY,
  SET_CURRENT_TABLE
} from "@/graphql/queries/tableQueries";
import ButtonStyled from "@/primatives/Button.vue";

@Component({
  apollo: {
    tables: GET_ALL_TABLES_QUERY
  },
  components: {
    ButtonStyled
  }
})
export default class TableView extends Vue {
  private tables: any = null;

  private async setCurrentTable(id: number) {
    await this.$apollo.mutate({
      mutation: SET_CURRENT_TABLE,
      variables: {
        id
      }
    });

    this.$router.push({
      path: "/"
    });
  }
}
</script>

<style scoped>
</style>
