import { useEffect, useState } from "react";
import { Box, Typography, Paper, List, ListItem, ListItemIcon, ListItemText, Container, Stack } from "@mui/material";
import EmojiEventsIcon from "@mui/icons-material/EmojiEvents";
import trophyData from "../../assets/trophy.json";
import TopSection from "../utils/TopSection"; // Import the new TopSection component
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";

const StatsContainer = () => {
  const [trophies, setTrophies] = useState([]);

  // Fetch the trophy data from trophy.json
  useEffect(() => {
    setTrophies(trophyData.trophies);
  }, []);

  return (
    <Container sx={{ display: "flex", flexDirection: "column", justifyContent: "space-between", height: "100vh" }}>
      <Container sx={{ pt: 3 }}>
        {/* Top Section */}
        <TopSection />
        <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
          <Box sx={{ mb: 3 }}>
            <Typography variant="h4" sx={{ mb: 4, fontWeight: "bold", textAlign: "left" }}>
              <EmojiEventsIcon sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }} />
              参加データ
            </Typography>

            <Paper sx={{ m: "0 auto", p: 4, width: "90%", backgroundColor: "secondary.main", borderRadius: "16px" }}>
              <Box sx={{ display: "flex", justifyContent: "space-around", alignItems: "center" }}>
                <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
                  <Typography variant="h4" sx={{ color: "#FF007F" }}>
                    5回
                  </Typography>
                  <Typography textAlign="center">回数</Typography>
                </Box>
                <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
                  <Typography variant="h4" sx={{ color: "#FF007F" }}>
                    20
                  </Typography>
                  <Typography>セッション数</Typography>
                </Box>
              </Box>
            </Paper>
          </Box>

          {/* Trophies List */}
          <Box>
            <Typography variant="h6" sx={{ mb: 1, textAlign: "left" }}>
              与えられたトロフィー
            </Typography>
            <List>
              {trophies.map((trophy, index) => (
                <ListItem key={index}>
                  <ListItemIcon>
                    <EmojiEventsIcon sx={{ fontSize: 30 }} />
                  </ListItemIcon>
                  <ListItemText primary={trophy.description} />
                </ListItem>
              ))}
            </List>
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};

export const Stats = () => {
  return (
    <BottomNavigationTemplate value="other">
      <StatsContainer />
    </BottomNavigationTemplate>
  );
};
