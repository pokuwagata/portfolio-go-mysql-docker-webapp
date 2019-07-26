import * as React from 'react';
import { HeaderButton } from './HeaderButton';

export interface HeaderProps {}

export const Header = (props: HeaderProps) => (
  <div>
    <HeaderButton name="Home" path="/" />
    <HeaderButton name="ユーザ登録" path="/signup" />
    <HeaderButton name="ログイン" path="/login" />
  </div>
);
