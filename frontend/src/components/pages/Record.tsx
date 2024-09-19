import { Stack } from "@mui/system";
import { Box, Container, Typography } from "@mui/material";
import DescriptionIcon from "@mui/icons-material/Description";
import FolderOutlinedIcon from "@mui/icons-material/FolderOutlined";
import EmojiEventsIcon from "@mui/icons-material/EmojiEvents";
import EditNoteIcon from "@mui/icons-material/EditNote";
import { IconButton } from "../utils/IconButton";
import TopSection from "../utils/TopSection";
import { LibraryBooks } from "@mui/icons-material";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";

const RecordContainer = () => {
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
            <LibraryBooks sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }} />
            記録
          </Typography>
          <Box
            sx={{
              display: "flex",
              placeContent: "center",
              flexWrap: "wrap",
              justifyContent: "space-between",
            }}
          >
            <IconButton icon={<DescriptionIcon sx={{ fontSize: "60px" }} />} text="会話の記録" url="conversation_history" />
            <IconButton icon={<FolderOutlinedIcon sx={{ fontSize: "60px" }} />} text="履歴" url="session_history_friendlist" />
            <IconButton icon={<EmojiEventsIcon sx={{ fontSize: "60px" }} />} text="データ" url="stats" />
            <IconButton icon={<EditNoteIcon sx={{ fontSize: "60px" }} />} text="持ち込みメモ" url="memo" />
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};

export const Record = () => {
  return (
    <BottomNavigationTemplate value="record">
      <RecordContainer />
    </BottomNavigationTemplate>
  );
};
