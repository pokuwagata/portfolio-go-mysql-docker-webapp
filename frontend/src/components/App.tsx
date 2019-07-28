import * as React from 'react';
import { Route, Switch } from 'react-router-dom';

import { Header } from './Header';
import { Signup } from './Signup';
import { Login } from './Login';
import { ArticleList } from './ArticleList';
import { Flush, FlushType } from './Flush';

type AppProps = {};
type AppState = {};
export type FlushState = {
  isDisplay: boolean;
  type: FlushType;
  message: string;
};

export const App = (props: AppProps) => {
  const [isLoggedIn, setIsLoggedIn] = React.useState(false);
  const [flushState, setFlushState] = React.useState({
    isDisplay: false,
    type: undefined,
    message: '',
  });

  return (
    <div>
      <Header isLoggedIn={false} />
      <Flush {...flushState} setFlushState={setFlushState} />
      <div className="container">
        <Switch>
          <Route exact path="/" component={ArticleList} />
          <Route
            path="/signup"
            render={props => (
              <Signup
                isLoggedIn={isLoggedIn}
                setIsLoggedIn={setIsLoggedIn}
                setFlushState={setFlushState}
                {...props}
              />
            )}
          />
          <Route path="/login" component={Login} />
        </Switch>
      </div>
    </div>
  );
};
