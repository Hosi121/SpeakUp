import { useState, useEffect } from "react";
import {
  Box,
  Typography,
  IconButton,
  Avatar,
  TextField,
  Button,
  Paper,
  Grid,
  Divider,
  useTheme,
  styled,
  Container,
} from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import SettingsIcon from "@mui/icons-material/Settings";
import LogoutIcon from "@mui/icons-material/Logout";
import AddIcon from "@mui/icons-material/Add";
import CheckIcon from "@mui/icons-material/Check";
import CloseIcon from "@mui/icons-material/Close";
import { BottomNavigationTemplate } from "../templates/BottomNavigationTemplate";
import api from "../../services/api";
export interface UserData {
  id: number;
  username: string;
  email: string;
  avatar_url: string;
  role: string;
  created_at: string;
  updated_at: string;
}

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

const SettingsContainer = () => {
  const theme = useTheme();
  const [user, setUser] = useState<UserData | null>(null);
  const [editName, setEditName] = useState(false);
  const [editEmail, setEditEmail] = useState(false);
  const [newName, setNewName] = useState("");
  const [newEmail, setNewEmail] = useState("");

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await api.get<UserData>("/user/info");
        console.log("Fetched user data:", response.data);
        setUser(response.data);
        setNewName(response.data.username);
        setNewEmail(response.data.email);
      } catch (error) {
        console.error("Failed to fetch user data:", error);
      }
    };

    fetchUserData();
  }, []);

  const getFullAvatarUrl = (avatarUrl: string) => {
    if (!avatarUrl) return ''; // デフォルトのアバター画像のURLを設定することもできます
    if (avatarUrl.startsWith('http')) return avatarUrl; // すでに完全なURLの場合
    return `http://localhost:8081${avatarUrl}`; // ローカル開発環境の場合
  };

  const handleAvatarUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const formData = new FormData();
      formData.append("avatar", event.target.files[0]);

      try {
        const response = await api.put("/user/avatar", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        });
        setUser((prevUser) => prevUser ? { ...prevUser, avatar_url: response.data.avatar_url } : null);
      } catch (error) {
        console.error("Failed to upload avatar:", error);
        // Handle error (e.g., show notification)
      }
    }
  };

  const handleSaveName = () => {
    if (user) {
      // Update the user's name via API
      const updateUser = async () => {
        try {
          await api.put(`/user/update`, { username: newName });
          setUser({ ...user, username: newName });
          setEditName(false);
        } catch (error) {
          console.error("Failed to update username:", error);
          // Handle error
        }
      };
      updateUser();
    }
  };

  const handleSaveEmail = () => {
    if (user) {
      // Update the user's email via API
      const updateUser = async () => {
        try {
          await api.put(`/user/update`, { email: newEmail });
          setUser({ ...user, email: newEmail });
          setEditEmail(false);
        } catch (error) {
          console.error("Failed to update email:", error);
          // Handle error
        }
      };
      updateUser();
    }
  };


  const handleLogout = () => {
    // Implement logout logic
    localStorage.removeItem("token"); // Remove the token
    // Redirect to login page or update state accordingly
  };

  if (!user) {
    // Show a loading state or skeleton until user data is fetched
    return <Typography>Loading...</Typography>;
  }

  return (
    <Container
      sx={{ display: "flex", flexDirection: "column", justifyContent: "space-between", height: "100vh" }}
    >
      <Container sx={{ pt: 3 }}>
        <Typography
          variant="h4"
          sx={{ mb: 4, fontWeight: "bold", textAlign: "left" }}
        >
          <SettingsIcon sx={{ fontSize: 40, mr: 2, verticalAlign: "bottom" }} />
          設定
        </Typography>

        <Grid container spacing={4}>
          <Grid item xs={12} md={6}>
            <StyledPaper elevation={0}>
              <Box sx={{ display: "flex", alignItems: "center", mb: 3 }}>
                <AvatarUpload>
                  <Avatar src={getFullAvatarUrl(user.avatar_url)} sx={{ width: 100, height: 100 }} />
                  <UploadButton component="label" size="small">
                    <AddIcon />
                    <input type="file" hidden accept="image/*" onChange={handleAvatarUpload} />
                  </UploadButton>
                </AvatarUpload>
                <Box>
                  {editName ? (
                    <Box sx={{ display: "flex", alignItems: "center", mb: 1 }}>
                      <StyledTextField
                        value={newName}
                        onChange={(e) => setNewName(e.target.value)}
                        variant="outlined"
                        size="small"
                      />
                      <IconButton onClick={handleSaveName} size="small" sx={{ ml: 1 }}>
                        <CheckIcon />
                      </IconButton>
                      <IconButton onClick={() => setEditName(false)} size="small">
                        <CloseIcon />
                      </IconButton>
                    </Box>
                  ) : (
                    <Box sx={{ display: "flex", alignItems: "center", mb: 1 }}>
                      <Typography variant="h6">{user.username}</Typography>
                      <IconButton onClick={() => setEditName(true)} size="small" sx={{ ml: 1 }}>
                        <EditIcon fontSize="small" />
                      </IconButton>
                    </Box>
                  )}
                  {editEmail ? (
                    <Box sx={{ display: "flex", alignItems: "center" }}>
                      <StyledTextField
                        value={newEmail}
                        onChange={(e) => setNewEmail(e.target.value)}
                        variant="outlined"
                        size="small"
                      />
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
              {/* Additional settings can go here */}
              <Divider sx={{ my: 2 }} />
              <Typography variant="h6" sx={{ mb: 2 }}>
                クレジット情報
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {/* Display credit information if applicable */}
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
            onClick={handleLogout}
          >
            ログアウト
          </Button>
        </Box>
      </Container>
    </Container>
  );
};

export const Settings = () => {
  return (
    <BottomNavigationTemplate value="other">
      <SettingsContainer />
    </BottomNavigationTemplate>
  );
};
