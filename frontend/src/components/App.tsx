import * as React from 'react';
import { Route, Switch } from 'react-router-dom';

import { Header } from './Header';
import { Signup } from './Signup';
import { Login } from './Login';
import { ArticleList } from './ArticleList';
import { FlushType } from './Flush';
import { ArticlePost } from './ArticlePost';
import { ArticleManagement } from './ArticleManagement';
import { FlushProvider } from './FlushProvider';
import { ArticleDetail } from './ArticleDetail';
import { PrivateRoute } from './PrivateRoute';
import * as Const from '../const';

type AppProps = {};
type AppState = {};
export type FlushState = {
  isDisplay: boolean;
  type?: FlushType;
  message?: string;
};

type getSessionResponse = {
  ok: boolean;
  loginUsername: string;
};

export const App = (props: AppProps) => {
  const [isLoggedIn, setIsLoggedIn] = React.useState(false);
  const [loginUsername, setLoginUsername] = React.useState();

  React.useEffect(() => {
    fetchSession();
  }, []); // マウント時のみ実行するため[]を渡す

  const fetchSession = async () => {
    const token = localStorage.getItem(Const.jwtTokenKey);
    if (!token) {
      setIsLoggedIn(false);
      setLoginUsername('');
    } else {
      try {
        const res = await fetch('api/admin/session', {
          method: 'GET',
          headers: { Authorization: 'Bearer' + ' ' + token },
        });
        const json = await res.json();
        if (res.ok) {
          setIsLoggedIn(true);
          setLoginUsername(json.username);
        } else {
          throw new Error();
        }
      } catch (error) {
        // TODO: エラーフラッシュの設定
        setIsLoggedIn(false);
        setLoginUsername('');
      }
    }
  };

  return (
    <div>
      <Header
        isLoggedIn={isLoggedIn}
        setIsLoggedIn={setIsLoggedIn}
        username={loginUsername}
      />
      <FlushProvider>
        <div className="container">
          <Switch>
            <Route exact path="/" component={ArticleList} />
            <Route
              path="/signup"
              render={props => (
                <Signup // TODO: 各コンポーネントのprops共通化
                  isLoggedIn={isLoggedIn}
                  setIsLoggedIn={setIsLoggedIn}
                  setLoginUsername={setLoginUsername}
                  {...props}
                />
              )}
            />
            <Route
              path="/login"
              render={props => (
                <Login
                  isLoggedIn={isLoggedIn}
                  setIsLoggedIn={setIsLoggedIn}
                  setLoginUsername={setLoginUsername}
                  {...props}
                />
              )}
            />
            <PrivateRoute path="/post" isLoggedIn={isLoggedIn}>
              <ArticlePost
                isLoggedIn={isLoggedIn}
                setIsLoggedIn={setIsLoggedIn}
                {...props}
              />
            </PrivateRoute>
            <PrivateRoute path="/management" isLoggedIn={isLoggedIn}>
              <ArticleManagement />
            </PrivateRoute>
            <Route path="/article" render={() => <ArticleDetail />} />
          </Switch>
        </div>
      </FlushProvider>
    </div>
  );
};
