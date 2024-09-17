import { useState, useEffect } from "react";
import {
  Box,
  List,
  ListItem,
  ListItemText,
  Paper,
  TextField,
  Button,
  IconButton,
  Container,
  Typography,
} from "@mui/material";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { useNavigate, useParams } from "react-router-dom";
import TopSection from "../utils/TopSection";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import api from "../../services/api";

const MessageContainer = () => {
  const { friendname } = useParams(); // URLパラメータからfriendnameを取得
  const [friendInfo, setFriendInfo] = useState(null);
  const [messages, setMessages] = useState([]);
  const [inputMessage, setInputMessage] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    // 友達の情報を取得
    api
      .get(`/friend/${friendname}`)
      .then((response) => {
        setFriendInfo(response.data);
      })
      .catch((error) => {
        console.error("Failed to fetch friend info", error);
      });
  }, [friendname]);

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

  if (!friendInfo) {
    return <div>Loading...</div>; // データがまだロードされていない場合の処理
  }

  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        height: "calc(100vh - 70px)",
        boxSizing: "border-box",
      }}
    >
      <Container sx={{ pt: 3 }}>
        <TopSection />
        {/* Back Button at the top-left corner */}
        <Box
          sx={{
            position: "relative",
            display: "flex",
            justifyContent: "flex-start",
            pt: 2,
            pb: 2,
          }}
        >
          <IconButton onClick={handleGoBack}>
            <ArrowBackIcon sx={{ fontSize: 40 }} />
          </IconButton>
          <Container
            sx={{
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -25%)",
              width: "fit-content",
              textAlign: "center",
            }}
          >
            <Box
              sx={{
                width: "80px",
                height: "80px",
                borderRadius: "50%",
                backgroundColor: "secondary.main",
                backgroundImage: `url(${friendInfo.avatar})`,
                backgroundSize: "cover",
                backgroundPosition: "center",
                margin: "0 auto",
              }}
            ></Box>
            <Typography sx={{ fontSize: "1.5rem" }}>
              {friendInfo.username}
            </Typography>
          </Container>
        </Box>

        {/* Chat Messages */}
        <Box
          sx={{
            mt: 4,
            flexGrow: 1,
            overflow: "auto",
            pt: 2,
            pb: 2,
            height: "55vh",
          }}
        >
          <List>
            {/* 固定の受信メッセージ */}
            <ListItem sx={{ justifyContent: "flex-start" }}>
              <Box
                sx={{
                  width: "40px",
                  height: "40px",
                  borderRadius: "50%",
                  backgroundColor: "secondary.main",
                  backgroundImage: `url(${friendInfo.avatar})`,
                  backgroundSize: "cover",
                  backgroundPosition: "center",
                  mr: 2,
                }}
              ></Box>
              <Paper
                sx={{
                  padding: "10px",
                  backgroundColor: "#ccc",
                  maxWidth: "60%",
                  wordWrap: "break-word",
                }}
              >
                <ListItemText primary="こんにちは" />
              </Paper>
            </ListItem>

            {/* ユーザーからの送信メッセージ */}
            {messages.map((message, index) => (
              <ListItem key={index} sx={{ justifyContent: "flex-end" }}>
                <Paper
                  sx={{
                    padding: "10px",
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
      </Container>
      {/* Message Input */}
      <Box
        sx={{
          display: "flex",
          alignItems: "center",
          p: 2,
          borderTop: "1px solid #ccc",
        }}
      >
        <TextField
          variant="outlined"
          placeholder="メッセージを入力"
          fullWidth
          value={inputMessage}
          onChange={(e) => setInputMessage(e.target.value)}
          sx={{ mr: 2 }}
        />
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
