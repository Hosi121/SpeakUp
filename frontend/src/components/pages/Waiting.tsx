import { Box, Container, LinearProgress, IconButton, Typography } from "@mui/material";
import TopSection from "../utils/TopSection";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import { MemoInputField } from "../utils/MemoInputField";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ArrowBack } from "@mui/icons-material";
import { CircularLineSpinner } from "../utils/CircularLineSpinner";

const WaitingContainer = () => {
  const navigate = useNavigate();
  const [memo, setMemo] = useState("");
  const scheduledTime = new Date("2024-09-20T07:51:00"); // mock
  const startedAt = new Date();
  const allWaitingTime = scheduledTime.getTime() - startedAt.getTime();
  const [remainTimeProgressRatio, setRemainTimeProgressRatio] = useState(0);

  const handleGoBack = () => {
    navigate(-1);
  };

  useEffect(() => {
    const intervalId = setInterval(() => {
      const now = new Date();
      const remainTime = scheduledTime.getTime() - now.getTime();
      const progressRaito = 1 - remainTime / allWaitingTime;
      const actualProgressRaito = Math.min(1, progressRaito);
      setRemainTimeProgressRatio(actualProgressRaito);
      if (remainTime <= 0) {
        clearInterval(intervalId);
        navigate("/miccheck");
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
          position: "relative",
        }}
      >
        <IconButton onClick={handleGoBack} sx={{ position: "absolute", top: 0, left: 3 }}>
          <ArrowBack sx={{ fontSize: 40 }} />
        </IconButton>
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
          <MemoInputField value={memo} onChange={setMemo} label="メモ入力" />
        </Box>

        <Box mt={3} width="100%">
          <Box mt={3}>
            <Typography>セッション開始まで</Typography>
            <Typography>もうしばらくお待ち下さい</Typography>
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
            }}
          >
            <LinearProgress
              variant="determinate"
              value={remainTimeProgressRatio * 100}
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
                },
              }}
            />
          </Box>
          <Box mt={2} display="flex" alignItems="center" ml="auto" sx={{ width: "fit-content" }}>
            <Typography pr={1}>Now loading...</Typography>
            <CircularLineSpinner size={18} strokeWidth={1.6} />
          </Box>
        </Box>
      </Container>
    </Container>
  );
};

export const Waiting = () => {
  return (
    <BottomNavigationTemplate value="session">
      <WaitingContainer />
    </BottomNavigationTemplate>
  );
};
