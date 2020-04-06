import { ApolloServer } from "apollo-server-express";
import { Request, Response } from "express";
import express from "express";
import resolvers from "./resolvers";
import expressPlayground from "graphql-playground-middleware-express";
import "graphql-import-node";
import typeDefs from "./typeDefs.graphql";
import { createStream } from "rotating-file-stream";
import path from "path";
import morgan from "morgan";
const { ApolloLogExtension } = require("apollo-log");

var app = express();

const dev = app.get("env") === "DEV";
const extensions = [() => new ApolloLogExtension()];

const server = new ApolloServer({
  typeDefs,
  resolvers,
  extensions,
  debug: dev,
});
server.applyMiddleware({ app });

const accessLogStream = createStream("access.log", {
  interval: "1d", // rotate daily
  path: path.join(__dirname, "logs"),
});
app.use(morgan("combined", { stream: accessLogStream }));

if (dev) {
  app.get(`/playground`, expressPlayground({ endpoint: `/graphql` }));
}

app.get(`/health-check`, (req: Request, res: Response) => res.end(`It works!`));

app.listen({ port: 4000 }, () =>
  console.log(
    `GraphQL Server running @ http://localhost:4000${server.graphqlPath}\n` +
    `mode: ${dev ? 'dev' : 'prod'}`
  )
);
