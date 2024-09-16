import { useState } from "react";
import { Box, List, ListItem, ListItemText, Paper, TextField, Button, IconButton, Container } from "@mui/material";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { useNavigate } from "react-router-dom";
import TopSection from "../utils/TopSection"; // Import TopSection component
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";

const MessageContainer = () => {
  const [messages, setMessages] = useState<string[]>([]);
  const [inputMessage, setInputMessage] = useState<string>("");
  const navigate = useNavigate(); // Initialize useNavigate

  // Function to handle sending the message
  const handleSendMessage = () => {
    if (inputMessage.trim() !== "") {
      setMessages([...messages, inputMessage]);
      setInputMessage(""); // Clear the input field after sending
    }
  };

  // Function to go back to the previous page
  const handleGoBack = () => {
    navigate(-1); // Go back to the previous page
  };

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "calc(100vh - 70px)",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />
        {/* Back Button at the top-left corner */}
        <Box sx={{ position: "relative", display: "flex", justifyContent: "flex-start", pt: 2, pb: 2 }}>
          <IconButton onClick={handleGoBack}>
            <ArrowBackIcon sx={{ fontSize: 40 }} />
          </IconButton>
        </Box>

        {/* Chat Messages */}
        <Box sx={{ flexGrow: 1, overflow: "auto", p: 2 }}>
          <List>
            {/* Render each message from the state */}
            {messages.map((message, index) => (
              <ListItem key={index} sx={{ justifyContent: "flex-end" }}>
                <Paper sx={{ padding: "10px", backgroundColor: "#f0f0f0", maxWidth: "60%", wordWrap: "break-word" }}>
                  <ListItemText primary={message} />
                </Paper>
              </ListItem>
            ))}
          </List>
        </Box>
      </Container>
      {/* Message Input */}
      <Box sx={{ display: "flex", alignItems: "center", p: 2, borderTop: "1px solid #ccc" }}>
        <TextField variant="outlined" placeholder="メッセージを入力" fullWidth value={inputMessage} onChange={(e) => setInputMessage(e.target.value)} sx={{ mr: 2 }} />
        <Button variant="contained" color="primary" onClick={handleSendMessage}>
          送信
        </Button>
      </Box>
    </Container>
  );
};

export const Message = () => {
  return (
    <BottomNavigationTemplate value="other">
      <MessageContainer />
    </BottomNavigationTemplate>
  );
};
