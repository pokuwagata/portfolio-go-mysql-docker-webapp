var photos = [];

const resolvers = {
  Query: {
    // 写真を格納した配列の長さを返す
    totalPhotos: () => photos.length
  },
  // postPhotoミューテーションと対応するリゾルバ
  Mutation: {
    postPhoto(parent: any, args: any) {
      photos.push(args);
      return true;
    }
  }
};

export default resolvers;
