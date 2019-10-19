import gql from "graphql-tag";

export const GET_LATEST_ORDER_QUERY = gql`
  query GetLatestOrder($table: ID!) {
    table(id: $table) {
      orders {
        id
        payments {
          id
          amount
        }
        items {
          id
          products {
            id
            name
            price
            category {
              id
              name
            }
          }
        }
      }
    }
  }
`;
