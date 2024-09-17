import { Card, CardContent, Typography, Avatar, Grid, Box, Container, Stack, List, ListItem, Paper, ListItemText, TextField, Button } from "@mui/material";
import { Favorite, Person } from "@mui/icons-material";
import TopSection from "../utils/TopSection";
import { SessionBottomNavigationTemplate } from "../templates/SessionBottomNavigationTemplate";
import SessionCountDownModal from "../utils/SessionCountDownModal";
import { HalfModal } from "../utils/HalfModal";
import { useState } from "react";
import HomeLogo from "../../assets/homeLogo";

const SessionContainer = ({ theme, users }: { theme: string; users: { name: string; icon: JSX.Element; description: string }[] }) => {
  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "100%",
        position: "relative",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />

        <Stack sx={{ margin: "30px auto 0", width: "100%" }}>
          <Box
            sx={{
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              zIndex: 1000,
            }}
          >
            <SessionCountDownModal />
          </Box>
          <Typography variant="h5" align="center" gutterBottom color="primary.main" fontWeight="bolder">
            テーマ: {theme}
          </Typography>
          <Box sx={{ display: "flex", flexDirection: "column", justifyContent: "center", mb: 2 }}>
            {users.map((user, index) => (
              <Card key={index} sx={{ bgcolor: "secondary.main", mb: 2, width: "100%", height: "25vh", borderRadius: 5, border: "6px solid #eee", boxSizing: "border-box", p: 1, display: "grid", placeContent: "center" }}>
                <CardContent>
                  <Grid container spacing={2} justifyContent="center">
                    <Grid item>
                      <Avatar sx={{ bgcolor: "#eee", width: "80px", height: "80px" }}>{user.icon}</Avatar>
                      <Typography variant="h6" align="center" sx={{ mt: 1 }}>
                        {user.name}
                      </Typography>
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            ))}
          </Box>
        </Stack>
      </Container>
    </Container>
  );
};

const users = [
  { name: "User1", icon: <Person />, description: "英語" },
  { name: "User2", icon: <Favorite />, description: "苗字" },
];

const theme = "好きな言葉";

export const Session = () => {
  const [memoOpen, setMemoOpen] = useState(false);
  const [assistantOpen, setAssistantOpen] = useState(false);
  const [messages, setMessages] = useState<string[]>([]);
  const [inputMessage, setInputMessage] = useState<string>("");
  const handleMemoClose = () => setMemoOpen(false);
  const handleAssistantClose = () => setAssistantOpen(false);
  const handleSendMessage = () => {
    if (inputMessage.trim() !== "") {
      setMessages([...messages, inputMessage]);
      setInputMessage(""); // Clear the input field after sending
    }
  };
  const memo = "I'm Hanako. Please call me Hanako.";

  return (
    <SessionBottomNavigationTemplate value="other" isMute={false} setMemoOpen={setMemoOpen} setAssistantOpen={setAssistantOpen}>
      <SessionContainer theme={theme} users={users} />
      <HalfModal open={memoOpen} handleClose={handleMemoClose} title="持ち込みメモ">
        <Typography variant="body1">{memo}</Typography>
      </HalfModal>

      <HalfModal open={assistantOpen} handleClose={handleAssistantClose} title="アシスタント">
        <Box>
          {" "}
          <Box sx={{ overflow: "auto", pt: 1, pb: 1, maxHeight: "30vh" }}>
            <List>
              <ListItem sx={{ justifyContent: "flex-start" }}>
                <Box sx={{ width: "40px", height: "40px", borderRadius: "50%", backgroundColor: "secondary.main", mr: 2, display: "grid", placeContent: "center" }}>
                  <HomeLogo style={{ width: "70%", height: "fit-content" }} />
                </Box>
                <Paper sx={{ padding: "5px", backgroundColor: "background.default", maxWidth: "60%", wordWrap: "break-word" }}>
                  <ListItemText primary="何かお困りですか？" />
                </Paper>
              </ListItem>

              {messages.map((message, index) => (
                <ListItem key={index} sx={{ justifyContent: "flex-end" }}>
                  <Paper sx={{ padding: "5px", backgroundColor: "#f0f0f0", maxWidth: "60%", wordWrap: "break-word" }}>
                    <ListItemText primary={message} />
                  </Paper>
                </ListItem>
              ))}
            </List>
          </Box>
          <Box sx={{ display: "flex", alignItems: "center", pb: 2, position: "fixed", bottom: 0, backgroundColor: "secondary.main" }}>
            <TextField
              variant="outlined"
              placeholder="メッセージを入力"
              fullWidth
              value={inputMessage}
              onChange={(e) => setInputMessage(e.target.value)}
              sx={{ mr: 2 }}
              InputProps={{
                style: {
                  height: "40px",
                },
              }}
            />
            <Button variant="contained" color="primary" onClick={handleSendMessage}>
              送信
            </Button>
          </Box>
        </Box>
      </HalfModal>
    </SessionBottomNavigationTemplate>
  );
};
