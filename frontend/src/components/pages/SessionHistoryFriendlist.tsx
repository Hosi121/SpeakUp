import { Stack } from "@mui/system";
import { Box, Container, Tab } from "@mui/material";
import TopSection from "../utils/TopSection";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import { useEffect, useState } from "react";
import TabContext from "@mui/lab/TabContext";
import TabList from "@mui/lab/TabList";
import TabPanel from "@mui/lab/TabPanel";
import FriendList from "../utils/FriendList";
import SessionHistory from "../utils/SessionHistory";
import { fetchSessionHistory } from "../../services/appData";
import type { SessionHistoryItem } from "../../types/types";

const SessionHistoryFriendlistContainer = () => {
  const [history, setHistory] = useState<SessionHistoryItem[]>([]);
  const [value, setValue] = useState("1");

  const handleChange = (_: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
  };

  useEffect(() => {
    const loadHistory = async () => {
      try {
        const data = await fetchSessionHistory();
        setHistory(data);
      } catch (error) {
        console.error("Failed to fetch session history", error);
      }
    };
    loadHistory();
  }, []);

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
                <TabList
                  onChange={handleChange}
                  sx={{ display: "grid", placeContent: "center" }}
                >
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
