import * as React from 'react';

export interface SignupProps {}

export const Signup = (props: SignupProps) => (
    <div className="justify-content-center d-flex">
        <div className="mx-auto">
            <h1>ユーザを登録する</h1>
            <form>
                <div className="form-group row">
                    <input
                        type="text"
                        className="form-control"
                        placeholder="Username"
                    />
                </div>
                <div className="form-group row">
                    <input
                        type="password"
                        className="form-control"
                        placeholder="Password"
                    />
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
