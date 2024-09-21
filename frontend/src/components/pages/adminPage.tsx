import { useState, useEffect } from "react";
import {
  Box,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Typography,
  Snackbar,
  Tab,
  Tabs,
  Paper,
  Alert,
  List,
  ListItem,
  ListItemText,
  ListItemAvatar,
  Avatar,
  CircularProgress,
} from "@mui/material";
import api from "../../services/api";
import { Event, EventDetails, User } from "../../types/types";
import * as eventService from "../../services/eventService";

const AdminPage: React.FC = () => {
  const [openDialog, setOpenDialog] = useState<boolean>(false);
  const [dateTime, setDateTime] = useState<string>("");
  const [theme, setTheme] = useState<string>("");
  const [topics, setTopics] = useState<string[]>(["", "", ""]);
  const [successMessage, setSuccessMessage] = useState<string>("");
  const [errorMessage, setErrorMessage] = useState<string>("");
  const [eventDetails, setEventDetails] = useState<EventDetails | null>(null);
  const [activeTab, setActiveTab] = useState<number>(0);
  const [, setEvents] = useState<Event[]>([]);
  const [searchQuery, setSearchQuery] = useState<string>("");
  const [users, setUsers] = useState<User[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchInitialEvents = async () => {
      try {
        const fetchedEvents = await eventService.fetchEvents();
        setEvents(fetchedEvents);
      } catch (error) {
        console.error("Failed to fetch initial events", error);
        setErrorMessage("初期イベントの取得に失敗しました");
      }
    };
    fetchInitialEvents();
  }, []);

  const handleGenerateTheme = async () => {
    try {
      const generatedTheme = await eventService.generateTheme();
      setTheme(generatedTheme);
    } catch (error) {
      setErrorMessage("テーマの生成に失敗しました");
      console.log(error);
    }
  };

  const handleSubmit = async () => {
    try {
      const isoDateTime = new Date(dateTime).toISOString();
      const eventDetails: EventDetails = {
        dateTime: isoDateTime,
        theme,
        topics,
      };

      await eventService.createEvent(eventDetails);
      setSuccessMessage("イベントが正常に作成されました");
      setEventDetails(eventDetails);
      handleCloseDialog();
      const fetchedEvents = await eventService.fetchEvents();
      setEvents(fetchedEvents);
    } catch (error) {
      setErrorMessage("イベントの作成に失敗しました");
      console.log(error);
    }
  };

  const handleOpenDialog = () => {
    setOpenDialog(true);
    setSuccessMessage("");
    setErrorMessage("");
    setEventDetails(null);
  };

  const handleCloseDialog = () => {
    setOpenDialog(false);
    setDateTime("");
    setTheme("");
    setTopics(["", "", ""]);
  };

  const handleTopicChange = (index: number, value: string) => {
    const newTopics = [...topics];
    newTopics[index] = value;
    setTopics(newTopics);
  };

  const handleTabChange = (_: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue);
  };

  const handleSnackbarClose = () => {
    setSuccessMessage("");
    setErrorMessage("");
  };

  const handleSearchUsers = async () => {
    setIsLoading(true);
    try {
      const response = await api.get(`/users/search?q=${searchQuery}`);
      const usersWithAvatars = await Promise.all(
        response.data.map(async (user: User) => {
          try {
            const avatarResponse = await api.get(`/users/${user.id}/avatar`);
            return { ...user, avatarURL: avatarResponse.data.avatarURL };
          } catch (error) {
            console.error(`Failed to fetch avatar for user ${user.id}`, error);
            return { ...user, avatarURL: "" };
          }
        })
      );
      setUsers(usersWithAvatars);
    } catch (error) {
      console.error("Failed to search users", error);
      setErrorMessage("ユーザーの検索に失敗しました");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Box sx={{ p: 2 }}>
      <Tabs value={activeTab} onChange={handleTabChange} centered>
        <Tab label="イベント管理" />
        <Tab label="ユーザー情報" />
      </Tabs>

      {activeTab === 0 && (
        <Box sx={{ mt: 4 }}>
          <Button variant="contained" onClick={handleOpenDialog}>
            イベント作成
          </Button>

          {eventDetails && (
            <Paper sx={{ mt: 4, p: 5 }}>
              <Typography
                variant="h6"
                fontWeight="bolder"
                color="primary.main"
                sx={{ mb: 1 }}
              >
                最後に作成したイベント
              </Typography>
              <Typography variant="body1">
                予定日時: {new Date(eventDetails.dateTime).toLocaleString()}
              </Typography>
              <Typography
                variant="h6"
                fontWeight="bolder"
                fontSize="1.1rem"
                sx={{ mt: 3, mb: 1 }}
              >
                テーマ
              </Typography>
              <Typography variant="body1">{eventDetails.theme}</Typography>
              <Typography
                variant="h6"
                fontWeight="bolder"
                fontSize="1.1rem"
                sx={{ mt: 3, mb: 1 }}
              >
                トピック
              </Typography>
              <List sx={{ width: "100%", p: 0 }}>
                {eventDetails.topics.map((topic, index) => (
                  <ListItem key={index} sx={{ p: 0 }}>
                    <ListItemText sx={{ textAlign: "center" }}>
                      {topic || "(未入力)"}
                    </ListItemText>
                  </ListItem>
                ))}
              </List>
            </Paper>
          )}
          <Dialog open={openDialog} onClose={handleCloseDialog}>
            <DialogTitle>イベント作成</DialogTitle>
            <DialogContent>
              <Box sx={{ mt: 2 }}>
                <TextField
                  fullWidth
                  label="予定日時"
                  type="datetime-local"
                  value={dateTime}
                  onChange={(e) => setDateTime(e.target.value)}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  inputProps={{
                    step: 1,
                  }}
                />
              </Box>
              <Box sx={{ mt: 2 }}>
                <TextField
                  fullWidth
                  label="テーマ"
                  value={theme}
                  onChange={(e) => setTheme(e.target.value)}
                />
                <Button
                  variant="outlined"
                  sx={{ mt: 1 }}
                  onClick={handleGenerateTheme}
                >
                  AIによる生成
                </Button>
              </Box>
              {[0, 1, 2].map((index) => (
                <Box sx={{ mt: 2 }} key={index}>
                  <TextField
                    fullWidth
                    label={`トピック ${index + 1}`}
                    value={topics[index]}
                    onChange={(e) => handleTopicChange(index, e.target.value)}
                    placeholder="ここにトピックを入力してください"
                  />
                </Box>
              ))}
            </DialogContent>
            <DialogActions>
              <Button onClick={handleCloseDialog}>キャンセル</Button>
              <Button onClick={handleSubmit} variant="contained">
                作成
              </Button>
            </DialogActions>
          </Dialog>
        </Box>
      )}

      {activeTab === 1 && (
        <Box sx={{ mt: 4 }}>
          <Typography variant="h6">ユーザー情報検索</Typography>
          <Box sx={{ display: "flex", alignItems: "center", mt: 2 }}>
            <TextField
              fullWidth
              label="ユーザー名で検索"
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
            />
            <Button
              variant="contained"
              sx={{ ml: 2 }}
              onClick={handleSearchUsers}
              disabled={isLoading}
            >
              {isLoading ? <CircularProgress size={24} /> : "検索"}
            </Button>
          </Box>

          <List sx={{ mt: 2 }}>
            {isLoading ? (
              <ListItem>
                <CircularProgress />
              </ListItem>
            ) : users && users.length > 0 ? (
              users.map((user) => (
                <ListItem key={user.id}>
                  <ListItemAvatar>
                    <Avatar
                      src={user.avatarURL || "/default-avatar.png"}
                      alt={user.username}
                    />
                  </ListItemAvatar>
                  <ListItemText
                    primary={user.username}
                    secondary={
                      <>
                        <Typography
                          component="span"
                          variant="body2"
                          color="text.primary"
                        >
                          メール: {user.email}
                        </Typography>
                        <br />
                        登録日: {new Date(user.createdAt).toLocaleString()}
                      </>
                    }
                  />
                </ListItem>
              ))
            ) : (
              <ListItem>
                <ListItemText primary="ユーザーが見つかりません" />
              </ListItem>
            )}
          </List>
        </Box>
      )}

      <Snackbar
        open={!!successMessage || !!errorMessage}
        autoHideDuration={6000}
        onClose={handleSnackbarClose}
      >
        <Alert
          onClose={handleSnackbarClose}
          severity={successMessage ? "success" : "error"}
          sx={{ width: "100%" }}
        >
          {successMessage || errorMessage}
        </Alert>
      </Snackbar>
    </Box>
  );
};

export default AdminPage;
