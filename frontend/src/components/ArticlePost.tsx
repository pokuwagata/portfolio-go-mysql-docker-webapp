import * as React from 'react';
import { Redirect, useLocation } from 'react-router-dom';
import { FlushType } from './Flush';
import { FlushDispatchContext, FlushActionType } from './FlushProvider';

export type ArticlePostProps = {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
};

export const ArticlePost = (props: ArticlePostProps) => {
  const id = new URLSearchParams(useLocation().search).get('id');

  const [loading, setLoading] = React.useState(true);
  const [title, setTitle] = React.useState('');
  const [titleErrors, setTitleErrors] = React.useState([]);
  const [content, setContent] = React.useState('');
  const [contentErrors, setContentErrors] = React.useState([]);
  const [postDone, setPostDone] = React.useState(false);

  const flushDispath = React.useContext(FlushDispatchContext);

  const fetchArticle = async (id: string) => {
    try {
      const res = await fetch('api/articles/' + id, {
        method: 'GET',
      });
      const json = await res.json();
      if (res.ok) {
        flushDispath({
          type: FlushActionType.HIDDEN,
        });
        setTitle(json.title);
        setContent(json.content);
        setLoading(false);
      } else {
        throw new Error(json.message);
      }
    } catch (error) {
      flushDispath({
        type: FlushActionType.VISIBLE,
        payload: {
          type: FlushType.ERROR,
          message: '記事の取得に失敗しました。' + error,
        },
      });
      setLoading(false);
    }
  };

  React.useEffect(() => {
    if (id) {
      // 記事のIdをURLから取得できる場合は編集
      fetchArticle(id);
    } else {
      // 記事のIdをURLから取得できない場合は新規作成
      setLoading(false);
    }
  }, []);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // 少なくとも1つのフォームにバリデーションエラーが発生している場合は処理を中断
    const isValidTitle = validateTitle();
    const isValidContent = validateContent();
    if (!(isValidTitle && isValidContent)) return;
    postArticle();
  };

  const postArticle = async () => {
    try {
      const res = await fetch('api/admin/article' + (id ? '/' + id : ''), {
        method: 'POST',
        headers: {
          'content-type': 'application/json',
          Authorization:
            'Bearer ' + localStorage.getItem('portfolio-jwt-token'),
        },
        body: JSON.stringify({
          title: title,
          content: content,
        }),
      });
      const json = await res.json();
      if (res.ok) {
        flushDispath({
          type: FlushActionType.VISIBLE,
          payload: {
            type: FlushType.SUCCESS,
            message: '記事の' + (id ? '更新' : '投稿') + 'に成功しました',
          },
        });
        setPostDone(true);
      } else {
        throw new Error(json.message);
      }
    } catch (error) {
      flushDispath({
        type: FlushActionType.VISIBLE,
        payload: {
          type: FlushType.ERROR,
          message:
            '記事の' + (id ? '更新' : '投稿') + 'に失敗しました。' + error,
        },
      });
    }
    setLoading(false);
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
  };

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
  };

  return !props.isLoggedIn || postDone ? (
    <Redirect to={id ? '/management' : '/'} />
  ) : loading ? (
    <p>loading...</p>
  ) : (
    // ログイン状態かつ記事未投稿状態の場合
    <div className="justify-content-center">
      <div>
        <h1 className="mb-3">記事の投稿</h1>
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
                value={title}
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
                value={content}
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
  );
};
