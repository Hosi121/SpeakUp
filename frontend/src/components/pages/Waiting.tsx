import { Box, Container, LinearProgress, Typography } from "@mui/material";
import TopSection from "../utils/TopSection";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import { MemoInputField } from "../utils/MemoInputField";
import { useEffect, useState } from "react";

const WaitingContainer = () => {
  const [memo, setMemo] = useState("");
  const scheduledTime = new Date("2024-09-20T06:47:00"); // mock
  const startedAt = new Date();
  const allWaitingTime = scheduledTime.getTime() - startedAt.getTime();
  const [remainTimeProgressRaito, setRemainTimeProgressRaito] = useState(0);

  useEffect(() => {
    const intervalId = setInterval(() => {
      const now = new Date();
      const remainTime = scheduledTime.getTime() - now.getTime();
      const progressRaito = 1 - remainTime / allWaitingTime;
      const actualProgressRaito = Math.min(1, progressRaito);
      setRemainTimeProgressRaito(actualProgressRaito);
      if (remainTime <= 0) {
        clearInterval(intervalId);
      }
    }, 800);
  }, []);

  const scheduledTimeStr = `${scheduledTime.getMonth() + 1}月${scheduledTime.getDate()}日${scheduledTime.getHours()}:${scheduledTime.getMinutes().toString().padStart(2, "0")}`;

  return (
    <Container sx={{ height: "100vh", padding: 3 }}>
      <TopSection />
      <Container
        sx={{
          mt: "5%",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <Typography
          variant="h6"
          sx={{
            mb: 2,
            width: "70%",
            maxWidth: "300px",
            backgroundColor: "secondary.main",
            color: "primary.main",
            borderRadius: 1,
            padding: 1,
            textAlign: "center",
          }}
        >
          {scheduledTimeStr}〜
        </Typography>

        <Box
          sx={{
            width: "100%",
            maxWidth: 500,
            bgcolor: "background.paper",
            borderRadius: 1,
            p: 2,
            boxShadow: 3,
          }}
        >
          <Typography variant="subtitle1" gutterBottom component="div">
            持ち込みメモ
          </Typography>
          <MemoInputField
            value={memo}
            onChange={setMemo}
            label="メモ入力"
          />
        </Box>

        <Box
          sx={{
            width: "100%",
            mt: 2,
            height: "22px",
            border: "solid black 2px",
            borderRadius: "5px",
            backgroundColor: "secondary.main",
            padding: 0,
            position: "relative",
          }}>
          <LinearProgress
            variant="determinate"
            value={remainTimeProgressRaito * 100}
            sx={{
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              width: "98%",
              borderRadius: "3px",
              height: "66%",
              backgroundColor: "transparent",
              span: {
                transition: "transform 0.8s linear",
                WebkitTransform: "-webkit-transform 0.8s linear",
              }
            }} />
        </Box>
      </Container>
    </Container >
  );
};

export const Waiting = () => {
  return (
    <BottomNavigationTemplate value="session">
      <WaitingContainer />
    </BottomNavigationTemplate>
  );
};
