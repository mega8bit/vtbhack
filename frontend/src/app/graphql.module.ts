import {NgModule} from '@angular/core';
import {ApolloModule, APOLLO_OPTIONS, Apollo} from 'apollo-angular';
import {HttpLinkModule, HttpLink} from 'apollo-angular-link-http';
import {InMemoryCache} from 'apollo-cache-inmemory';
import {ApolloLink} from "apollo-link";

const authLink = new ApolloLink((operation, forward) => {
  // Retrieve the authorization token from local storage.
  const token = localStorage.getItem('auth_token');

  // Use the setContext method to set the HTTP headers.
  operation.setContext({
    headers: {
      Authorization: token ? `Bearer ${token}` : ''
    }
  });

  // Call the next link in the middleware chain.
  return forward(operation);
});

// const uri = 'api'; // <-- add the URL of the GraphQL server here
// export function createApollo(httpLink: HttpLink) {
//   return {
//     link: authLink.concat(httpLink.create({uri})),
//     cache: new InMemoryCache(),
//   };
// }

@NgModule({
  exports: [ApolloModule, HttpLinkModule],
})
export class GraphQLModule {
  private readonly URI1: string = '/api';
  private readonly URI2: string = '/chat';

  constructor(
    apollo: Apollo,
    httpLink: HttpLink
  ) {
    const options1: any = { uri: this.URI1 };
    apollo.createDefault({
      link: authLink.concat(httpLink.create(options1)),
      cache: new InMemoryCache()
    });

    const options2: any = { uri: this.URI2 };
    apollo.createNamed('chat', {
      link: httpLink.create(options2),
      cache: new InMemoryCache()
    });
  }
}
