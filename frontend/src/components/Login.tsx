import * as React from 'react';
import { UserForm, FormType } from './UserForm';

export type LoginProps = {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setLoginUsername: (state: boolean) => void;
};

export const Login = (props: LoginProps) => {
  return <UserForm formDetail={{
    formType : FormType.LOGIN,
    requestUri: 'api/session',
    successMsg : 'ログインに成功しました',
    titleText : 'ログインする',
    buttonText : 'ログイン'
  }} {...props}></UserForm>
}
