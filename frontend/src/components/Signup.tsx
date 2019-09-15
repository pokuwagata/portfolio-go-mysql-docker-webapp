import * as React from 'react';
import { RouteComponentProps } from 'react-router-dom';
import { FlushState } from './App';
import { UserForm, FormType } from './UserForm';

// TODO: RouteComponentPropsは必要？
export type SignupProps = RouteComponentProps & {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setFlushState: (state: FlushState) => void;
};

export const Signup = (props: SignupProps) => {
  return <UserForm formDetail={{
    formType : FormType.SIGNUP,
    requestUri: 'api/user',
    successMsg : 'ユーザの登録に成功しました',
    errorMsg: 'エラーが発生しました。管理者にお問い合わせください',
    titleText : 'ユーザを登録する',
    buttonText : '登録'
  }} {...props}></UserForm>
}
