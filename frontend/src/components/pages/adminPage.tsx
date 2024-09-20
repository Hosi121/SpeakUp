import React, { useState, useEffect } from "react";
import { Box, Button, Dialog, DialogTitle, DialogContent, DialogActions, TextField, Typography, Snackbar, Tab, Tabs, Paper, Alert, List, ListItem, ListItemText } from "@mui/material";
import api from "../../services/api";

interface EventDetails {
  dateTime: string;
  theme: string;
  topics: string[];
}

interface ChatResponse {
  choices: Array<{
    message: {
      role: string;
      content: string;
    };
    index: number;
    finish_reason: string;
  }>;
}

interface Event {
  id: number;
  event_start: string;
  event_end: string;
  theme_id: number;
  theme: {
    theme_text: string;
    topic1: string;
    topic2: string;
    topic3: string;
  };
}

const AdminPage: React.FC = () => {
  const [openDialog, setOpenDialog] = useState<boolean>(false);
  const [dateTime, setDateTime] = useState<string>("");
  const [theme, setTheme] = useState<string>("");
  const [topics, setTopics] = useState<string[]>(["", "", ""]);
  const [successMessage, setSuccessMessage] = useState<string>("");
  const [errorMessage, setErrorMessage] = useState<string>("");
  const [eventDetails, setEventDetails] = useState<EventDetails | null>(null);
  const [activeTab, setActiveTab] = useState<number>(0);
  const [events, setEvents] = useState<Event[]>([]);

  useEffect(() => {
    fetchEvents();
  }, []);

  const fetchEvents = async () => {
    try {
      const response = await api.get("/events");
      setEvents(response.data);
    } catch (error) {
      console.error("Failed to fetch events", error);
      setErrorMessage("イベントの取得に失敗しました");
      setEvents([]);
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

  const handleGenerateTheme = async () => {
    try {
      const prompt = "イベントのテーマを提案してください。";
      const response = await api.post<ChatResponse>("/chat/theme", { content: prompt });
      const generatedTheme = response.data.choices[0].message.content;
      setTheme(generatedTheme);
    } catch (error) {
      console.error("Failed to generate theme", error);
      setErrorMessage("テーマの生成に失敗しました");
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

      const response = await api.post("/events", {
        dateTime: isoDateTime,
        theme,
        topics,
      });
      setSuccessMessage("イベントが正常に作成されました");
      setEventDetails({ dateTime: isoDateTime, theme, topics });
      handleCloseDialog();
      fetchEvents(); // イベントリストを更新
    } catch (error) {
      console.error("Failed to create event", error);
      setErrorMessage("イベントの作成に失敗しました");
    }
  };

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue);
  };

  const handleSnackbarClose = () => {
    setSuccessMessage("");
    setErrorMessage("");
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

          <Snackbar open={!!successMessage || !!errorMessage} autoHideDuration={6000} onClose={handleSnackbarClose}>
            <Alert onClose={handleSnackbarClose} severity={successMessage ? "success" : "error"} sx={{ width: "100%" }}>
              {successMessage || errorMessage}
            </Alert>
          </Snackbar>

          {eventDetails && (
            <Paper sx={{ mt: 4, p: 2 }}>
              <Typography variant="h6">最後に作成したイベント</Typography>
              <Typography variant="body1">予定日時: {new Date(eventDetails.dateTime).toLocaleString()}</Typography>
              <Typography variant="body1">テーマ: {eventDetails.theme}</Typography>
              <Typography variant="body1">トピック:</Typography>
              <ul>
                {eventDetails.topics.map((topic, index) => (
                  <li key={index}>{topic || "(未入力)"}</li>
                ))}
              </ul>
            </Paper>
          )}

          <Paper sx={{ mt: 4, p: 2 }}>
            <Typography variant="h6">イベントリスト</Typography>
            {events ? (
              <List>
                {events.map((event) => (
                  <ListItem key={event.id}>
                    <ListItemText
                      primary={`${new Date(event.event_start).toLocaleString()} - ${new Date(event.event_end).toLocaleString()}`}
                      secondary={
                        <>
                          <Typography component="span" variant="body2" color="text.primary">
                            テーマ: {event.theme.theme_text}
                          </Typography>
                          <br />
                          トピック1: {event.theme.topic1}
                          <br />
                          トピック2: {event.theme.topic2}
                          <br />
                          トピック3: {event.theme.topic3}
                        </>
                      }
                    />
                  </ListItem>
                ))}
              </List>
            ) : (
              <Typography>イベントがありません</Typography>
            )}
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
                  onChange={(e: React.ChangeEvent<HTMLInputElement>) => setDateTime(e.target.value)}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  inputProps={{
                    step: 1, // 1秒単位
                  }}
                />
              </Box>
              <Box sx={{ mt: 2 }}>
                <TextField fullWidth label="テーマ" value={theme} onChange={(e: React.ChangeEvent<HTMLInputElement>) => setTheme(e.target.value)} />
                <Button variant="outlined" sx={{ mt: 1 }} onClick={handleGenerateTheme}>
                  AIによる生成
                </Button>
              </Box>
              {[0, 1, 2].map((index) => (
                <Box sx={{ mt: 2 }} key={index}>
                  <TextField fullWidth label={`トピック ${index + 1}`} value={topics[index]} onChange={(e: React.ChangeEvent<HTMLInputElement>) => handleTopicChange(index, e.target.value)} placeholder="ここにトピックを入力してください" />
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
          <TextField
            fullWidth
            label="ユーザー名で検索"
            sx={{ mt: 2 }}
            // 検索機能は未実装のため、入力欄のみ配置
          />
        </Box>
      )}
    </Box>
  );
};

export default AdminPage;
