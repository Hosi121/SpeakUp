import React, { useState, useEffect } from "react";
import { Box, Typography, IconButton, Avatar, TextField, Button, Switch, Paper, Grid, Divider, useTheme, styled } from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import SettingsIcon from "@mui/icons-material/Settings";
import LogoutIcon from "@mui/icons-material/Logout";
import AddIcon from "@mui/icons-material/Add";
import CheckIcon from "@mui/icons-material/Check";
import CloseIcon from "@mui/icons-material/Close";

// Mock data files
import userData from "../../mock/user.json";
import creditData from "../../assets/credit.json";
import { MainBottomNavigation } from "../utils/MainBottomNavigation";

const StyledPaper = styled(Paper)(({ theme }) => ({
  padding: theme.spacing(3),
  borderRadius: theme.shape.borderRadius * 2,
  boxShadow: "0 4px 20px rgba(0,0,0,0.1)",
}));

const StyledTextField = styled(TextField)(({ theme }) => ({
  "& .MuiOutlinedInput-root": {
    borderRadius: theme.shape.borderRadius * 2,
  },
}));

const AvatarUpload = styled(Box)(({ theme }) => ({
  position: "relative",
  width: 100,
  height: 100,
  marginRight: theme.spacing(3),
}));

const UploadButton = styled(IconButton)(({ theme }) => ({
  position: "absolute",
  right: -8,
  bottom: -8,
  backgroundColor: theme.palette.primary.main,
  color: theme.palette.common.white,
  "&:hover": {
    backgroundColor: theme.palette.primary.dark,
  },
}));

const Settings = () => {
  const theme = useTheme();
  const [user, setUser] = useState({ name: "", email: "", avatar: "" });
  const [creditInfo, setCreditInfo] = useState("");
  const [editName, setEditName] = useState(false);
  const [editEmail, setEditEmail] = useState(false);
  const [newName, setNewName] = useState("");
  const [newEmail, setNewEmail] = useState("");

  useEffect(() => {
    setUser(userData);
    setCreditInfo(creditData.creditInfo);
    setNewName(userData.name);
    setNewEmail(userData.email);
  }, []);

  const handleAvatarUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const reader = new FileReader();
      reader.onload = (e) => {
        if (e.target && e.target.result) {
          setUser((prevUser) => ({ ...prevUser, avatar: e.target.result as string }));
        }
      };
      reader.readAsDataURL(event.target.files[0]);
    }
  };

  const handleSaveName = () => {
    setUser({ ...user, name: newName });
    setEditName(false);
  };

  const handleSaveEmail = () => {
    setUser({ ...user, email: newEmail });
    setEditEmail(false);
  };

  return (
    <Box sx={{ display: "flex", flexDirection: "column", justifyContent: "space-between", height: "100vh" }}>
      <Box sx={{ padding: theme.spacing(3), boxSizing: "border-box" }}>
        <Typography variant="h4" sx={{ mb: 4, fontWeight: "bold", textAlign: "left" }}>
          <SettingsIcon sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }} />
          設定
        </Typography>

        <Grid container spacing={4}>
          <Grid item xs={12} md={6}>
            <StyledPaper elevation={0}>
              <Box sx={{ display: "flex", alignItems: "center", mb: 3 }}>
                <AvatarUpload>
                  <Avatar src={user.avatar} sx={{ width: 100, height: 100 }} />
                  <UploadButton component="label" size="small">
                    <AddIcon />
                    <input type="file" hidden accept="image/*" onChange={handleAvatarUpload} />
                  </UploadButton>
                </AvatarUpload>
                <Box>
                  {editName ? (
                    <Box sx={{ display: "flex", alignItems: "center", mb: 1 }}>
                      <StyledTextField value={newName} onChange={(e) => setNewName(e.target.value)} variant="outlined" size="small" />
                      <IconButton onClick={handleSaveName} size="small" sx={{ ml: 1 }}>
                        <CheckIcon />
                      </IconButton>
                      <IconButton onClick={() => setEditName(false)} size="small">
                        <CloseIcon />
                      </IconButton>
                    </Box>
                  ) : (
                    <Box sx={{ display: "flex", alignItems: "center", mb: 1 }}>
                      <Typography variant="h6">{user.name}</Typography>
                      <IconButton onClick={() => setEditName(true)} size="small" sx={{ ml: 1 }}>
                        <EditIcon fontSize="small" />
                      </IconButton>
                    </Box>
                  )}
                  {editEmail ? (
                    <Box sx={{ display: "flex", alignItems: "center" }}>
                      <StyledTextField value={newEmail} onChange={(e) => setNewEmail(e.target.value)} variant="outlined" size="small" />
                      <IconButton onClick={handleSaveEmail} size="small" sx={{ ml: 1 }}>
                        <CheckIcon />
                      </IconButton>
                      <IconButton onClick={() => setEditEmail(false)} size="small">
                        <CloseIcon />
                      </IconButton>
                    </Box>
                  ) : (
                    <Box sx={{ display: "flex", alignItems: "center" }}>
                      <Typography variant="body1" color="text.secondary">
                        {user.email}
                      </Typography>
                      <IconButton onClick={() => setEditEmail(true)} size="small" sx={{ ml: 1 }}>
                        <EditIcon fontSize="small" />
                      </IconButton>
                    </Box>
                  )}
                </Box>
              </Box>
            </StyledPaper>
          </Grid>

          <Grid item xs={12} md={6}>
            <StyledPaper elevation={0}>
              <Typography variant="h6" sx={{ mb: 2 }}>
                アカウント設定
              </Typography>
              <Box sx={{ display: "flex", justifyContent: "space-between", alignItems: "center", mb: 2 }}>
                <Typography variant="body1">通知</Typography>
                <Switch />
              </Box>
              <Box sx={{ display: "flex", justifyContent: "space-between", alignItems: "center", mb: 2 }}>
                <Typography variant="body1">ダークモード</Typography>
                <Switch />
              </Box>
              <Divider sx={{ my: 2 }} />
              <Typography variant="h6" sx={{ mb: 2 }}>
                クレジット情報
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {creditInfo}
              </Typography>
            </StyledPaper>
          </Grid>
        </Grid>

        <Box sx={{ mt: 4, textAlign: "center" }}>
          <Button
            variant="contained"
            startIcon={<LogoutIcon />}
            sx={{
              borderRadius: theme.shape.borderRadius * 2,
              textTransform: "none",
              px: 4,
              py: 1,
            }}
          >
            ログアウト
          </Button>
        </Box>
      </Box>
      <MainBottomNavigation value="home" />
    </Box>
  );
};

export default Settings;
