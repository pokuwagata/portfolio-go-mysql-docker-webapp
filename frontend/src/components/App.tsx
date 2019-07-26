import * as React from 'react';
import { Route, Switch } from 'react-router-dom';

import { Header } from './Header';
import { Signup } from './Signup';
import { Login } from './Login';
import { ArticleList } from './ArticleList';

export interface AppProps {}

export const App = (props: AppProps) => (
  <div>
    <Header isLoggedIn={true}/>
    <Switch>
      <Route exact path="/" component={ArticleList} />
      <Route path="/signup" component={Signup} />
      <Route path="/login" component={Login} />
    </Switch>
  </div>
);
