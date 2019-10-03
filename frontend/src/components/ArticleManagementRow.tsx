import * as React from 'react';

export interface ArticleManagementRowProps {
  id: string;
  title: string;
  updatedAt: string;
}

export const ArticleManagementRow = (props: ArticleManagementRowProps) => {
  return (
    <div className="row mb-3 align-items-center">
      <div className="sm-col-auto mr-1">
        <input
          className="form-check-input position-static"
          type="checkbox"
          id={props.id}
        />
      </div>
      <div className="sm-col-6 mr-5">
        <label className="form-check-label" htmlFor={props.id}>
          {props.title}
        </label>
      </div>
      <div className="sm-col-auto">{new Date(props.updatedAt).toString()}</div>
      <div className="ml-md-auto">
        <button type="button" className="btn btn-primary">
          編集
        </button>
      </div>
    </div>
  );
};
