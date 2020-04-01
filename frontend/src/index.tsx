import ApolloClient from 'apollo-boost';
import 'bootstrap/dist/css/bootstrap.min.css';
import * as React from 'react';
import { ApolloProvider } from 'react-apollo';
import * as ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import { App } from './components/App';
import { ScrollToTop } from './components/ScrollToTop';

const client = new ApolloClient({ uri: '/bff/graphql' });

ReactDOM.render(
  <ApolloProvider client={client}>
    <BrowserRouter>
      <ScrollToTop />
      <App />
    </BrowserRouter>
  </ApolloProvider>,
  document.getElementById('app-root')
);
