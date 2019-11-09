import gql from "graphql-tag";

export const GENERATE_CUST_CODE = gql`
  mutation CreateCustCode($order: ID!) {
    createCustCode(id: $order) {
      id
    }
  }
`;
