### CubiCasa Test

**Read First**

- Test case is done for repo(local tested) but not api but api testing need manually tested
- Design is diff with standard because new idea(7 day is short to fix design, so I choose design base on my view)
- Config is running for docker. if test local please change data in .env then run postgresql in local
- Details is provided in doc and some comment in code please check it for api some other information

**Todo**

- [x] Implement a Creation for each _hub, team, and user_.
- [x] Implement a Search which will return _team and hub information_.
- [x] Implement a Join for user into team, and team into hub (for simplicity: one user belongs to one team, one team
  belongs to one hub).
- [x] Write the test cases
- [x] Provide a SQL script which creates tables needed for the API.
- [x] Good to use `docker/docker-compose` for local development setup(not mandatory) (basic setup)
- [x] Good to provide the solution with security concern (base Auth)

**Run**
Command : docker-compose up

**Framework used**
- GIN
- GORM
- Testify
- Sql (docker init)
