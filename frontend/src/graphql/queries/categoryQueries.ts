import gql from "graphql-tag";

export const GET_ALL_CATEGORIES_QUERIES = gql`
  query GetAllCategories {
    categories {
      id
      name
    }
  }
`;

export const GET_PRODUCTS_IN_CATEGORY_QUERIES = gql`
  query GetProductsInCategory($id: ID!) {
    category(id: $id) {
      id
      name

      products {
        id
        name
      }
    }
  }
`;
