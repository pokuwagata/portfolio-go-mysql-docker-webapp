import fetch from "node-fetch";
import { ApolloError } from "apollo-server-express";

var photos = [];

const resolvers = {
  Query: {
    // 写真を格納した配列の長さを返す
    totalPhotos: async () => {
      return await fetch("http://api:5000").then((res) => res.text());
    },
  },
  // postPhotoミューテーションと対応するリゾルバ
  Mutation: {
    postPhoto(_: any, args: any) {
      photos.push(args);
      return true;
    },
    async createSession(_: any, args: any) {
      return await fetch("http://api:5000/session", {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          username: args.input.username,
          password: args.input.password,
        }),
      }).then((res) =>
        res.json().then((json) => {
          if (res.ok) {
            return { username: json.username, token: json.token };
          } else {
            throw new ApolloError(json.message);
          }
        })
      );
    },
  },
};

export default resolvers;
