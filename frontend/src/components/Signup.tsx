import * as React from 'react';
import { UserForm, FormType } from './UserForm';

export type SignupProps = {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setLoginUsername: (state: boolean) => void;
};

export const Signup = (props: SignupProps) => {
  return <UserForm formDetail={{
    formType : FormType.SIGNUP,
    requestUri: 'api/user',
    successMsg : 'ユーザの登録に成功しました',
    titleText : 'ユーザを登録する',
    buttonText : '登録'
  }} {...props}></UserForm>
}
