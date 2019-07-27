import * as React from 'react';
import { HeaderButton } from './HeaderButton';

export interface HeaderProps {
  isLoggedIn: boolean;
}

export const Header = (props: HeaderProps) => (
  <nav className="navbar navbar-expand-lg navbar-light bg-light">
    <div className="navbar-collapse">
      <ul className="navbar-nav mr-auto">
        <HeaderButton name="Home" path="/" />
      </ul>
      <ul className="navbar-nav">
        {props.isLoggedIn ? (
          <>
            <HeaderButton name="投稿する" path="/post" />
            <HeaderButton name="管理" path="/management" />
          </>
        ) : (
          <>
            <HeaderButton name="ユーザ登録" path="/signup" />
            <HeaderButton name="ログイン" path="/login" />
          </>
        )}
      </ul>
    </div>
  </nav>
);
