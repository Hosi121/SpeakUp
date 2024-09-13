import { useState } from "react";
import { Box, Typography, List, ListItem, ListItemText, IconButton, TextField, Button, Paper } from "@mui/material";
import SettingsIcon from '@mui/icons-material/Settings';
import { useNavigate } from "react-router-dom";
import { MainBottomNavigation } from "../utils/MainBottomNavigation"; 
import NotificationModal from "../utils/NotificationModal"; 

const Message = () => {
  const [messages, setMessages] = useState([]);
  const [inputMessage, setInputMessage] = useState('');
  const navigate = useNavigate(); 

  // Function to handle sending the message
  const handleSendMessage = () => {
    if (inputMessage.trim() !== "") {
      setMessages([...messages, inputMessage]);
      setInputMessage('');  // Clear the input field after sending
    }
  };

  // Function to navigate to settings page
  const handleNavigateSettings = () => {
    navigate("/settings");
  };

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', height: '100vh', justifyContent: 'space-between' }}>
      
      {/* Top Section */}
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', p: 2 }}>
        {/* Notification Icon */}
        <NotificationModal />

        {/* User Name */}
        <Typography variant="h6">Mike</Typography>

        {/* Settings Icon (navigate to /settings) */}
        <IconButton onClick={handleNavigateSettings}>
          <SettingsIcon sx={{ fontSize: 40 }} />
        </IconButton>
      </Box>

      {/* Chat Messages */}
      <Box sx={{ flexGrow: 1, overflow: 'auto', p: 2 }}>
        <List>
          {/* Render each message from the state */}
          {messages.map((message, index) => (
            <ListItem key={index} sx={{ justifyContent: 'flex-end' }}>
              <Paper sx={{ padding: '10px', backgroundColor: '#f0f0f0', maxWidth: '60%', wordWrap: 'break-word' }}>
                <ListItemText primary={message} />
              </Paper>
            </ListItem>
          ))}
        </List>
      </Box>

      {/* Message Input */}
      <Box sx={{ display: 'flex', alignItems: 'center', p: 2, borderTop: '1px solid #ccc' }}>
        <TextField
          variant="outlined"
          placeholder="メッセージを入力"
          fullWidth
          value={inputMessage}
          onChange={(e) => setInputMessage(e.target.value)}
          sx={{ mr: 2 }}
        />
        <Button 
          variant="contained" 
          color="primary" 
          onClick={handleSendMessage}
        >
          送信
        </Button>
      </Box>

      {/* Footer Navigation */}
      <MainBottomNavigation value="home" />
    </Box>
  );
};

export default Message;