import gql from "graphql-tag";
import { GET_CURRENT_TABLE } from "./queries/tableQueries";

export const typeDefs = gql`
  extend type Query {
    currentTable: Int
  }
`;

export const resolvers = {
  Mutation: {
    setCurrentTable: (_: any, { id }: any, { cache }: any) => {
      cache.writeQuery({
        query: GET_CURRENT_TABLE,
        data: {
          currentTable: id
        }
      });
    }
  }
};
