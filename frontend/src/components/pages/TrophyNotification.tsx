import { Box, Typography, Container } from "@mui/material";
import EmojiEventsIcon from "@mui/icons-material/EmojiEvents";

const TrophyNotification = () => {
  return (
    <Container sx={{ p: 0, width: "100%", height: "100vh", display: "grid", placeItems: "center" }}>
      <Container
        sx={{
          p: 0,
          width: "100%",
          height: "60vh",
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <Box
          sx={{
            width: "100%",
            bgcolor: "white",
            textAlign: "center",
            py: 3,
            mt: 4,
          }}
        >
          <Typography variant="h4" component="h1" sx={{ color: "primary.main", fontWeight: "bold" }}>
            トロフィー獲得！
          </Typography>
        </Box>

        <EmojiEventsIcon sx={{ fontSize: 200, color: "black" }} />

        <Typography variant="body1" sx={{ mb: 4, fontWeight: "bold" }}>
          初めてのセッションに参加
        </Typography>
      </Container>
    </Container>
  );
};

export default TrophyNotification;
