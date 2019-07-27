import * as React from 'react';
import { Redirect, RouteComponentProps } from 'react-router-dom';
import { FlushType } from './Flush';

export type SignupProps = RouteComponentProps & {
  // TODO: 後で型定義
  isLoggedIn: boolean;
  setIsLoggedIn: any;
  setFlushState: any;
};

export const Signup = (props: SignupProps) => {
  const [username, setUserName] = React.useState('');
  const handleSubmit = (e: any) => {
    e.preventDefault();
    // TODO: APIパス設定
    fetch('/').then(
      res => {
        // TODO: 後で消す
        console.log(res);
        props.setFlushState({
          isDisplay: true,
          type: FlushType.SUCCESS,
          message: 'ユーザの登録に成功しました',
        });
        props.setIsLoggedIn(true);
      },
      error => {
        // TODO: 後で消す
        console.log(error);
        props.setFlushState({
          isDisplayFlush: true,
          flushType: FlushType.ERROR,
        });
      }
    );
  };

  return props.isLoggedIn ? (
    <Redirect to="/" />
  ) : (
    <div className="justify-content-center d-flex">
      <div className="mx-auto">
        <h1>ユーザを登録する</h1>
        <form onSubmit={handleSubmit}>
          <div className="form-group row">
            <input
              type="text"
              className="form-control"
              placeholder="Username"
              minLength={4}
              maxLength={16}
              // TODO: patternの挙動は後で調査
              // pattern="[A-Za-z0-9]"
              required
              onChange={e => setUserName(e.target.value)}
            />
          </div>
          <div className="form-group row">
            <input
              type="password"
              className="form-control"
              placeholder="Password"
              minLength={8}
              maxLength={32}
              // pattern="[A-Za-z0-9]"
              required
            />
          </div>
          <div className="form-group row justify-content-center">
            <button type="submit" className="btn btn-primary">
              登録
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};
