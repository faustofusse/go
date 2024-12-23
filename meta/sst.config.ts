/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
    app(input) {
        return {
            name: "fausto",
            removal: input?.stage === "prod" ? "retain" : "remove",
            protect: ["prod"].includes(input?.stage),
            home: "aws",
        };
    },
    async run() {
        const server = new sst.aws.Function("Function", {
            name: "fausto-golib",
            runtime: "go",
            handler: "./main.go",
            timeout: "3 minutes",
        });
        const api = new sst.aws.ApiGatewayV1("Api", {
            endpoint: { type: "regional" },
            domain: {
                name: "go.fausto.ar",
                dns: sst.cloudflare.dns(),
            },
        });
        api.route("ANY /", server.arn, {});
        api.route("ANY /{proxy+}", server.arn, {});
        api.deploy();
    },
});
