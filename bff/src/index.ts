// const { ApolloServer } = require(`apollo-server-express`);
import {Request, Response} from "express";
import express from 'express';

var app = express();

app.get(`/`, (req :Request, res :Response) => res.end(`hellooo`));

app.listen({ port: 4000 }, () =>
  // console.log(`GraphQL Server running @ http://localhost:4000${server.graphqlPath}`)
  console.log("port 4000")
);
