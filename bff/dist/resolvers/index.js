"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var photos = [];
var resolvers = {
    Query: {
        // 写真を格納した配列の長さを返す
        totalPhotos: function () { return photos.length; }
    },
    // postPhotoミューテーションと対応するリゾルバ
    Mutation: {
        postPhoto: function (parent, args) {
            photos.push(args);
            return true;
        }
    }
};
exports.default = resolvers;
