import { Stack } from "@mui/system";
import { Box, Typography } from "@mui/material";
import DescriptionIcon from "@mui/icons-material/Description";
import FolderOutlinedIcon from "@mui/icons-material/FolderOutlined";
import EmojiEventsIcon from "@mui/icons-material/EmojiEvents";
import EditNoteIcon from "@mui/icons-material/EditNote";
import { IconButton } from "../utils/IconButton";

export const Record = () => {
  return (
    <Stack sx={{ margin: "0 auto", width: "90%" }}>
      <Typography align="left" sx={{ marginBottom: "30px", fontSize: "2rem", color: "#da0063" }}>
        記録
      </Typography>
      <Box sx={{ display: "flex", placeContent: "center", flexWrap: "wrap", justifyContent: "space-between" }}>
        <IconButton icon={<DescriptionIcon sx={{ fontSize: "60px" }} />} text="会話の記録" url="conversationrecords" />
        <IconButton icon={<FolderOutlinedIcon sx={{ fontSize: "60px" }} />} text="履歴" url="history" />
        <IconButton icon={<EmojiEventsIcon sx={{ fontSize: "60px" }} />} text="データ" url="stats" />
        <IconButton icon={<EditNoteIcon sx={{ fontSize: "60px" }} />} text="持ち込みメモ" url="memo" />
      </Box>
    </Stack>
  );
};
