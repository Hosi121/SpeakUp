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
    <Container sx={{ padding: 3, paddingBottom: 4 }}>
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

        <Box mt={3} width="100%">
          <Box mt={3}>
            <Typography>
              セッション開始まで
            </Typography>
            <Typography>
              もうしばらくお待ち下さい
            </Typography>
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
          <Box mt={2} display="flex" alignItems="center" ml="auto"
            sx={{ width: "fit-content" }}>
            <Typography pr={1}>
              Now loading...
            </Typography>
            <CircularLineSpinner size={18} strokeWidth={1.6} />
          </Box>
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

const CircularLineSpinner = ({ size = 48, strokeWidth = 2, speed = 1 }) => {
  const lines = 12;
  const gapSize = 2;
  const radius = (size - strokeWidth) / 2;
  const center = size / 2;
  const color = "#000000";
  const transParentLines = 6;
  const calcStrokeColor = (num: number): string => {
    if (num < (lines - gapSize) - transParentLines) {
      return color;
    }
    const ratio = (num - ((lines - gapSize) - transParentLines)) / transParentLines;
    return `rgba(0, 0, 0, ${1 - ratio})`;
  }

  return (
    <div style={{ width: size, height: size }}>
      <svg
        width={size}
        height={size}
        viewBox={`0 0 ${size} ${size}`}
        xmlns="http://www.w3.org/2000/svg"
        style={{ animation: `spin ${1 / speed}s linear infinite` }}
      >
        {[...Array(lines - gapSize)].map((_, i) => {
          const angle = -(i * 360) / lines;
          const x1 = center + radius * Math.cos((angle * Math.PI) / 180);
          const y1 = center + radius * Math.sin((angle * Math.PI) / 180);
          const x2 = center + (radius - strokeWidth * 2) * Math.cos((angle * Math.PI) / 180);
          const y2 = center + (radius - strokeWidth * 2) * Math.sin((angle * Math.PI) / 180);
          return (
            <line
              key={i}
              x1={x1}
              y1={y1}
              x2={x2}
              y2={y2}
              stroke={calcStrokeColor(i)}
              strokeWidth={strokeWidth}
              strokeLinecap="round"
            />
          );
        })}
      </svg>
      <style>{`
        @keyframes spin {
          0% { transform: rotate(0deg); }
          100% { transform: rotate(360deg); }
        }
      `}</style>
    </div>
  );
};
