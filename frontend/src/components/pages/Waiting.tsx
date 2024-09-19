import { Box, Button, Container, Typography } from "@mui/material";
import CircularProgress from "@mui/material/CircularProgress";
import TopSection from "../utils/TopSection";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import { MemoInputField } from "../utils/MemoInputField";
import { useState } from "react";

const WaitingContainer = () => {
  const [memo, setMemo] = useState("");

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
          6月30日21:00〜
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

        <Button
          variant="contained"
          sx={{
            mt: 4,
            color: "#000",
            bgcolor: "secondary.main",
            "&:hover": { bgcolor: "secondary.dark" },
            width: "80%",
            maxWidth: "300px",
          }}
        >
          <CircularProgress size={24} sx={{ mr: 3 }} />
          セッション開始まで
          <br />
          しばらくお待ちください
        </Button>
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
