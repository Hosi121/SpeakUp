import { Box, Button, Container, Typography } from "@mui/material";
import CircularProgress from "@mui/material/CircularProgress";
import Memo from "./Memo";

const Waiting = () => {
  return (
    <Container sx={{ height: "100vh", display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center", padding: 3 }}>
      <Typography variant="h6" sx={{ mb: 2, width: "70%", maxWidth: "300px", backgroundColor: "secondary.main", color: "primary.main", borderRadius: 1, padding: 1, textAlign: "center" }}>
        6月30日21:00〜
      </Typography>

      <Box sx={{ width: "100%", maxWidth: 500, bgcolor: "background.paper", borderRadius: 1, p: 2, boxShadow: 3 }}>
        <Typography variant="subtitle1" gutterBottom component="div">
          持ち込みメモ
        </Typography>
        <Memo />
      </Box>

      <Button variant="contained" sx={{ mt: 4, color: "#000", bgcolor: "secondary.main", "&:hover": { bgcolor: "secondary.dark" }, width: "70%", maxWidth: "300px" }}>
        <CircularProgress size={24} sx={{ mr: 3 }} />
        セッション開始まで
        <br />
        しばらくお待ちください
      </Button>
    </Container>
  );
};

export default Waiting;
