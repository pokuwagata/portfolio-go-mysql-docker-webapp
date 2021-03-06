import * as React from 'react';
import { Redirect } from 'react-router-dom';
import { FlushType } from './Flush';
import { FlushActionType, FlushDispatchContext } from './FlushProvider';
import * as Const from '../const';
import gql from 'graphql-tag';
import { useMutation } from 'react-apollo';

export enum FormType {
  SIGNUP,
  LOGIN,
}

export type FormDetail = {
  formType: FormType;
  requestUri: string;
  successMsg: string;
  titleText: string;
  buttonText: string;
};

export type UserFormProps = {
  formDetail: FormDetail;
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setLoginUsername: (state: boolean) => void;
};

export const UserForm = (props: UserFormProps) => {
  const [username, setUsername] = React.useState('');
  const [usernameErrors, setUsernameErrors] = React.useState([]);
  const [password, setPassword] = React.useState('');
  const [passwordErrors, setPasswordErrors] = React.useState([]);
  const flushDispatch = React.useContext(FlushDispatchContext);

  const CREATE_SESSION = gql`
    mutation createSession($username: String, $password: String) {
      createSession(input: { username: $username, password: $password }) {
        token
        username
      }
    }
  `;
  const [createSession] = useMutation(CREATE_SESSION, {
    ignoreResults: false,
    onCompleted: data => {
        flushDispatch({
          type: FlushActionType.VISIBLE,
          payload: {
            type: FlushType.SUCCESS,
            message: props.formDetail.successMsg,
          },
        });
        localStorage.setItem(Const.jwtTokenKey, data.createSession.token);
        props.setIsLoggedIn(true);
        props.setLoginUsername(data.createSession.username);
    },
    onError: err => {
      flushDispatch({
          type: FlushActionType.VISIBLE,
          payload: {
            type: FlushType.ERROR,
            message: err.message,
          },
        });
    },
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // 少なくとも1つのフォームにバリデーションエラーが発生している場合は処理を中断
    const isValidUsername = validateUsername();
    const isValidPassword = validatePassword();
    if (!(isValidUsername && isValidPassword)) return;

    createSession({ variables: { username: username, password: password } });
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
    <Redirect to="/management" />
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
              placeholder="ユーザ名を入力"
              maxLength={16}
              onChange={e => setUsername(e.target.value)}
              autoFocus
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
              placeholder="パスワードを入力"
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
};
