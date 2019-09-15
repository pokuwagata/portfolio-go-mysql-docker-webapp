import * as React from 'react';
import { Redirect, RouteComponentProps } from 'react-router-dom';
import { FlushType } from './Flush';
import { FlushState } from './App';

export type SignupProps = RouteComponentProps & {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setFlushState: (state: FlushState) => void;
};

export const Signup = (props: SignupProps) => {
  const [username, setUsername] = React.useState('');
  const [usernameErrors, setUsernameErrors] = React.useState([]);
  const [password, setPassword] = React.useState('');
  const [passwordErrors, setPasswordErrors] = React.useState([]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // 少なくとも1つのフォームにバリデーションエラーが発生している場合は処理を中断
    if (!(validateUsername() && validatePassword())) return;

    // TODO: APIパス設定
    fetch('/api/user', {
      method : 'POST',
      headers: {'content-type': 'application/json'},
      body: JSON.stringify({
        username: username,
        password: password
      })
    }).then(
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
          isDisplay: true,
          type: FlushType.ERROR,
          message: 'error'
        });
      }
    );
  };

  const validateUsername = (): boolean => {
    const errors = checkUsernameError();
    setUsernameErrors(errors);

    return errors.length === 0;
  };

  const checkUsernameError = (): Array<string> => {
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

  const validatePassword = (): boolean => {
    const errors = checkPasswordError();
    setPasswordErrors(errors);

    return errors.length === 0;
  };

  const checkPasswordError = (): Array<string> => {
    let errors = [];
    if (password.length === 0) {
      errors.push('パスワードを入力してください。');
      return errors;
    }
    if (!RegExp('^[A-Za-z0-9]+$').test(password))
      errors.push('半角英数字のみで入力してください。');

    if (password.length < 8)
      errors.push('パスワードは8文字以上で入力してください。');
    return errors;
  };

  //TODO: ユーザ名とパスワードのフォームは別コンポーネントとして切り出し検討
  return props.isLoggedIn ? (
    <Redirect to="/" />
  ) : (
    <div className="justify-content-center d-flex">
      <div className="mx-auto" style={{ flex: '0 0 400px' }}>
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
              className={
                'form-control' +
                (passwordErrors.length > 0 ? ' ' + 'is-invalid' : '')
              }
              placeholder="Password"
              maxLength={32}
              onChange={e => setPassword(e.target.value)}
            />
            {passwordErrors.length > 0 && (
              <div className="invalid-feedback">{passwordErrors.join('')}</div>
            )}
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
