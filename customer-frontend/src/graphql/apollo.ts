import Vue from "vue";
import VueApollo from "vue-apollo";
import { ApolloClient } from "apollo-client";
import { HttpLink } from "apollo-link-http";
import { setContext } from "apollo-link-context";
import { onError } from "apollo-link-error";
import { InMemoryCache } from "apollo-cache-inmemory";
import { resolvers, typeDefs } from "./localResolvers";

// HTTP link to graphql API
const httpLink = new HttpLink({
  uri: "/graphql"
});

// Set auth JWT on each request if available
const authLink = setContext((_, { headers }) => {
  const code = localStorage.getItem("customer-code");

  return {
    headers: {
      ...headers,
      authorization: code ? `Customer ${code}` : ""
    }
  };
});

// Error Handling
const errorLink = onError(({ graphQLErrors, networkError }) => {
  if (graphQLErrors) {
    graphQLErrors.map(({ message, locations, path }) =>
      console.log(
        `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
      )
    );
  }

  if (networkError) {
    console.log(`[Network error]: ${networkError}`);
  }
});

// Prepare memory cache
const cache = new InMemoryCache();

// Create the apollo client
export const apolloClient = new ApolloClient({
  link: errorLink.concat(authLink.concat(httpLink)),
  connectToDevTools: true,
  cache,
  resolvers,
  typeDefs
});

cache.writeData({
  data: {
    currentTable: null
  }
});

// Install the Vue plugin
Vue.use(VueApollo);

export const apolloProvider = new VueApollo({
  defaultClient: apolloClient
});
