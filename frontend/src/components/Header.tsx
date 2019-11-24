import * as React from 'react';
import { HeaderButton } from './HeaderButton';
import { SignoutButton } from './SignoutButton';

export interface HeaderProps {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  username: string;
}

export const Header = (props: HeaderProps) => (
  <nav className="navbar navbar-expand-sm navbar-light bg-light mb-3">
    <div className="navbar-collapse">
      <ul className="navbar-nav mr-auto">
        <HeaderButton name="Home" path="/" />
        {props.isLoggedIn ? <li className="navbar-text" style={{marginLeft: "1rem"}}>ユーザ:{props.username}</li> : null}
      </ul>
      <ul className="navbar-nav">
        {props.isLoggedIn ? (
          <>
            <HeaderButton name="投稿する" path="/post" />
            <HeaderButton name="管理" path="/management" />
            <SignoutButton setIsLoggedIn={props.setIsLoggedIn}></SignoutButton>
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
