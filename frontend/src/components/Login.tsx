import * as React from 'react';
import { RouteComponentProps } from 'react-router-dom';
import { FlushState } from './App';
import { UserForm, FormType } from './UserForm';

// TODO: RouteComponentPropsは必要？
export type LoginProps = RouteComponentProps & {
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
