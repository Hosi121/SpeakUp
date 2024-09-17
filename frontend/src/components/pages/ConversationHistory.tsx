import React, { useState, useEffect } from "react";
import { Accordion, AccordionSummary, AccordionDetails, Typography, Container, Stack, Box } from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import conversation_history from "../../mock/conversation_history.json"; // Import the JSON file directly
import TopSection from "../utils/TopSection";
import DescriptionIcon from "@mui/icons-material/Description";

type SessionData = {
  date: string;
  previousDate: string;
  sessions: number;
  completionRate: string;
  comment: string;
  examples: string[];
};

const ConversationHistoryContainer = () => {
  const [data, setData] = useState<SessionData[]>([]);

  useEffect(() => {
    setData(conversation_history);
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
        <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
          <Box sx={{ mb: 3 }}>
            <Typography variant="h4" sx={{ mb: 4, fontWeight: "bold", textAlign: "left" }}>
              <DescriptionIcon sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }} />
              会話の記録
            </Typography>
          </Box>
          {data.map((item, index) => (
            <Accordion
              key={index}
              sx={{
                mb: 2,
                borderRadius: 3,
                "&:before": {
                  display: "none",
                },
              }}
            >
              <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                <Typography>{item.date}</Typography>
              </AccordionSummary>
              <AccordionDetails sx={{ p: 3 }}>
                <Box sx={{ display: "flex", justifyContent: "space-around", alignItems: "center", mb: 3 }}>
                  <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
                    <Typography variant="h4" sx={{ color: "primary.main" }}>
                      {item.sessions}
                    </Typography>
                    <Typography textAlign="center">セッション数</Typography>
                  </Box>
                  <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
                    <Typography variant="h4" sx={{ color: "primary.main" }}>
                      {item.completionRate}
                    </Typography>
                    <Typography>満足度</Typography>
                  </Box>
                </Box>
                <Typography variant="body2" fontSize={"large"} textAlign={"left"}>
                  感想
                </Typography>
                <Typography variant="body2" textAlign={"left"}>
                  {item.comment}
                </Typography>
                <Typography variant="body2" fontSize={"large"} textAlign={"left"} sx={{ mt: 3 }}>
                  学んだ表現
                </Typography>
                <Typography variant="body2" textAlign={"left"}>
                  {item.examples.join(", ")}
                </Typography>
              </AccordionDetails>
            </Accordion>
          ))}{" "}
        </Stack>
      </Container>
    </Container>
  );
};

export const ConversationHistory = () => {
  return (
    <BottomNavigationTemplate value="other">
      <ConversationHistoryContainer />
    </BottomNavigationTemplate>
  );
};
