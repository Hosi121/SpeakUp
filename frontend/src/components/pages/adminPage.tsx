import { useState, useEffect } from 'react';
import { 
  Box, Button, Dialog, DialogTitle, DialogContent, DialogActions, 
  TextField, Typography, Snackbar, Tab, Tabs, Paper, Alert,
  List, ListItem, ListItemText, ListItemAvatar, Avatar, CircularProgress
} from '@mui/material';
import api from '../../services/api';
import { ChatResponse, Event, EventDetails, User } from '../../types';

const AdminPage: React.FC = () => {
  const [openDialog, setOpenDialog] = useState<boolean>(false);
  const [dateTime, setDateTime] = useState<string>('');
  const [theme, setTheme] = useState<string>('');
  const [topics, setTopics] = useState<string[]>(['', '', '']);
  const [successMessage, setSuccessMessage] = useState<string>('');
  const [errorMessage, setErrorMessage] = useState<string>('');
  const [eventDetails, setEventDetails] = useState<EventDetails | null>(null);
  const [activeTab, setActiveTab] = useState<number>(0);
  const [events, setEvents] = useState<Event[]>([]);
  const [searchQuery, setSearchQuery] = useState<string>('');
  const [users, setUsers] = useState<User[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  useEffect(() => {
    const fetchInitialEvents = async () => {
      try {
        await fetchEvents();
      } catch (error) {
        console.error('Failed to fetch initial events', error);
        setErrorMessage('初期イベントの取得に失敗しました');
      }
    };
    fetchInitialEvents();
  }, []);

  const fetchEvents = async () => {
    try {
      const response = await api.get('/events');
      setEvents(response.data);
    } catch (error) {
      console.error('Failed to fetch events', error);
      setErrorMessage('イベントの取得に失敗しました');
    }
  };

  const handleOpenDialog = () => {
    setOpenDialog(true);
    setSuccessMessage('');
    setErrorMessage('');
    setEventDetails(null);
  };

  const handleCloseDialog = () => {
    setOpenDialog(false);
    setDateTime('');
    setTheme('');
    setTopics(['', '', '']);
  };

  const handleGenerateTheme = async () => {
    try {
      const prompt = 'イベントのテーマを提案してください。';
      const response = await api.post<ChatResponse>('/chat/theme', { content: prompt });
      const generatedTheme = response.data.choices[0].message.content;
      setTheme(generatedTheme);
    } catch (error) {
      console.error('Failed to generate theme', error);
      setErrorMessage('テーマの生成に失敗しました');
    }
  };

  const handleTopicChange = (index: number, value: string) => {
    const newTopics = [...topics];
    newTopics[index] = value;
    setTopics(newTopics);
  };

  const handleSubmit = async () => {
    try {
      const isoDateTime = new Date(dateTime).toISOString();
      
      const response = await api.post('/events', {
        dateTime: isoDateTime,
        theme,
        topics
      });
      setSuccessMessage('イベントが正常に作成されました');
      setEventDetails({ dateTime: isoDateTime, theme, topics });
      handleCloseDialog();
      fetchEvents();
    } catch (error) {
      console.error('Failed to create event', error);
      setErrorMessage('イベントの作成に失敗しました');
    }
  };

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue);
  };

  const handleSnackbarClose = () => {
    setSuccessMessage('');
    setErrorMessage('');
  };

  const handleSearchUsers = async () => {
    setIsLoading(true);
    try {
      const response = await api.get(`/users/search?q=${searchQuery}`);
      const usersWithAvatars = await Promise.all(response.data.map(async (user: User) => {
        try {
          const avatarResponse = await api.get(`/users/${user.id}/avatar`);
          return { ...user, avatarURL: avatarResponse.data.avatarURL };
        } catch (error) {
          console.error(`Failed to fetch avatar for user ${user.id}`, error);
          return { ...user, avatarURL: '' };
        }
      }));
      setUsers(usersWithAvatars);
    } catch (error) {
      console.error('Failed to search users', error);
      setErrorMessage('ユーザーの検索に失敗しました');
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
            <Paper sx={{ mt: 4, p: 2 }}>
              <Typography variant="h6">最後に作成したイベント</Typography>
              <Typography variant="body1">
                予定日時: {new Date(eventDetails.dateTime).toLocaleString()}
              </Typography>
              <Typography variant="body1">テーマ: {eventDetails.theme}</Typography>
              <Typography variant="body1">トピック:</Typography>
              <ul>
                {eventDetails.topics.map((topic, index) => (
                  <li key={index}>{topic || '(未入力)'}</li>
                ))}
              </ul>
            </Paper>
          )}

          <Paper sx={{ mt: 4, p: 2 }}>
            <Typography variant="h6">イベントリスト</Typography>
            <List>
              {events && events.length > 0 ? (
                events.map((event) => (
                  <ListItem key={event.id}>
                    <ListItemText
                      primary={`${new Date(event.event_start).toLocaleString()} - ${new Date(event.event_end).toLocaleString()}`}
                      secondary={
                        <>
                          <Typography component="span" variant="body2" color="text.primary">
                            テーマ: {event.theme.theme_text}
                          </Typography>
                          <br />
                          トピック1: {event.theme.topic1}<br />
                          トピック2: {event.theme.topic2}<br />
                          トピック3: {event.theme.topic3}
                        </>
                      }
                    />
                  </ListItem>
                ))
              ) : (
                <ListItem>
                  <ListItemText primary="イベントがありません" />
                </ListItem>
              )}
            </List>
          </Paper>

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
          <Box sx={{ display: 'flex', alignItems: 'center', mt: 2 }}>
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
              {isLoading ? <CircularProgress size={24} /> : '検索'}
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
                    <Avatar src={user.avatarURL || '/default-avatar.png'} alt={user.username} />
                  </ListItemAvatar>
                  <ListItemText
                    primary={user.username}
                    secondary={
                      <>
                        <Typography component="span" variant="body2" color="text.primary">
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
          sx={{ width: '100%' }}
        >
          {successMessage || errorMessage}
        </Alert>
      </Snackbar>
    </Box>
  );
};

export default AdminPage;