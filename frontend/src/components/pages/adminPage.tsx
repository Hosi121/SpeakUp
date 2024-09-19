import { useState } from 'react';
import { 
  Box, Button, Dialog, DialogTitle, DialogContent, DialogActions, 
  TextField, Typography, Snackbar, Tab, Tabs, Paper, Alert 
} from '@mui/material';
import api from '../../services/api';

interface EventDetails {
  dateTime: string;
  theme: string;
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

const AdminPage: React.FC = () => {
  const [openDialog, setOpenDialog] = useState<boolean>(false);
  const [dateTime, setDateTime] = useState<string>('');
  const [theme, setTheme] = useState<string>('');
  const [successMessage, setSuccessMessage] = useState<string>('');
  const [eventDetails, setEventDetails] = useState<EventDetails | null>(null);
  const [activeTab, setActiveTab] = useState<number>(0);

  const handleOpenDialog = () => {
    setOpenDialog(true);
    setSuccessMessage('');
    setEventDetails(null);
  };

  const handleCloseDialog = () => {
    setOpenDialog(false);
    setDateTime('');
    setTheme('');
  };

  const handleGenerateTheme = async () => {
    try {
      const prompt = 'イベントのテーマを提案してください。';
      const response = await api.post<ChatResponse>('/chat/theme', { content: prompt });
      const generatedTheme = response.data.choices[0].message.content;
      setTheme(generatedTheme);
    } catch (error) {
      console.error('Failed to generate theme', error);
    }
  };

  const handleSubmit = () => {
    // TODO: Backendへのデータ送信処理をここに記述
    // 送信が成功したと仮定
    setSuccessMessage('成功');
    setEventDetails({ dateTime, theme });
    handleCloseDialog();
  };

  const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
    setActiveTab(newValue);
  };

  const handleSnackbarClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === 'clickaway') {
      return;
    }
    setSuccessMessage('');
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
            イベント作成欄
          </Button>

          {successMessage && (
            <Snackbar
              open={true}
              autoHideDuration={6000}
              onClose={handleSnackbarClose}
            >
              <Alert
                onClose={handleSnackbarClose}
                severity="success"
                sx={{ width: '100%' }}
              >
                {successMessage}
              </Alert>
            </Snackbar>
          )}

          {eventDetails && (
            <Paper sx={{ mt: 4, p: 2 }}>
              <Typography variant="h6">イベント内容</Typography>
              <Typography variant="body1">
                予定日時: {new Date(eventDetails.dateTime).toLocaleString()}
              </Typography>
              <Typography variant="body1">テーマ: {eventDetails.theme}</Typography>
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
                  onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setDateTime(e.target.value)
                  }
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </Box>
              <Box sx={{ mt: 2 }}>
                <TextField
                  fullWidth
                  label="テーマ"
                  value={theme}
                  onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                    setTheme(e.target.value)
                  }
                />
                <Button
                  variant="outlined"
                  sx={{ mt: 1 }}
                  onClick={handleGenerateTheme}
                >
                  AIによる生成
                </Button>
              </Box>
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