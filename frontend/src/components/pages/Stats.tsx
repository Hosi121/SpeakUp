import React, { useEffect, useState } from 'react';
import { Box, Typography, Paper, List, ListItem, ListItemIcon, ListItemText } from '@mui/material';
import NotificationsIcon from '@mui/icons-material/Notifications';
import SettingsIcon from '@mui/icons-material/Settings';
import EmojiEventsIcon from '@mui/icons-material/EmojiEvents';
import { MainBottomNavigation } from '../utils/MainBottomNavigation';
import trophyData from '../../assets/trophy.json';

const Stats = () => {
  const [trophies, setTrophies] = useState([]);

  // Fetch the trophy data from trophy.json
  useEffect(() => {
    setTrophies(trophyData.trophies);
  }, []);

  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', height: '100vh', justifyContent: 'space-between', p: 2 }}>
      {/* Top section with Notification and Settings */}
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
        <NotificationsIcon sx={{ fontSize: 40, color: '#1E00A1' }} />
        <SettingsIcon sx={{ fontSize: 40 }} />
      </Box>

      {/* Participation Data */}
      <Box sx={{ mb: 3 }}>
        <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
          <EmojiEventsIcon sx={{ fontSize: 40, color: '#1E00A1' }} />
          <Typography variant="h6" sx={{ ml: 1 }}>参加データ</Typography>
        </Box>
        <Paper sx={{ p: 2, backgroundColor: '#FFD700', borderRadius: '16px' }}>
          <Box sx={{ display: 'flex', justifyContent: 'space-around', alignItems: 'center' }}>
            <Typography variant="h5" sx={{ color: '#FF007F' }}>5回</Typography>
            <Typography variant="h5" sx={{ color: '#FF007F' }}>20</Typography>
          </Box>
          <Box sx={{ display: 'flex', justifyContent: 'space-around', alignItems: 'center', mt: 1 }}>
            <Typography>回数</Typography>
            <Typography>セッション数</Typography>
          </Box>
        </Paper>
      </Box>

      {/* Trophies List */}
      <Box>
        <Typography variant="h6" sx={{ mb: 1 }}>与えられたトロフィー</Typography>
        <List>
          {trophies.map((trophy, index) => (
            <ListItem key={index}>
              <ListItemIcon>
                <EmojiEventsIcon sx={{ fontSize: 30 }} />
              </ListItemIcon>
              <ListItemText primary={trophy.description} />
            </ListItem>
          ))}
        </List>
      </Box>

      {/* Footer Navigation */}
      <MainBottomNavigation value="stats" />
    </Box>
  );
};

export default Stats;