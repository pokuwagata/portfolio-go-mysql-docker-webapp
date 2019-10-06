import * as React from 'react';

export interface ArticleManagementRowProps {
  id: string;
  title: string;
  updatedAt: string;
  checked: boolean;
  setSelected: (id: string) => void;
}

export const ArticleManagementRow = (props: ArticleManagementRowProps) => {
  const onClick = (event: React.ChangeEvent<HTMLInputElement>, id: string) => {
    if (event.target.checked) {
      props.setSelected(id);
    } else {
      props.setSelected(null);
    }
  };

  return (
    <div className="row mb-3 align-items-center">
      <div className="sm-col-auto mr-1">
        <input
          className="form-check-input position-static"
          type="checkbox"
          id={props.id}
          onChange={event => onClick(event, props.id)}
          checked={props.checked}
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
