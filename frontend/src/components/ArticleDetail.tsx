import * as React from 'react';
import { FlushType } from './Flush';
import { FlushDispatchContext, FlushActionType } from './FlushProvider';
import { useLocation, Redirect, Link } from 'react-router-dom';

export interface ArticleDetailProps {
  isLoggedIn: boolean;
}

export const ArticleDetail = (props: ArticleDetailProps) => {
  const [loading, setLoading] = React.useState(true);
  const flushDispath = React.useContext(FlushDispatchContext);
  const [article, setArticle] = React.useState({
    title: '',
    content: '',
    username: '',
    updatedAt: '',
  });
  const [canFetch, setCanFetch] = React.useState(true);

  const id = new URLSearchParams(useLocation().search).get('id');

  const fetchArticle = async (id: string) => {
    setLoading(true);
    try {
      const res = await fetch('api/articles/' + id, {
        method: 'GET',
      });
      const json = await res.json();
      if (res.ok) {
        flushDispath({
          type: FlushActionType.HIDDEN,
        });
        setArticle(json);
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
      setArticle(null);
      setLoading(false);
    }
  };

  React.useEffect(() => {
    id ? fetchArticle(id) : setCanFetch(false);
  }, []);

  return !canFetch ? (
    <Redirect to="/" />
  ) : loading ? (
    <p>loading...</p>
  ) : (
    <div>
      <div className="d-flex align-items-center mb-2">
        <h2>{article.title}</h2>
        <div className="ml-auto">
          {new Date(article.updatedAt).toLocaleString()}
        </div>
      </div>
      <p>{article.content}</p>
      <div className="d-flex">
        <Link to="/">戻る</Link>
        <p className="ml-auto">posted by {article.username} </p>
      </div>
    </div>
  );
};
