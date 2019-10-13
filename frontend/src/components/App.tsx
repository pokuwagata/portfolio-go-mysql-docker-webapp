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
  const [flushState, setFlushState] = React.useState<FlushState>({
    isDisplay: false,
    type: undefined,
    message: '',
  });
  const [loginUsername, setLoginUsername] = React.useState();

  React.useEffect(() => {
    const token = localStorage.getItem('portfolio-jwt-token');
    if (!token) {
      setIsLoggedIn(false);
      setLoginUsername('');
    } else {
      fetch('api/admin/session', {
        method: 'GET',
        headers: { Authorization: 'Bearer' + ' ' + token },
      })
        .then(res => {
          return new Promise(resolve =>
            res.json().then(json =>
              resolve({
                ok: res.ok,
                json,
              })
            )
          );
        })
        .then(res => {
          if ((res as getSessionResponse).ok) {
            setIsLoggedIn(true);
            setLoginUsername((res as getSessionResponse).loginUsername);
          } else {
            throw new Error();
          }
        })
        .catch(error => {
          // TODO: エラーフラッシュの設定
          setIsLoggedIn(false);
          setLoginUsername('');
        });
    }
  }, []); // マウント時のみ実行するため[]を渡す

  return (
    <div>
      <Header isLoggedIn={isLoggedIn} />
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
            <Route
              path="/post"
              render={props => (
                <ArticlePost
                  isLoggedIn={isLoggedIn}
                  setIsLoggedIn={setIsLoggedIn}
                  {...props}
                />
              )}
            />
            <Route
              path="/management"
              render={props => (
                <ArticleManagement
                  isLoggedIn={isLoggedIn}
                  setIsLoggedIn={setIsLoggedIn}
                  {...props} // TODO: 不要な場合は削除する
                />
              )}
            />
            <Route
              path="/article"
              render={() => <ArticleDetail isLoggedIn={isLoggedIn} />}
            />
          </Switch>
        </div>
      </FlushProvider>
    </div>
  );
};
