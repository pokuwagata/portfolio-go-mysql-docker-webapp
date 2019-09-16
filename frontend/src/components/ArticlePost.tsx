import * as React from 'react';
import { Redirect } from 'react-router-dom';
import { FlushState } from './App';

export type ArticlePostProps = {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setFlushState: (state: FlushState) => void;
};

export const ArticlePost = (props: ArticlePostProps) => {
  const [title, setTitle] = React.useState('');
  const [titleErrors, setTitleErrors] = React.useState([]);
  const [content, setContent] = React.useState('');
  const [contentErrors, setContentErrors] = React.useState([]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // 少なくとも1つのフォームにバリデーションエラーが発生している場合は処理を中断
    const isValidTitle =  validateTitle();
    const isValidContent = validateContent();
    if(isValidTitle && isValidContent) return
  };

  const validateTitle = (): boolean => {
    const errors = checkTitleError();
    setTitleErrors(errors);

    return errors.length === 0;
  };

  const checkTitleError = (): Array<string> => {
    let errors = [];
    if (title.length === 0) {
      errors.push('タイトルを入力してください。');
    }
    return errors;
  }

  const validateContent = (): boolean => {
    const errors = checkContentError();
    setContentErrors(errors);

    return errors.length === 0;
  };

  const checkContentError = (): Array<string> => {
    let errors = [];
    if (content.length === 0) {
      errors.push('内容を入力してください。');
    }
    return errors;
  }

  return props.isLoggedIn ? (
    <div className="justify-content-center">
      <div>
        <h1>記事の投稿</h1>
        <form onSubmit={handleSubmit}>
          <div className="form-group row">
            <div className="col-8">
              <input
                type="text"
                className={
                  'form-control' +
                  (titleErrors.length > 0 ? ' ' + 'is-invalid' : '')
                }
                placeholder="タイトルを入力"
                maxLength={45}
                onChange={e => setTitle(e.target.value)}
              />
              {titleErrors.length > 0 && (
                <div className="invalid-feedback">{titleErrors.join('')}</div>
              )}
            </div>
          </div>
          <div className="form-group row">
            <div className="col-12">
              <textarea
                className={
                  'form-control' +
                  (contentErrors.length > 0 ? ' ' + 'is-invalid' : '')
                }
                placeholder="投稿したい内容を入力"
                maxLength={1000}
                onChange={e => setContent(e.target.value)}
                rows={15}
              />
              {contentErrors.length > 0 && (
                <div className="invalid-feedback">{contentErrors.join('')}</div>
              )}
            </div>
          </div>
          <div className="form-group row">
            <div className="col-2">
              <button type="submit" className="btn btn-primary">
                投稿
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  ) : (
    <Redirect to="/" />
  );
};
