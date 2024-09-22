import { useState, useEffect } from "react";
import { Box, Container, Typography, CircularProgress } from "@mui/material";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import TopSection from "../utils/TopSection";
import TopWaves from "../utils/TopWaves";
import HomeLogo from "../../assets/homeLogo";
import { IconButton } from "../utils/IconButton";
import { LibraryBooks, Mic } from "@mui/icons-material";
import api from "../../services/api";

type SessionData = {
  dateTime: string;
  sessions: number[];
  theme: string;
};

type EventData = {
  id: number;
  event_start: string;
  event_end: string;
  theme_id: number;
  theme: {
    theme_text: string;
    topic1: string;
    topic2: string;
    topic3: string;
  };
};

const HomeContainer = () => {
  const [, setSessionData] = useState<SessionData[]>([]);
  const [eventData, setEventData] = useState<EventData[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // APIからイベントデータを取得
        const response = await api.get("/events");
        setEventData(response.data);

        // 既存のJSONデータを取得（実際のアプリケーションでは、このデータもAPIから取得するかもしれません）
        const sessionResponse = await import("../../mock/sessions.json");
        setSessionData([
          {
            ...sessionResponse.default,
            theme: "",
            dateTime: "",
            sessions: [],
          },
        ]);

        setLoading(false);
      } catch (error) {
        console.error("Failed to fetch data", error);
        setError("データの取得に失敗しました");
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const getNextEvent = () => {
    if (!eventData || eventData.length === 0) return null;

    const now = new Date();
    return (
      eventData.find((event) => new Date(event.event_start) > now) ||
      eventData[0]
    );
  };

  const nextEvent = getNextEvent();

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "calc(100vh - 70px)",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />
        <Container sx={{ mb: 3 }}>
          <TopWaves isFlipped={false} />
          <Box
            sx={{
              margin: "0 calc(50% - 50vw)",
              width: "100vw",
              height: "calc(15vh + 1px)",
              marginTop: "-1px",
              backgroundColor: "secondary.main",
              display: "grid",
              placeItems: "center",
            }}
          >
            <HomeLogo style={{ width: "70%" }} />
          </Box>
          <TopWaves isFlipped={true} />
        </Container>
        <Box
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
        >
          <h3>直近の参加予定</h3>
          {loading ? (
            <CircularProgress />
          ) : error ? (
            <Typography color="error">{error}</Typography>
          ) : nextEvent ? (
            <>
              <Typography
                sx={{
                  mt: 1,
                  mb: 1,
                  color: "primary.main",
                  fontSize: "1.3rem",
                  fontWeight: "bolder",
                  textAlign: "center",
                }}
              >
                {new Date(nextEvent.event_start).toLocaleString("ja-JP", {
                  month: "long",
                  day: "numeric",
                  hour: "2-digit",
                  minute: "2-digit",
                })}
              </Typography>
              <Typography
                sx={{
                  color: "primary.main",
                  fontSize: "1rem",
                  fontWeight: "bolder",
                  textAlign: "center",
                }}
              >
                テーマ: {nextEvent.theme.theme_text}
              </Typography>
              <Typography
                sx={{
                  color: "primary.main",
                  fontSize: "0.9rem",
                  textAlign: "center",
                  mt: 1,
                }}
              >
                トピック1: {nextEvent.theme.topic1}
              </Typography>
              <Typography
                sx={{
                  color: "primary.main",
                  fontSize: "0.9rem",
                  textAlign: "center",
                }}
              >
                トピック2: {nextEvent.theme.topic2}
              </Typography>
              <Typography
                sx={{
                  color: "primary.main",
                  fontSize: "0.9rem",
                  textAlign: "center",
                }}
              >
                トピック3: {nextEvent.theme.topic3}
              </Typography>
            </>
          ) : (
            <Typography>予定されているイベントはありません</Typography>
          )}
        </Box>

        <Box
          sx={{
            display: "flex",
            placeContent: "center",
            flexWrap: "wrap",
            justifyContent: "space-between",
          }}
        >
          <IconButton
            icon={<LibraryBooks sx={{ fontSize: "60px" }} />}
            text="記録"
            url="record"
          />
          <IconButton
            icon={<Mic sx={{ fontSize: "60px" }} />}
            text="セッション"
            url="sessionlist"
          />
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
