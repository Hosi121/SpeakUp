import { Box, Container, Typography } from "@mui/material";
import { Stack } from "@mui/system";
import TopSection from "../utils/TopSection";
import { MemoInputField } from "../utils/MemoInputField";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export const SessionInterval = () => {
  const [countdown, setCountdown] = useState(60);
  const navigate = useNavigate();

  useEffect(() => {
    if (countdown > 0) {
      const timer = setTimeout(() => setCountdown(countdown - 1), 1000);
      return () => clearTimeout(timer);
    } else {
      navigate("/session");
    }
  }, [countdown]);
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
          <Typography variant="h3" sx={{ mb: 1, textAlign: "left", color: "primary.main", fontFamily: "bangers" }}>
            GREAT!
          </Typography>
          <Typography textAlign="left" sx={{ mb: 2 }}>
            今のセッションで学んだことをメモしておこう！
          </Typography>
          <MemoInputField value={""} onChange={() => {}} label="学んだ表現" height="30vh" />
          <Box sx={{ textAlign: "center", mt: 2 }}>
            <Typography variant="h6" component="p" fontWeight="bold" textAlign="left">
              次のセッションまで...
            </Typography>
            <Typography variant="h1" component="p" color="primary" fontWeight="bold" sx={{ mt: 2, fontSize: "5rem", fontFamily: "bangers" }}>
              {countdown}
            </Typography>
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};
