import { useState, useEffect } from 'react';
import { Box, Button, Typography, IconButton, Avatar, List, ListItem, ListItemAvatar, ListItemText, Dialog, DialogTitle, DialogContent } from '@mui/material';
import SettingsIcon from '@mui/icons-material/Settings';
import NotificationsIcon from '@mui/icons-material/Notifications';
import notifications from '../../mock/notifications.json'; // Import the JSON file directly

const NotificationModal = () => {
  // State to control the modal open/close
  const [open, setOpen] = useState(false);
  const [notificationsData, setNotificationsData] = useState([]);

  // Function to open the modal
  const handleOpen = () => {
    setOpen(true);
  };

  // Function to close the modal
  const handleClose = () => {
    setOpen(false);
  };

  // Load notifications from imported JSON
  useEffect(() => {
    setNotificationsData(notifications); // Set the imported JSON data as state
  }, []);

  return (
    <div>
      {/* Button or Icon to trigger the modal */}
      <IconButton onClick={handleOpen}>
        <NotificationsIcon sx={{ fontSize: 40, color: '#1a1a1a' }} />
      </IconButton>

      {/* Dialog (Modal) */}
      <Dialog open={open} onClose={handleClose} fullWidth maxWidth="xs">
        <DialogTitle>
          {/* Header with Notification Icon and Settings */}
          <Box
            sx={{
              display: 'flex',
              justifyContent: 'space-between',
              width: '100%',
            }}
          >
            <Box sx={{ display: 'flex', alignItems: 'center' }}>
              <NotificationsIcon sx={{ fontSize: 40, color: '#1a1a1a' }} />
              <Typography variant="h6" sx={{ ml: 1 }}>
                通知
              </Typography>
            </Box>
            <IconButton>
              <SettingsIcon sx={{ fontSize: 32, color: '#1a1a1a' }} />
            </IconButton>
          </Box>
        </DialogTitle>

        <DialogContent>
          {/* Notifications List */}
          <List sx={{ width: '100%' }}>
            {notificationsData.map((notification) => (
              <ListItem key={notification.id} sx={{ mb: 2 }}>
                <ListItemAvatar>
                  <Avatar src={notification.profileIcon} alt={`${notification.user} icon`} />
                </ListItemAvatar>
                <ListItemText
                  primary={
                    <Box sx={{ display: 'flex', alignItems: 'center' }}>
                      <Typography variant="body1" sx={{ color: '#FF007F', mr: 1 }}>
                        {notification.user}
                      </Typography>
                      <Typography variant="body2" sx={{ color: '#1a1a1a' }}>
                        {notification.message}
                      </Typography>
                    </Box>
                  }
                  secondary={
                    <Typography variant="caption" sx={{ color: '#1a1a1a' }}>
                      {notification.time}
                    </Typography>
                  }
                />
                {notification.type === 'friendRequest' && (
                  <Box>
                    <Button variant="contained" size="small" sx={{ backgroundColor: '#FF007F', color: 'white', mr: 1 }}>
                      承認
                    </Button>
                    <Button variant="outlined" size="small" sx={{ color: '#1a1a1a' }}>
                      拒否
                    </Button>
                  </Box>
                )}
              </ListItem>
            ))}
          </List>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default NotificationModal;
