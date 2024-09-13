import React, { useState } from "react";
import { Avatar, Box, Typography, List, ListItem, ListItemAvatar, ListItemText, IconButton, TextField, Button } from "@mui/material";
import { Star } from "@mui/icons-material";
import SettingsIcon from '@mui/icons-material/Settings';
import { MainBottomNavigation } from "../utils/MainBottomNavigation";  // Assuming you have this component already
import NotificationModal from "../utils/NotificationModal";  // Assuming you have this component already

const Message = () => {
  // State to store messages and the current input
  const [messages, setMessages] = useState([]);
  const [inputMessage, setInputMessage] = useState('');

  // Function to handle sending the message
  const handleSendMessage = () => {
    if (inputMessage.trim() !== "") {
      setMessages([...messages, inputMessage]);
      setInputMessage('');  // Clear the input field after sending
    }
  };

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', height: '100vh', justifyContent: 'space-between' }}>
      
      {/* Top Section */}
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', p: 2 }}>
        {/* Notification Icon */}
        <NotificationModal />

        {/* User Avatar and Name */}
        <Box sx={{ display: 'flex', alignItems: 'center' }}>
          <Avatar sx={{ width: 64, height: 64, bgcolor: '#FFD700' }}>
            <Star sx={{ fontSize: 40, color: '#FF007F' }} />
          </Avatar>
          <Typography variant="h6" sx={{ ml: 2 }}>Mike</Typography>
        </Box>

        {/* Settings Icon */}
        <IconButton>
          <SettingsIcon sx={{ fontSize: 40 }} />
        </IconButton>
      </Box>

      {/* Chat Messages */}
      <Box sx={{ flexGrow: 1, overflow: 'auto', p: 2 }}>
        <List>
          {/* Render each message from the state */}
          {messages.map((message, index) => (
            <ListItem key={index}>
              <ListItemAvatar>
                <Avatar sx={{ bgcolor: '#FFD700' }}>
                  <Star sx={{ color: '#FF007F' }} />
                </Avatar>
              </ListItemAvatar>
              <ListItemText primary={message} />
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