import * as React from 'react';
import { FlushType } from './Flush';
import { FlushDispatchContext, FlushActionType } from './FlushProvider';
import { Pagination } from './Pagination';
import { ArticleRow } from './ArticleRow';

export interface ArticleListProps {}

export const ArticleList = (props: ArticleListProps) => {
  const [loading, setLoading] = React.useState(true);
  const [articleList, setArticleList] = React.useState([]);
  const [pageNumber, setPageNumber] = React.useState(1);
  const [maxPageNumber, setMaxNumber] = React.useState(10);

  const flushDispath = React.useContext(FlushDispatchContext);

  const fetchArticles = async (pageNumber: number) => {
    try {
      const res = await fetch('api/admin/articles?number=' + pageNumber, {
        method: 'GET',
        headers: {
          'content-type': 'application/json',
          Authorization:
            'Bearer ' + localStorage.getItem('portfolio-jwt-token'),
        },
      });
      const json = await res.json();
      if (res.ok) {
        flushDispath({
          type: FlushActionType.HIDDEN,
        });
        setArticleList(json.articles);
        json.maxNumber && setMaxNumber(json.maxNumber);
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
      setArticleList([]);
      setLoading(false);
    }
  };

  React.useEffect(() => {
    fetchArticles(pageNumber);
  }, [pageNumber]);

  return (
    <div>
      <div className="container mb-5">
        {loading ? (
          <p>loading...</p>
        ) : articleList.length === 0 ? (
          <p>記事が見つかりません</p>
        ) : (
          articleList.map(article => (
            <ArticleRow
              key={article.id}
              id={article.id}
              username={article.username}
              title={article.title}
              content={article.content}
              updatedAt={article.updatedAt}
            />
          ))
        )}
      </div>
      <Pagination
        pageNumber={pageNumber}
        setPageNumber={setPageNumber}
        maxPageNumber={maxPageNumber}
        setMaxPageNumber={setPageNumber}
        loading={loading}
      />
    </div>
  );
};
