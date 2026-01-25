import { List, ListItem, ListItemAvatar, Avatar, ListItemText, Typography, Button, Box } from "@mui/material";
import type { SessionHistoryItem } from "../../types/types";

interface SessionHistoryProps {
  history: SessionHistoryItem[];
}

const SessionHistory = ({ history }: SessionHistoryProps) => {
  return (
    <List>
      {history.map((data) => (
        <ListItem alignItems="flex-start" sx={{ mb: 2, display: "flex", flexDirection: "column" }}>
          <Box sx={{ mb: 1, display: "flex", flexDirection: "row", width: "100%", flexWrap: "wrap" }}>
            <ListItemAvatar sx={{ width: "30%" }}>
              <Avatar alt={data.user} src={data.avatar} />
            </ListItemAvatar>
            <ListItemText
              sx={{ width: "70%" }}
              primary={
                <Box display="flex" alignItems="center">
                  <Typography component="span" variant="body1" color="text.primary">
                    {data.user}
                  </Typography>
                </Box>
              }
              secondary={
                <>
                  <Typography component="span" variant="body2" color="text.primary">
                    話したテーマ: {data.theme}
                  </Typography>
                  <br />
                  <Typography component="span" variant="body2" color="text.secondary">
                    {`最終日: ${data.date}`}
                  </Typography>
                  <br />
                  <Typography component="span" variant="body2" color="text.secondary">
                    {`ランク${data.rank}`}
                  </Typography>
                </>
              }
            />
          </Box>
          {data.friendState === "friend" && (
            <Box sx={{ display: "flex", justifyContent: "space-between", width: "100%" }}>
              <Button sx={{ width: "55%" }} variant="contained" disabled color="primary">
                フレンド申請済
              </Button>
              <Button sx={{ width: "43%" }} variant="contained" color="primary">
                メッセージ
              </Button>
            </Box>
          )}
          {data.friendState === "pending" && (
            <Box sx={{ display: "flex", justifyContent: "space-between", width: "100%" }}>
              <Button sx={{ width: "55%" }} variant="contained" disabled color="primary">
                フレンド申請済
              </Button>
              <Button sx={{ width: "43%" }} variant="contained" disabled color="primary">
                メッセージ
              </Button>
            </Box>
          )}
          {data.friendState === "unapplied" && (
            <Box sx={{ display: "flex", justifyContent: "space-between", width: "100%" }}>
              <Button sx={{ width: "55%" }} variant="contained" color="primary">
                フレンド申請
              </Button>
              <Button sx={{ width: "43%" }} variant="contained" disabled color="primary">
                メッセージ
              </Button>
            </Box>
          )}
        </ListItem>
      ))}
    </List>
  );
};

export default SessionHistory;
