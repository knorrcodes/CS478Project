import gql from "graphql-tag";

export const GET_SERVER_QUERY = gql`
  query GetServer($code: Int!) {
    server(code: $code) {
      id
      name
    }
  }
`;
