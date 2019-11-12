import gql from "graphql-tag";

export const GENERATE_CUST_CODE = gql`
  mutation CreateCustCode($order: ID!) {
    createCustCode(id: $order) {
      id
    }
  }
`;

export const CHECK_FOR_CUST_CODE = gql`
  query orders{
    cust_code{
      code
    }
  }
`;
