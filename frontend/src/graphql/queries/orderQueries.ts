import gql from "graphql-tag";

export const GET_LATEST_ORDER_QUERY = gql`
  query GetLatestOrder($table: ID!) {
    table(id: $table) {
      num
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

export const GET_LATEST_ORDER_ID_QUERY = gql`
  query GetLatestOrderId($table: ID!) {
    table(id: $table) {
      orders {
        id
      }
    }
  }
`;

export const START_NEW_ORDER_MUTATION = gql`
  mutation StartNewOrder($input: NewOrderInput!) {
    startOrder(input: $input) {
      id
    }
  }
`;

export const CLOSE_ORDER_MUTATION = gql`
  mutation CloseOrder($id: ID!) {
    closeOrder(id: $id) {
      id
      end_time
    }
  }
`;

export const ADD_ITEMS_TO_ORDER_MUTATION = gql`
  mutation AddItemsToOrder($order: ID!, $products: [Int!]!) {
    addItemToOrder(order: $order, products: $products) {
      id
    }
  }
`;

export const APPLY_PAYMENT = gql`
  mutation ApplyPayment($order: ID!, $amount: Int!) {
    applyPayment(input: {
      order: $order
      amount: $amount
    }){
      id
    }
  }
`;
