import { HalfModal } from "../utils/HalfModal";
import { useState, useEffect } from "react";
import { Typography, Box, List, ListItem, Paper, ListItemText, TextField, Button, Tab } from "@mui/material";
import HomeLogo from "../../assets/homeLogo";
import { Favorite, Person } from "@mui/icons-material";
import { SessionBottomNavigationTemplate } from "../templates/SessionBottomNavigationTemplate";
import SessionContainer from "../utils/SessionContainer";
import api from "../../services/api";
import { fetchMemo } from "../../services/memoService"; // Import the fetchMemo function
import { TabContext, TabList, TabPanel } from "@mui/lab";
import { useNavigate } from "react-router-dom";

const users = [
  { name: "User1", icon: <Person />, description: "英語" },
  { name: "User2", icon: <Favorite />, description: "苗字" },
];

const theme = "好きな言葉";

export const Session = () => {
  const [memoOpen, setMemoOpen] = useState(false);
  const [assistantOpen, setAssistantOpen] = useState(false);
  const [messages, setMessages] = useState<string[]>([]);
  const [inputMessage, setInputMessage] = useState<string>("");
  const [isLoading, setIsLoading] = useState(false); // ローディング状態を管理
  const [memo1, setMemo1] = useState(""); // メモ1を管理
  const [memo2, setMemo2] = useState(""); // メモ2を管理
  const [value, setValue] = useState("1");
  const [countdown, setCountdown] = useState(303);
  const navigate = useNavigate();

  useEffect(() => {
    if (countdown > 0) {
      const timer = setTimeout(() => setCountdown(countdown - 1), 1000);
      return () => clearTimeout(timer);
    } else {
      navigate("/sessioninterval");
    }
  }, [countdown, navigate]);
  useEffect(() => {
    // コンポーネント読み込み時にメモを取得
    const getMemo = async () => {
      try {
        const data = await fetchMemo();
        setMemo1(data.memo1 || ""); // memo1だけ取得
        setMemo2(data.memo2 || "");
      } catch (error) {
        console.error("Failed to fetch memo", error);
      }
    };
    getMemo();
  }, []);

  const handleMemoClose = () => setMemoOpen(false);
  const handleAssistantClose = () => setAssistantOpen(false);

  const handleChange = (event: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
  };
  const handleSendMessage = async () => {
    if (inputMessage.trim() === "") return;

    setIsLoading(true); // ローディング状態を開始
    setMessages([...messages, `You: ${inputMessage}`]); // ユーザーのメッセージを表示
    const userMessage = inputMessage;
    setInputMessage(""); // 送信後に入力フィールドをクリア

    try {
      // サーバーにリクエストを送信
      const response = await api.post("/chat/ask", {
        content: userMessage,
      });

      // 応答メッセージを表示
      if (response.data && response.data.choices && response.data.choices[0].message.content) {
        const assistantMessage = response.data.choices[0].message.content;
        setMessages((prevMessages) => [...prevMessages, `Assistant: ${assistantMessage}`]);
      } else {
        setMessages((prevMessages) => [...prevMessages, "Error: Invalid response format"]);
      }
    } catch (error) {
      setMessages((prevMessages) => [...prevMessages, "Error: Failed to get response"]);
    } finally {
      setIsLoading(false); // ローディング状態を終了
    }
  };

  return (
    <SessionBottomNavigationTemplate value="other" isMute={false} setMemoOpen={setMemoOpen} setAssistantOpen={setAssistantOpen}>
      <SessionContainer theme={theme} users={users} />
      <HalfModal open={memoOpen} handleClose={handleMemoClose} title="">
        <TabContext value={value}>
          <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
            <TabList onChange={handleChange} sx={{ display: "grid", placeContent: "center" }}>
              <Tab label=" 持ち込みメモ" value="1" />
              <Tab label="ワードリスト" value="2" />
            </TabList>
          </Box>
          <TabPanel value="1">
            <Typography variant="body1">{memo1}</Typography>
          </TabPanel>
          <TabPanel value="2">
            <Typography variant="body1">{memo2}</Typography>
          </TabPanel>
        </TabContext>
      </HalfModal>

      <HalfModal open={assistantOpen} handleClose={handleAssistantClose} title="アシスタント">
        <Box sx={{ overflow: "auto", pt: 1, pb: 1, maxHeight: "30vh" }}>
          <List>
            {/* アシスタントからの初期メッセージ */}
            <ListItem sx={{ justifyContent: "flex-start" }}>
              <Box
                sx={{
                  width: "40px",
                  height: "40px",
                  borderRadius: "50%",
                  backgroundColor: "secondary.main",
                  mr: 2,
                  display: "grid",
                  placeContent: "center",
                }}
              >
                <HomeLogo style={{ width: "70%", height: "fit-content" }} />
              </Box>
              <Paper
                sx={{
                  padding: "5px",
                  backgroundColor: "background.default",
                  maxWidth: "60%",
                  wordWrap: "break-word",
                }}
              >
                <ListItemText primary="何かお困りですか？" />
              </Paper>
            </ListItem>

            {/* メッセージのリスト */}
            {messages.map((message, index) => (
              <ListItem key={index} sx={{ justifyContent: message.startsWith("You:") ? "flex-end" : "flex-start" }}>
                <Paper
                  sx={{
                    padding: "5px",
                    backgroundColor: message.startsWith("You:") ? "#f0f0f0" : "background.default",
                    maxWidth: "60%",
                    wordWrap: "break-word",
                  }}
                >
                  <ListItemText primary={message} />
                </Paper>
              </ListItem>
            ))}
          </List>
        </Box>

        {/* メッセージ入力欄と送信ボタン */}
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            pb: 2,
            position: "fixed",
            bottom: 0,
            backgroundColor: "secondary.main",
          }}
        >
          <TextField
            variant="outlined"
            placeholder="メッセージを入力"
            fullWidth
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
            sx={{ mr: 2 }}
            InputProps={{
              style: {
                height: "40px",
              },
            }}
            disabled={isLoading} // ローディング中は入力を無効化
          />
          <Button variant="contained" color="primary" onClick={handleSendMessage} disabled={isLoading}>
            {isLoading ? "送信中..." : "送信"}
          </Button>
        </Box>
      </HalfModal>
    </SessionBottomNavigationTemplate>
  );
};

export default Session;
