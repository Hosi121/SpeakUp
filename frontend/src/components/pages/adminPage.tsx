import { useState } from 'react';
import { 
  Box, Button, Dialog, DialogTitle, DialogContent, DialogActions, 
  TextField, Typography, Snackbar, Tab, Tabs, Paper 
} from '@mui/material';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';

function AdminPage() {
  const [openDialog, setOpenDialog] = useState(false);
  const [dateTime, setDateTime] = useState(null);
  const [theme, setTheme] = useState('');
  const [successMessage, setSuccessMessage] = useState('');
  const [eventDetails, setEventDetails] = useState(null);
  const [activeTab, setActiveTab] = useState(0);

  const handleOpenDialog = () => {
    setOpenDialog(true);
    setSuccessMessage('');
    setEventDetails(null);
  };

  const handleCloseDialog = () => {
    setOpenDialog(false);
    setDateTime(null);
    setTheme('');
  };

  const handleGenerateTheme = () => {
    // AIによるテーマ生成をシミュレート
    setTheme('自動生成されたテーマ');
  };

  const handleSubmit = () => {
    // TODO: Backendへのデータ送信処理をここに記述
    // 送信が成功したと仮定
    setSuccessMessage('成功');
    setEventDetails({ dateTime, theme });
    handleCloseDialog();
  };

  const handleTabChange = (event, newValue) => {
    setActiveTab(newValue);
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
              message={successMessage}
              autoHideDuration={6000}
            />
          )}

          {eventDetails && (
            <Paper sx={{ mt: 4, p: 2 }}>
              <Typography variant="h6">イベント内容</Typography>
              <Typography variant="body1">
                予定日時: {eventDetails.dateTime?.format('YYYY-MM-DD HH:mm')}
              </Typography>
              <Typography variant="body1">テーマ: {eventDetails.theme}</Typography>
            </Paper>
          )}

          <Dialog open={openDialog} onClose={handleCloseDialog}>
            <DialogTitle>イベント作成</DialogTitle>
            <DialogContent>
              <Box sx={{ mt: 2 }}>
                <DateTimePicker
                  label="予定日時"
                  value={dateTime}
                  onChange={(newValue) => setDateTime(newValue)}
                  renderInput={(params) => <TextField fullWidth {...params} />}
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
}

export default AdminPage;