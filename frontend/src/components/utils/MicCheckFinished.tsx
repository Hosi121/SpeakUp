import { Box, Button, Stack, Typography } from "@mui/material";
import { MicIcon } from "./MicIcon";
import { useNavigate } from "react-router-dom";

export const MicCheckFinished = () => {
  const navigate = useNavigate();

  return (
    <Stack sx={{ margin: "30px auto 0", width: "90%" }}>
      <Typography variant="h3" sx={{ mb: 1, textAlign: "left", color: "primary.main", fontFamily: "bangers" }}>
        NICE!
      </Typography>
      <Typography variant="h5" fontWeight="bolder" textAlign="center" sx={{ mt: "10%", mb: "10%" }}>
        マイクチェック完了！
      </Typography>
      <Typography variant="h6" fontWeight="bolder" sx={{ fontSize: "1.1rem" }}>
        あと少しでセッションが始まるよ！
      </Typography>
      <Typography variant="h6" fontWeight="bolder" sx={{ fontSize: "1.1rem" }}>
        準備はいいかな？
      </Typography>
      <Button variant="contained" sx={{ width: "50%", p: 2, display: "block", m: "10% auto", fontSize: "1.2rem", borderRadius: 3 }} onClick={() => navigate("/session")}>
        準備OK!
      </Button>
      <Box sx={{ width: "100%", display: "grid", placeContent: "center", mt: 4 }}>
        <MicIcon />
      </Box>
    </Stack>
  );
};
