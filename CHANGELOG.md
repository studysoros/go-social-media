# Changelog

## [1.1.0](https://github.com/studysoros/go-social-media/compare/v1.0.0...v1.1.0) (2026-01-11)


### Features

* update api version workflow ([9b55c76](https://github.com/studysoros/go-social-media/commit/9b55c764480709ceb055396048dd8afa1871cf46))

## 1.0.0 (2026-01-11)


### Features

* add authorization, role based ([cf26b39](https://github.com/studysoros/go-social-media/commit/cf26b3947d3fb3ef70469662bcbff2f27d163328))
* add automation workflow ([e29b5b1](https://github.com/studysoros/go-social-media/commit/e29b5b1e66259f0f7b78a069b2897a77fae107ca))
* add comments creation and seed data ([18a8a28](https://github.com/studysoros/go-social-media/commit/18a8a2848fb2be91f819b557ac463e6b538d5716))
* add comments to posts ([3916883](https://github.com/studysoros/go-social-media/commit/39168832be3682567c1f1f4c54a4a98d2f6ca121))
* add cors ([a2efccd](https://github.com/studysoros/go-social-media/commit/a2efccd1c4a25706c6d09df7261a7e2cd609e7ba))
* add db and migrations ([43a3125](https://github.com/studysoros/go-social-media/commit/43a3125e8f5c6e7a18a5dbdf098f7bb88448f07f))
* add env handling ([e4c4944](https://github.com/studysoros/go-social-media/commit/e4c49446fec5e4050c9ba6fe2b6270d999f21c4b))
* add follow, unfollow features ([427fb28](https://github.com/studysoros/go-social-media/commit/427fb28100d93fc9ea7ea12dabc316e80b01113e))
* add get user by id route ([e8d2205](https://github.com/studysoros/go-social-media/commit/e8d2205bf496d770779590c3578724709f21a643))
* add indexes ([db89b1b](https://github.com/studysoros/go-social-media/commit/db89b1bd82ff5ff571bba5a3a59242774f07ed35))
* add metrics ([ba94758](https://github.com/studysoros/go-social-media/commit/ba9475896629a876124072d85f1d371bfa96fb95))
* add mock implementations for testing user authentication and store functionality ([284d518](https://github.com/studysoros/go-social-media/commit/284d51851dfa6f70a548dd9d043e4d9648fc7f59))
* add optimistic concurrency control ([a063d0e](https://github.com/studysoros/go-social-media/commit/a063d0ea746be3cd16f42f12b75106c232d6d136))
* add pagination, filtering and sorting ([8031650](https://github.com/studysoros/go-social-media/commit/803165035c44f19e3dfbfa20c35c8d1667595df1))
* add payload validation on post creation ([7d98ace](https://github.com/studysoros/go-social-media/commit/7d98acee3b024dee6df3d1dc0dde8e9322d52b4e))
* add post creation and retrieval endpoints ([a997f97](https://github.com/studysoros/go-social-media/commit/a997f97e939f1ff939f96e9ea945cd1bb11cafe1))
* add sql query timeout ([23c1b22](https://github.com/studysoros/go-social-media/commit/23c1b2271b0864b676d5553097175aa2fff46012))
* add structured logger ([3707a5e](https://github.com/studysoros/go-social-media/commit/3707a5e7aed163e66eb0c71083a3456a3210813f))
* add swagger docs ([15ae642](https://github.com/studysoros/go-social-media/commit/15ae642c678efb9d8ffcbff1665318153bc187b9))
* add update and delete routes ([d30068e](https://github.com/studysoros/go-social-media/commit/d30068e42b2236680d015aecb3b3167e490d9904))
* authentication ([d493348](https://github.com/studysoros/go-social-media/commit/d4933485062135915a4412d57ab74a7fcf8566eb))
* confirmation ui ([30cf11a](https://github.com/studysoros/go-social-media/commit/30cf11ab62686200c5b6f7c8826216eab227cf22))
* email confirmation flow ([7a8d151](https://github.com/studysoros/go-social-media/commit/7a8d151c77ce4d77f779d3e90fe198fc89f287f3))
* get user feed ([5c296c5](https://github.com/studysoros/go-social-media/commit/5c296c59fb678ad6f68841384016dffada962e87))
* implement errors pkg ([b96eeb6](https://github.com/studysoros/go-social-media/commit/b96eeb61007cd7396783d5b79e26fadaab8f7192))
* implement graceful server shutdown ([abc508c](https://github.com/studysoros/go-social-media/commit/abc508c1911ba07366963d4bc133c069bee8ab7c))
* implement rate limiter ([e09c0bf](https://github.com/studysoros/go-social-media/commit/e09c0bf8448ebd1a03d9698d5c23e65e460ee0a0))
* implement redis caching for user data ([1c40954](https://github.com/studysoros/go-social-media/commit/1c40954b66dac3bf5ccad903e0280b8899d8dfd2))
* implement user registration and activation via invitation token ([ed1d2c4](https://github.com/studysoros/go-social-media/commit/ed1d2c489e16905bc8389e6bffd5f895288f8204))
* init http server and api ([8a06f07](https://github.com/studysoros/go-social-media/commit/8a06f075ff996499837b03e82a882a69216fa204))
* release please script ([ea47471](https://github.com/studysoros/go-social-media/commit/ea4747110eaeaa520476c521bd0df2a1f92d1d22))
* standardize json responses ([5baaddc](https://github.com/studysoros/go-social-media/commit/5baaddc4d5031dd8a9dc6f2c63a126edca9438dd))


### Bug Fixes

* change user fields to string ([2ef8433](https://github.com/studysoros/go-social-media/commit/2ef8433a173fab4437e310f54d69aed0ccc77446))
* check for retry error ([2b94688](https://github.com/studysoros/go-social-media/commit/2b9468848214b772acb8d49f83b4941efb8c2d94))
* correct column mapping in post creation ([436ccaf](https://github.com/studysoros/go-social-media/commit/436ccaf07bb39913fe46c785396a3157e6b0d141))
* sql syntax error in PostsStore.Create ([8f58821](https://github.com/studysoros/go-social-media/commit/8f5882117bd8b32d7698b927ab7a97eb8c38539e))
* sql syntax error in UsersStore.Create ([ee5d871](https://github.com/studysoros/go-social-media/commit/ee5d87104fd65fd1828b61e5e0cfd54dcded0b33))
* unused middleware ([8a000e5](https://github.com/studysoros/go-social-media/commit/8a000e5c2036feb4485615692d4b68a8bc2791ce))
