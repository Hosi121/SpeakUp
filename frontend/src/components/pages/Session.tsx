import { HalfModal } from "../utils/HalfModal";
import { useState } from "react";
import { Typography, Box, List, ListItem, Paper, ListItemText, TextField, Button} from "@mui/material";
import HomeLogo from "../../assets/homeLogo";
import { Favorite, Person } from "@mui/icons-material";
import { SessionBottomNavigationTemplate } from "../templates/SessionBottomNavigationTemplate";
import SessionContainer from "../utils/SessionContainer";

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
        <Box sx={{ overflow: "auto", pt: 1, pb: 1, maxHeight: "30vh" }}>
          <List>
            {/* アシスタントからの初期メッセージ */}
            <ListItem sx={{ justifyContent: "flex-start" }}>
              <Box
                sx={{
                  width: "40px",
                  height: "40px",
                  borderRadius: "50%",
                  backgroundColor: "secondary.main",
                  mr: 2,
                  display: "grid",
                  placeContent: "center",
                }}
              >
                <HomeLogo style={{ width: "70%", height: "fit-content" }} />
              </Box>
              <Paper
                sx={{
                  padding: "5px",
                  backgroundColor: "background.default",
                  maxWidth: "60%",
                  wordWrap: "break-word",
                }}
              >
                <ListItemText primary="何かお困りですか？" />
              </Paper>
            </ListItem>

            {/* ユーザーのメッセージを表示 */}
            {messages.map((message, index) => (
              <ListItem key={index} sx={{ justifyContent: "flex-end" }}>
                <Paper
                  sx={{
                    padding: "5px",
                    backgroundColor: "#f0f0f0",
                    maxWidth: "60%",
                    wordWrap: "break-word",
                  }}
                >
                  <ListItemText primary={message} />
                </Paper>
              </ListItem>
            ))}
          </List>
        </Box>

        {/* メッセージ入力欄と送信ボタン */}
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            pb: 2,
            position: "fixed",
            bottom: 0,
            backgroundColor: "secondary.main",
          }}
        >
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
      </HalfModal>


    </SessionBottomNavigationTemplate>
  );
};

export default Session;