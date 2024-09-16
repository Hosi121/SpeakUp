import { Container, Stack, Typography } from "@mui/material";
import EditNoteIcon from "@mui/icons-material/EditNote";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import TopSection from "../utils/TopSection";
import { MemoInputField } from "../utils/MemoInputField";

const MemoContainer = () => {
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
          <Typography variant="h4" sx={{ mb: 4, fontWeight: "bold", textAlign: "left" }}>
            <EditNoteIcon sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }} />
            持ち込みメモ
          </Typography>
          <MemoInputField />
        </Stack>
      </Container>
    </Container>
  );
};

export const Memo = () => {
  return (
    <BottomNavigationTemplate value="other">
      <MemoContainer />
    </BottomNavigationTemplate>
  );
};
