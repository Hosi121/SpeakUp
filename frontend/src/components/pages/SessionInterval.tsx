import { Box, Container, Typography } from "@mui/material";
import { Stack } from "@mui/system";
import TopSection from "../utils/TopSection";
import { MemoInputField } from "../utils/MemoInputField";
import { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { SessionStepContext } from "../utils/SessionStepContextProvider";

export const SessionInterval = () => {
  const intervalTime = 60; // [s]
  const [countdown, setCountdown] = useState(intervalTime);
  const navigate = useNavigate();
  const [memoValue, setMemoValue] = useState("");
  const { learnedExpressions, setLearnedExpressions } =
    useContext(SessionStepContext);

  useEffect(() => {
    if (countdown > 0) {
      const timer = setTimeout(() => setCountdown(countdown - 1), 1000);
      return () => clearTimeout(timer);
    } else {
      navigate("/session");
      const newLearnedExpression = learnedExpressions + "\n" + memoValue;
      setLearnedExpressions(newLearnedExpression);
    }
  }, [countdown, navigate]);
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
        {learnedExpressions}
        <TopSection />
        <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
          <Typography
            variant="h3"
            sx={{
              mb: 1,
              textAlign: "left",
              color: "primary.main",
              fontFamily: "bangers",
            }}
          >
            GREAT!
          </Typography>
          <Typography textAlign="left" sx={{ mb: 2 }}>
            今のセッションで学んだことをメモしておこう！
          </Typography>
          <MemoInputField
            value={memoValue}
            setValue={setMemoValue}
            label="学んだ表現"
            height="30vh"
          />
          <Box sx={{ textAlign: "center", mt: 2 }}>
            <Typography
              variant="h6"
              component="p"
              fontWeight="bold"
              textAlign="left"
            >
              次のセッションまで...
            </Typography>
            <Typography
              variant="h1"
              component="p"
              color="primary"
              fontWeight="bold"
              sx={{ mt: 2, fontSize: "5rem", fontFamily: "bangers" }}
            >
              {countdown}
            </Typography>
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};
