import React from "react";
import { List, ListItem, ListItemAvatar, Avatar, ListItemText, Typography, Button, Box } from "@mui/material";

interface History {
  avatar: string;
  user: string;
  theme: string;
  date: string;
  rank: number;
  friedstate: string;
}

interface SessionHistoryProps {
  history: History[];
}

const SessionHistory = ({ history }: SessionHistoryProps) => {
  return (
    <List>
      {history.map((data: History) => (
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
                <React.Fragment>
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
                </React.Fragment>
              }
            />
          </Box>
          {data.friedstate === "friend" && (
            <Box>
              <Button variant="outlined" color="primary" style={{ marginRight: 8 }}>
                フレンド申請済
              </Button>
              <Button variant="contained" color="primary">
                メッセージ
              </Button>
            </Box>
          )}
          {data.friedstate === "pending" && (
            <Box>
              <Button variant="outlined" color="primary" style={{ marginRight: 8 }}>
                フレンド申請済
              </Button>
              <Button variant="outlined" color="primary">
                メッセージ
              </Button>
            </Box>
          )}
          {data.friedstate === "unapplied" && (
            <Box>
              <Button variant="contained" color="primary" style={{ marginRight: 8 }}>
                フレンド申請
              </Button>
              <Button variant="outlined" color="primary">
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
