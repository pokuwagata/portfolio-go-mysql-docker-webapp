"use strict";
/// <reference path="./@types/graphql.d.ts" />
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var apollo_server_express_1 = require("apollo-server-express");
var express_1 = __importDefault(require("express"));
var resolvers_1 = __importDefault(require("./resolvers"));
var graphql_playground_middleware_express_1 = __importDefault(require("graphql-playground-middleware-express"));
// import { readFileSync } from 'fs';
var typeDefs_graphql_1 = __importDefault(require("./typeDefs.graphql"));
// const typeDefs = importSchema("./typeDefs.graphql");
// const typeDefs = readFileSync(`./typeDefs.graphql`, 'UTF-8');
var app = express_1.default();
var server = new apollo_server_express_1.ApolloServer({ typeDefs: typeDefs_graphql_1.default, resolvers: resolvers_1.default });
server.applyMiddleware({ app: app });
app.get("/", function (req, res) { return res.end("hellooooooooo"); });
app.get("/playground", graphql_playground_middleware_express_1.default({ endpoint: "/graphql" }));
app.listen({ port: 4000 }, function () {
    // console.log(`GraphQL Server running @ http://localhost:4000${server.graphqlPath}`)
    return console.log("port 4000");
});
