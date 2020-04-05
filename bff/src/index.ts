import { ApolloServer } from "apollo-server-express";
import { Request, Response } from "express";
import express from "express";
import resolvers from "./resolvers";
import expressPlayground from "graphql-playground-middleware-express";
import "graphql-import-node";
import typeDefs from "./typeDefs.graphql";

var app = express();

const server = new ApolloServer({
  typeDefs,
  resolvers,
  debug: true
});
server.applyMiddleware({ app });

app.get(`/`, (req: Request, res: Response) => res.end(`hellooooooooo`));
app.get(`/playground`, expressPlayground({ endpoint: `/graphql` }));

app.listen({ port: 4000 }, () =>
  // console.log(`GraphQL Server running @ http://localhost:4000${server.graphqlPath}`)
  console.log("port 4000")
);
