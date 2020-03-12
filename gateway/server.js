const { ApolloServer } = require('apollo-server');
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'articles', url: 'http://localhost:9092/query' },
        { name: 'users', url: 'http://localhost:9091/query' },
    ],
});

const server = new ApolloServer({
    gateway,

    subscriptions: false,
});

server.listen().then(({ url }) => {
    console.log(`🚀 Server ready at ${url}`);
});