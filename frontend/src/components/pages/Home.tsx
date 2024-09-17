import { useState, useEffect } from "react";
import { Box, Container, List, ListItem, ListItemText } from "@mui/material";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import TopSection from "../utils/TopSection";
import TopWaves from "../utils/TopWaves";
import HomeLogo from "../../assets/homeLogo";
import { IconButton } from "../utils/IconButton";
import { LibraryBooks, Mic } from "@mui/icons-material";
import session from "../../mock/sessions.json";

type SessionData = {
  dateTime: string;
  sessions: number[];
};

const HomeContainer = () => {
  const [sessionData, setSessionData] = useState<SessionData>();

  useEffect(() => {
    setSessionData(session);
  });

  return (
    <Container sx={{ display: "flex", flexDirection: "column", justifyContent: "space-between", height: "calc(100vh - 70px)" }}>
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Container sx={{ mb: 3 }}>
          <TopWaves isFlipped={false} />
          <Box sx={{ margin: "0 calc(50% - 50vw)", width: "100vw", height: "15vh", backgroundColor: "secondary.main", display: "grid", placeItems: "center" }}>
            <HomeLogo style={{ width: "70%" }} />
          </Box>
          <TopWaves isFlipped={true} />
        </Container>
        <Box sx={{ width: "100%", backgroundColor: "secondary.main", display: "flex", flexDirection: "column", alignItems: "center", mt: 3, mb: 3, p: 3, boxSizing: "border-box", borderRadius: 3 }}>
          <h2>直近の参加予定</h2>
          <List>
            <ListItem>
              <ListItemText primary={sessionData?.dateTime} primaryTypographyProps={{ color: "primary.main", fontSize: "1.3rem", fontWeight: "bolder", textAlign: "center" }} />
            </ListItem>
            {sessionData?.sessions.map((session) => (
              <ListItem sx={{ p: 0 }}>
                <ListItemText primary={"セッション" + session} primaryTypographyProps={{ color: "primary.main", textAlign: "center" }} />
              </ListItem>
            ))}
          </List>
        </Box>
        <Box
          sx={{
            display: "flex",
            placeContent: "center",
            flexWrap: "wrap",
            justifyContent: "space-between",
          }}
        >
          <IconButton icon={<LibraryBooks sx={{ fontSize: "60px" }} />} text="記録" url="record" />
          <IconButton icon={<Mic sx={{ fontSize: "60px" }} />} text="セッション" url="sessionlist" />
        </Box>
      </Container>
    </Container>
  );
};

export const Home = () => {
  return (
    <BottomNavigationTemplate value="home">
      <HomeContainer />
    </BottomNavigationTemplate>
  );
};
