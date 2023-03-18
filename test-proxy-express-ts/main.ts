import Express from "express";
import { createProxyMiddleware, RequestHandler } from "http-proxy-middleware";

const app = Express();

const db: Record<string, string> = {
  "my-doc": "preview/1671133747717_11b6e5f6-892d-4164-ab25-e32528130d5a/doc",
  "another-doc": "preview/1671131676871_32abc487-6a98-4e87-b982-92358d331bfa/doc",
};

const proxy: RequestHandler = createProxyMiddleware({
  target: "http://some-bucket.s3.amazonaws.com",
  secure: false,
  onProxyReq: (proxyReq, req, res) => {
    proxyReq.setHeader("host", "some-bucket.s3.amazonaws.com");
    if (!proxyReq.path.endsWith(".js")) {
      proxyReq.path = db[proxyReq.path.split("/")[3]] + "/index.html";
    } else {
      proxyReq.path =
        db[proxyReq.path.split("/")[3]] +
        proxyReq.path.split(proxyReq.path.split("/")[4])[1];
    }
    console.log(proxyReq.path);
  },
});

// /docs/preview/lol/v1/
// /docs/preview/lol/v1/_next/dsfofosk.js
// /1671131676871_32abc487-6a98-4e87-b982-92358d331bfa/doc/_next/dsfofosk.js
app.use("/docs/preview/:name/:version/*", proxy);

app.listen(3000);
