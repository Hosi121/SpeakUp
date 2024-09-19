import { Box, Button, Container, Typography } from "@mui/material";
import { Stack } from "@mui/system";
import TopSection from "../utils/TopSection";
import { MicIcon } from "../utils/MicIcon";
import AudioVisualizer from "../utils/AudioVisualizer";

export const MicCheck = () => {
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
          <Typography variant="h5" fontWeight="bolder" textAlign="left">
            セッション前に
            <br />
            マイクチェックをするよ！
          </Typography>
          <Button variant="contained" sx={{ width: "50%", p: 2, display: "block", m: "10% auto", fontSize: "1.2rem", borderRadius: 3 }}>
            準備OK!
          </Button>
          <Typography variant="h5" textAlign="left" fontWeight="bolder" color="primary.main" fontSize="1.7rem">
            "I'll&ensp;enjoy&ensp;speaking&ensp;English!!"
          </Typography>
          <Typography variant="h5" textAlign="right" fontWeight="bolder" color="primary.main" fontSize="1.7rem">
            と言おう！
          </Typography>
          <Box sx={{ width: "100%", display: "grid", placeContent: "center", mt: 4 }}>
            <AudioVisualizer />
            <MicIcon />
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};
