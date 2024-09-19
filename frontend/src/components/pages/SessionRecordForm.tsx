import React, { useState } from "react";
import { Box, Typography, TextField, Button, Paper, Container, Stack } from "@mui/material";
import { styled } from "@mui/system";
import { useNavigate } from "react-router-dom";
import TopSection from "../utils/TopSection";

const DashedBox = styled(Paper)({
  border: "2px dashed #FFD700",
  padding: "20px",
  marginBottom: "20px",
});
const SessionRecordForm: React.FC = () => {
  const [satisfaction, setSatisfaction] = useState<string>("");
  const [thoughts, setThoughts] = useState<string>("");
  const [learnedExpressions, setLearnedExpressions] = useState<string>("");
  const navigate = useNavigate();

  const handleNext = () => {
    navigate("/friendrequest");
  };

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <Container sx={{ pt: 3, pb: 3 }}>
        <TopSection />
        <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
          <Typography variant="h3" sx={{ mb: 1, textAlign: "left", color: "primary.main", fontFamily: "bangers" }}>
            GOOD JOB!
          </Typography>
          <Typography textAlign="left" sx={{ mb: 2 }}>
            記録を残してみませんか？
          </Typography>
          <DashedBox elevation={0}>
            <Typography variant="body1">5月15日</Typography>
            <Typography variant="body1">参加したセッション数: 3</Typography>
            <TextField fullWidth label="満足度 (%)" value={satisfaction} onChange={(e) => setSatisfaction(e.target.value)} margin="normal" />
            <TextField fullWidth label="感想" multiline rows={3} value={thoughts} onChange={(e) => setThoughts(e.target.value)} margin="normal" />
            <TextField fullWidth label="学んだ表現" multiline rows={3} value={learnedExpressions} onChange={(e) => setLearnedExpressions(e.target.value)} margin="normal" />
          </DashedBox>

          <Typography variant="h6" align="center" gutterBottom fontWeight="bolder">
            次に参加予定のセッション
          </Typography>

          <Box sx={{ backgroundColor: "secondary.main", padding: "10px", marginBottom: "20px", color: "primary.main", borderRadius: 2 }}>
            <Typography variant="body1" align="center" sx={{ fontWeight: "bolder" }}>
              7月17日20:00〜
            </Typography>
          </Box>

          <Button variant="contained" fullWidth onClick={handleNext}>
            次へ
          </Button>
        </Stack>
      </Container>
    </Container>
  );
};

export default SessionRecordForm;
