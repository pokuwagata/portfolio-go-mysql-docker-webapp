import * as React from 'react';
import { Redirect, RouteComponentProps } from 'react-router-dom';
import { FlushType } from './Flush';
import { FlushState } from './App';

export enum FormType {
  SIGNUP,
  LOGIN
}

export type FormDetail = {
  formType : FormType;
  requestUri : string;
  successMsg : string;
  errorMsg: string;
  titleText : string;
  buttonText : string;
}

export type UserFormProps = {
  formDetail: FormDetail;
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setFlushState: (state: FlushState) => void;
};

export const UserForm = (props: UserFormProps) => {
  const [username, setUsername] = React.useState('');
  const [usernameErrors, setUsernameErrors] = React.useState([]);
  const [password, setPassword] = React.useState('');
  const [passwordErrors, setPasswordErrors] = React.useState([]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // 少なくとも1つのフォームにバリデーションエラーが発生している場合は処理を中断
    if (!(validateUsername() && validatePassword())) return;

    fetch(props.formDetail.requestUri, {
      method : 'POST',
      headers: {'content-type': 'application/json'},
      body: JSON.stringify({
        username: username,
        password: password
      })
    })
    .then(res => {
      return new Promise((resolve) => res.json().then((json) => resolve({
        ok: res.ok,
        json
      })));
    })
    .then(res => {
        // TODO: 後で消す
        console.log(res);
        if((res as any).ok) {
          props.setFlushState({
            isDisplay: true,
            type: FlushType.SUCCESS,
            message: props.formDetail.successMsg,
          });
          props.setIsLoggedIn(true);
        } else {
          throw new Error((res as any).json.message)
        }
      }
    ).catch(
      (error)=>{
        // TODO: 後で消す
        props.setFlushState({
          isDisplay: true,
          type: FlushType.ERROR,
          message: props.formDetail.errorMsg
        });
      });
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

  return props.isLoggedIn ? (
    <Redirect to="/" />
  ) : (
    <div className="justify-content-center d-flex">
      <div className="mx-auto" style={{ flex: '0 0 400px' }}>
        <h1 className="text-center">{props.formDetail.titleText}</h1>
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
              {props.formDetail.buttonText}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}