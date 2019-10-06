import * as React from 'react';
import { FlushState } from './App';
import { FlushType } from './Flush';

export interface ArticleRemoveButtonProps {
  setFlushState: (state: FlushState) => void;
  fetchArticles: (pageNumber: number) => void;
  id: number;
}

export const ArticleRemoveButton = (props: ArticleRemoveButtonProps) => {
  const onClickButton = async (event: React.MouseEvent, id: number) => {
    event.preventDefault();
    try {
      const res = await fetch('api/admin/article?id=' + id, {
        method: 'DELETE',
        headers: {
          Authorization:
            'Bearer ' + localStorage.getItem('portfolio-jwt-token'),
        },
      });
      const json = await res.json();
      if (res.ok) {
        props.setFlushState({
          isDisplay: true,
          type: FlushType.SUCCESS,
          message: '記事の削除に成功しました。',
        });
        props.fetchArticles(1);
      } else {
        throw new Error(json.message);
      }
    } catch (error) {
      props.setFlushState({
        isDisplay: true,
        type: FlushType.ERROR,
        message: '記事の削除に失敗しました。' + error,
      });
    }
  };

  return (
    <button
      type="button"
      className="btn btn-danger"
      onClick={(event: React.MouseEvent) => onClickButton(event, props.id)}
    >
      選択した記事を削除
    </button>
  );
};
