import { useEffect, useState } from "react";
import { Stack } from "@mui/system";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import { Button, Container, Typography } from "@mui/material";
import TopSection from "../utils/TopSection";
import sessions from "../../mock/sessions.json";
import { useNavigate } from "react-router-dom";

type SessionData = {
  dateTime: string;
  sessions: number[];
};

const SessionListContainer = () => {
  const [sessionData, setSessionData] = useState<SessionData[]>([]);

  useEffect(() => {
    setSessionData(sessions);
  }, []);

  const navigate = useNavigate();

  return (
    <Container sx={{ display: "flex", flexDirection: "column", justifyContent: "space-between", height: "100vh" }}>
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
          <Typography variant="h5" sx={{ fontWeight: "bold", textAlign: "left" }}>
            参加予定のセッション
          </Typography>
          <Button
            sx={{
              width: "100%",
              backgroundColor: "secondary.main",
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
              mt: 3,
              mb: 3,
              p: 3,
              boxSizing: "border-box",
              borderRadius: 3,
            }}
            variant="contained"
            onClick={() => navigate("/waiting")}
          >
            <Typography sx={{ mb: 2, color: "primary.main", fontSize: "1.3rem", fontWeight: "bolder", textAlign: "center" }}>{sessionData[0]?.dateTime}</Typography>
            {sessionData[0]?.sessions.map((session) => <Typography sx={{ color: "primary.main", textAlign: "center", fontSize: "0.9rem" }}>セッション{session}</Typography>)}
          </Button>
        </Stack>
      </Container>
    </Container>
  );
};

export const SessionList = () => {
  return (
    <BottomNavigationTemplate value="session">
      <SessionListContainer />
    </BottomNavigationTemplate>
  );
};
