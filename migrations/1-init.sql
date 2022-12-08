-- +migrate Up
CREATE SCHEMA IF NOT EXISTS testdata;
-- Drop
DROP TABLE IF EXISTS testdata.Hubs;
DROP TABLE IF EXISTS testdata.Teams;
DROP TABLE IF EXISTS testdata.Users;
-- Hub 
CREATE TABLE testdata.Hubs
(
    HUB_ID     SERIAL,
    HUB_NAME   Char(20) not null,
    HUB_REGION Char(20) not null,
    PRIMARY KEY (HUB_ID)
);
INSERT INTO testdata.Hubs(HUB_NAME, HUB_REGION)
VALUES ('Vietnam', 'Asia');
INSERT INTO testdata.Hubs(HUB_NAME, HUB_REGION)
VALUES ('England', 'EU');
-- Team
CREATE TABLE testdata.Teams
(
    TEAM_ID   SERIAL,
    TEAM_NAME Char(20) not null,
    HUB_ID    INTEGER,
    PRIMARY KEY (TEAM_ID),
    CONSTRAINT fk_hub FOREIGN KEY (HUB_ID) REFERENCES testdata.Hubs (HUB_ID) ON DELETE CASCADE ON UPDATE CASCADE
);
INSERT INTO testdata.Teams(TEAM_NAME,
                           HUB_ID)
VALUES ('DEVOP', 1);
INSERT INTO testdata.Teams(TEAM_NAME,
                           HUB_ID)
VALUES ('DEVCORE', 1);
INSERT INTO testdata.Teams(TEAM_NAME,
                           HUB_ID)
VALUES ('DEVUI', 2);
INSERT INTO testdata.Teams(TEAM_NAME,
                           HUB_ID)
VALUES ('BA', 2);
--User
CREATE TABLE testdata.Users
(
    USER_ID   SERIAL,
    USER_NAME Char(20) not null,
    ROLE      Char(20) not null,
    TEAM_ID   INTEGER,
    PRIMARY KEY (USER_ID),
    CONSTRAINT fk_team FOREIGN KEY (TEAM_ID) REFERENCES testdata.Teams (TEAM_ID) ON DELETE CASCADE
);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('DEVOP-1', 'DEVOP', 1);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('DEVOP-2', 'DEVOP', 1);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('DEVCORE-1', 'DEVCORE', 1);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('DEVCORE-2', 'DEVCORE', 1);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('DEVUI-1', 'DEVUI', 2);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('DEVUI-2', 'DEVUI', 2);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('BA-1', 'BA', 2);
INSERT INTO testdata.Users (USER_NAME, ROLE, TEAM_ID)
VALUES ('BA-2', 'BA', 2);
