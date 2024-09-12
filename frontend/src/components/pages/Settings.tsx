import React, { useState, useEffect } from 'react';
import { Box, Typography, IconButton, Avatar, TextField, Button, Switch, Dialog, DialogActions, DialogContent, DialogTitle } from '@mui/material';
import EditIcon from '@mui/icons-material/Edit';
import SettingsIcon from '@mui/icons-material/Settings';
import LogoutIcon from '@mui/icons-material/Logout';
import UploadFileIcon from '@mui/icons-material/UploadFile';

// Mock data files
import userData from '../../mock/user.json';
import creditData from '../../assets/credit.json';

const Settings = () => {
  const [user, setUser] = useState({ name: '', email: '', avatar: '' });
  const [creditInfo, setCreditInfo] = useState('');
  const [editMode, setEditMode] = useState(false);
  const [newName, setNewName] = useState('');
  const [newEmail, setNewEmail] = useState('');

  // Load user and credit data
  useEffect(() => {
    setUser(userData);
    setCreditInfo(creditData.creditInfo);
    setNewName(userData.name);
    setNewEmail(userData.email);
  }, []);

  // Handle avatar file upload
  const handleAvatarUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const reader = new FileReader();
      reader.onload = (e) => {
        if (e.target && e.target.result) {  // Null check for e.target
          setUser((prevUser) => ({ ...prevUser, avatar: e.target.result as string }));
        }
      };
      reader.readAsDataURL(event.target.files[0]);
    }
  };

  // Handle save for name and email changes
  const handleSaveChanges = () => {
    setUser({ ...user, name: newName, email: newEmail });
    setEditMode(false);
  };

  return (
    <Box sx={{ padding: 2 }}>
      {/* Header with Settings Icon */}
      <Box sx={{ display: 'flex', alignItems: 'center', mb: 2 }}>
        <SettingsIcon sx={{ fontSize: 40, mr: 2 }} />
        <Typography variant="h5">設定</Typography>
      </Box>

      {/* User Information Box */}
      <Box
        sx={{
            backgroundColor: '#FFDD66',
            padding: 2,
            borderRadius: 2,
            mb: 4,
            position: 'relative',
            width: '100%', // 横幅を全体に広げる
            minHeight: '50vh', // 高さを画面の50%に設定（必要に応じて調整）
        }}
        >
        <Avatar src={user.avatar} sx={{ width: 80, height: 80, mb: 2 }} />
        <Typography variant="h6">{user.name}</Typography>
        <Typography variant="body1">{user.email}</Typography>

        {/* Edit Button */}
        <IconButton
          sx={{ position: 'absolute', top: 16, right: 16 }}
          onClick={() => setEditMode(true)}
        >
          <EditIcon />
        </IconButton>

        {/* Upload avatar */}
        <Box mt={2}>
          <Button
            variant="contained"
            component="label"
            startIcon={<UploadFileIcon />}
            sx={{ mt: 2, backgroundColor: '#FF007F', color: 'white' }}
          >
            アバターをアップロード
            <input type="file" hidden accept="image/*" onChange={handleAvatarUpload} />
          </Button>
        </Box>

      {/* Notifications and Dark Mode */}
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
        <Typography variant="body1">通知</Typography>
        <Switch />
      </Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 2 }}>
        <Typography variant="body1">ダークモード</Typography>
        <Switch />
      </Box>

      {/* Credit Information */}
      <Typography variant="body2" sx={{ mb: 2 }}>
        クレジット情報
      </Typography>
      <Typography variant="body2" sx={{ mb: 4 }}>
        {creditInfo}
      </Typography>

      {/* Logout */}
      <Button sx={{ color: 'red', textTransform: 'none' }} startIcon={<LogoutIcon />}>
        ログアウト
      </Button>

      {/* Edit Modal for Name and Email */}
      <Dialog open={editMode} onClose={() => setEditMode(false)}>
        <DialogTitle>登録情報を編集</DialogTitle>
        <DialogContent>
          <TextField
            label="名前"
            fullWidth
            variant="outlined"
            margin="normal"
            value={newName}
            onChange={(e) => setNewName(e.target.value)}
          />
          <TextField
            label="Email"
            fullWidth
            variant="outlined"
            margin="normal"
            value={newEmail}
            onChange={(e) => setNewEmail(e.target.value)}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setEditMode(false)} color="primary">
            キャンセル
          </Button>
          <Button onClick={handleSaveChanges} color="primary">
            保存
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
    </Box>
  );
};

export default Settings;