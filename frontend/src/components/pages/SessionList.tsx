import { useEffect, useState } from "react";
import { Stack } from "@mui/system";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import { Accordion, AccordionDetails, AccordionSummary, Button, Checkbox, Container, FormControlLabel, FormGroup, Typography } from "@mui/material";
import TopSection from "../utils/TopSection";
import sessions from "../../mock/sessions.json";
import { useNavigate } from "react-router-dom";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

type SessionData = {
  dateTime: string;
  sessions: number[];
};

const SessionListContainer = () => {
  const [sessionData, setSessionData] = useState<SessionData[]>([]);
  const [selectedSessions, setSelectedSessions] = useState<{ [key: string]: number[] }>({});

  useEffect(() => {
    setSessionData(sessions);
    const initialSelections: { [key: string]: number[] } = {};
    sessions.forEach((session) => {
      initialSelections[session.dateTime] = [...session.sessions];
    });
    setSelectedSessions(initialSelections);
  }, []);

  const handleCheckboxChange = (dateTime: string, session: number) => {
    setSelectedSessions((prevSelected) => {
      const newSelected = { ...prevSelected };
      if (newSelected[dateTime].includes(session)) {
        newSelected[dateTime] = newSelected[dateTime].filter((s) => s !== session);
      } else {
        newSelected[dateTime].push(session);
      }
      return newSelected;
    });
  };

  const handleSave = (dateTime: string) => {
    setSessionData((prevData) =>
      prevData.map((data) => {
        if (data.dateTime === dateTime) {
          return { ...data, sessions: selectedSessions[dateTime] };
        }
        return data;
      })
    );
  };

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
        <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
          <Typography variant="h5" sx={{ fontWeight: "bold", textAlign: "left", mb: 3 }}>
            参加可能なセッション
          </Typography>
          {sessionData.map((data, index) => (
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
              <AccordionSummary expandIcon={<ExpandMoreIcon />} sx={{ height: "50px" }}>
                <Typography fontWeight="bolder" fontSize="large" color="primary.main" sx={{ display: "block", m: "auto" }}>
                  {data.dateTime}
                </Typography>
              </AccordionSummary>
              <AccordionDetails sx={{ p: 2 }}>
                <FormGroup>
                  <FormControlLabel control={<Checkbox checked={selectedSessions[data.dateTime]?.includes(1)} onChange={() => handleCheckboxChange(data.dateTime, 1)} />} label="セッション1" sx={{ m: "0 auto" }} />
                  <FormControlLabel control={<Checkbox checked={selectedSessions[data.dateTime]?.includes(2)} onChange={() => handleCheckboxChange(data.dateTime, 2)} />} label="セッション2" sx={{ m: "0 auto" }} />
                  <FormControlLabel control={<Checkbox checked={selectedSessions[data.dateTime]?.includes(3)} onChange={() => handleCheckboxChange(data.dateTime, 3)} />} label="セッション3" sx={{ m: "0 auto" }} />
                  <Button
                    variant="contained"
                    fullWidth
                    sx={{
                      width: "50%",
                      backgroundColor: "primary.main",
                      color: "secondary.main",
                      m: "20px auto",
                      "&:hover": {
                        backgroundColor: "#FF3399",
                      },
                    }}
                    onClick={() => handleSave(data.dateTime)}
                  >
                    保存
                  </Button>
                </FormGroup>{" "}
              </AccordionDetails>
            </Accordion>
          ))}{" "}
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
