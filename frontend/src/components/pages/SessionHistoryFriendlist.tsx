import { Stack } from "@mui/system";
import { Box, Container, Tab } from "@mui/material";
import TopSection from "../utils/TopSection";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import React, { useEffect, useState } from "react";
import TabContext from "@mui/lab/TabContext";
import TabList from "@mui/lab/TabList";
import TabPanel from "@mui/lab/TabPanel";
import FriendList from "../utils/FriendList";
import SessionHistory from "../utils/SessionHistory";
import History from "../../mock/session_history.json";

interface History {
  avatar: string;
  user: string;
  theme: string;
  date: string;
  rank: number;
  friedstate: string;
}

const SessionHistoryFriendlistContainer = () => {
  const [history, setHistory] = useState<History[]>([]);
  const [value, setValue] = React.useState("1");

  const handleChange = (event: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
  };

  useEffect(() => {
    setHistory(History);
  });

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "100vh",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Stack sx={{ width: "100%" }}>
          <Box sx={{ width: "100%" }}>
            <TabContext value={value}>
              <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
                <TabList onChange={handleChange} sx={{ display: "grid", placeContent: "center" }}>
                  <Tab label=" セッション履歴" value="1" />
                  <Tab label="フレンド" value="2" />
                </TabList>
              </Box>
              <TabPanel value="1">
                <SessionHistory history={history} />
              </TabPanel>
              <TabPanel value="2">
                <FriendList />
              </TabPanel>
            </TabContext>
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};

export const SessionHistoryFriendlist = () => {
  return (
    <BottomNavigationTemplate value="record">
      <SessionHistoryFriendlistContainer />
    </BottomNavigationTemplate>
  );
};
