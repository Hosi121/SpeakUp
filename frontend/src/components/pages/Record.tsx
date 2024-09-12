import { Stack } from "@mui/system";
import { Box, Typography } from "@mui/material";
import DescriptionIcon from "@mui/icons-material/Description";
import FolderOutlinedIcon from "@mui/icons-material/FolderOutlined";
import EmojiEventsIcon from "@mui/icons-material/EmojiEvents";
import EditNoteIcon from "@mui/icons-material/EditNote";
import { AnchorButton } from "../utils/AnchorButton";

export const Record = () => {
  return (
    <Stack sx={{ margin: "0 auto", width: "90%" }}>
      <Typography align="left" sx={{ marginBottom: "30px", fontSize: "2rem", color: "#da0063" }}>
        記録
      </Typography>
      <Box sx={{ display: "flex", placeContent: "center", flexWrap: "wrap", justifyContent: "space-between" }}>
        <AnchorButton icon={<DescriptionIcon sx={{ fontSize: "60px" }} />} text="会話の記録" />
        <AnchorButton icon={<FolderOutlinedIcon sx={{ fontSize: "60px" }} />} text="履歴" />
        <AnchorButton icon={<EmojiEventsIcon sx={{ fontSize: "60px" }} />} text="データ" />
        <AnchorButton icon={<EditNoteIcon sx={{ fontSize: "60px" }} />} text="持ち込みメモ" />
      </Box>
    </Stack>
  );
};
