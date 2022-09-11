const express = require("express");
const app = express();
const jwt = require("express-jwt");
const jwks = require("jwks-rsa");
const port = process.env.PORT || 8080;

const isRevokedCallback = (req, payload, done) => {
  var issuer = payload.iss;
  var tokenId = payload.jti;

  return done();
};

const jwtCheck = jwt({
  secret: jwks.expressJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: process.env.JWK_URI,
  }),
  audience: process.env.JWT_AUDIENCE,
  issuer: process.env.JWT_ISSUER,
  algorithms: ["RS256"],
  isRevoked: isRevokedCallback,
});

app.use(jwtCheck);

app.get("/authorized", (req, res) => res.send("Secured Resource"));

app.listen(port, () => console.info("onAppStart", { port }));
