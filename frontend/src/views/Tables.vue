<template>
  <div>
    <div class="container text-center">
      <h1>Tables</h1>
      <div class="row">
        <div class="col" align="center">
          <button
            v-for="table in tables"
            v-bind:key="table.id"
            class="btn btn-secondary btn-lg mx-1 my-1"
            @click="setCurrentTable(table.id)"
          >{{ table.num }}</button>
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

@Component({
  apollo: {
    tables: GET_ALL_TABLES_QUERY
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
