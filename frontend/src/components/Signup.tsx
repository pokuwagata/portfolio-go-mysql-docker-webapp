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
  const [username, setUsername] = React.useState('');
  const [usernameErrors, setUsernameErrors] = React.useState([]);
  const handleSubmit = (e: any) => {
    e.preventDefault();
    if (!validateForm()) return;

    // TODO: APIパス設定
    fetch('/api').then(
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

  const validateForm = (): boolean => {
    if (validateUserName()) {
      return true;
    } else {
      return false;
    }
  };

  const validateUserName = (): boolean => {
    const errors = checkUserNameError();
    setUsernameErrors(errors);

    return errors.length === 0;
  };

  // TODO: userName → username
  const checkUserNameError = (): Array<string> => {
    let errors = [];
    if (username.length === 0) {
      errors.push('ユーザ名を入力してください。');
      return errors;
    }
    if (!RegExp('^[A-Za-z0-9]+$').test(username)) 
      errors.push('半角英数字のみで入力してください。');
    
    if (username.length < 4)
      errors.push('ユーザ名は4文字以上で入力してください。');
    return errors;
  };

  return props.isLoggedIn ? (
    <Redirect to="/" />
  ) : (
    <div className="justify-content-center d-flex">
      <div className="mx-auto" style={{flex: '0 0 400px'}}>
        <h1 className="text-center">ユーザを登録する</h1>
        <form onSubmit={handleSubmit}>
          <div className="form-group row">
            <input
              type="text"
              className={
                'form-control' +
                (usernameErrors.length > 0 ? ' ' + 'is-invalid' : '')
              }
              placeholder="Username"
              maxLength={16}
              onChange={e => setUsername(e.target.value)}
            />
            {usernameErrors.length > 0 && (
              <div className="invalid-feedback">{usernameErrors.join('')}</div>
            )}
          </div>
          <div className="form-group row">
            <input
              type="password"
              className="form-control"
              placeholder="Password"
              minLength={8}
              maxLength={32}
              // pattern="[A-Za-z0-9]"
              // required
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
