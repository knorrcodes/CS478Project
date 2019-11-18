import gql from "graphql-tag";

export const GENERATE_CUST_CODE = gql`
  mutation CreateCustCode($order: ID!) {
    createCustCode(id: $order) {
      id
    }
  }
`;

export const CHECK_FOR_CUST_CODE = gql`
  query CheckCustCode($code: String!) {
    custcode(code: $code) {
      id
    }
  }
`;

export const GET_CURRENT_CUST_ORDER = gql`
  query CheckCustCode($code: String!) {
    custcode(code: $code) {
      id
      order {
        id
        table {
          id
        }
        items {
          products {
            id
            name
            price
          }
        }
        payments {
          id
          amount
        }
        cust_code {
          id
          code
        }
      }
    }
  }
`;
