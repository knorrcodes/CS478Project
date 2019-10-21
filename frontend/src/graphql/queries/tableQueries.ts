import gql from "graphql-tag";

export const GET_CURRENT_TABLE = gql`
  query GetCurrentTable {
    currentTable @client
  }
`;

export const SET_CURRENT_TABLE = gql`
  mutation SetCurrentTable($id: Int) {
    setCurrentTable(id: $id) @client
  }
`;

export const GET_ALL_TABLES_QUERY = gql`
  query GetAllTables {
    tables {
      id
      num
    }
  }
`;
