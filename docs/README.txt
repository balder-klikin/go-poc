Running Suite: go-poc/app Suite
===============================
Random Seed: 1435190858
Will run 6 of 6 specs

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET   /check                    --> github.com/balder-klikin/go-poc/app.check (3 handlers)
[GIN-debug] GET   /ping                     --> github.com/balder-klikin/go-poc/app.pong (3 handlers)
[GIN-debug] POST  /upload                   --> github.com/balder-klikin/go-poc/app.uploadImageS3 (3 handlers)
••••••
Ran 6 of 6 Specs in 0.003 seconds
SUCCESS! -- 6 Passed | 0 Failed | 0 Pending | 0 Skipped PASS

Ginkgo ran 1 suite in 652.319957ms
Test Suite Passed
