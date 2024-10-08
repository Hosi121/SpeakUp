import { HalfModal } from "../utils/HalfModal";
import { useCallback, useContext, useEffect, useRef, useState } from "react";
import {
  Typography,
  Box,
  List,
  ListItem,
  Paper,
  ListItemText,
  TextField,
  Button,
  Tab,
  Avatar,
} from "@mui/material";
import HomeLogo from "../../assets/homeLogo";
import { Person } from "@mui/icons-material";
import { SessionBottomNavigationTemplate } from "../templates/SessionBottomNavigationTemplate";
import SessionContainer from "../utils/SessionContainer";
import api from "../../services/api";
import { fetchMemo } from "../../services/memoService"; // Import the fetchMemo function
import { TabContext, TabList, TabPanel } from "@mui/lab";
import { useNavigate } from "react-router-dom";
import { TopicPopup } from "../utils/TopicPopup";
import { AudioVolumeAnalyzer } from "../utils/AudioVolumeAnalyzer";
import { UserData } from "./Settings";
import { SessionStepContext } from "../utils/SessionStepContextProvider";

const theme = "好きな言葉";

const STUN_SERVERS = {
  iceServers: [
    { urls: "stun:stun.l.google.com:19302" },
    { urls: "stun:stun1.l.google.com:19302" },
  ],
};

type UserCardData = {
  name: string;
  icon: JSX.Element;
};

export const Session = () => {
  const sessionTime = 300; // [s]
  const [memoOpen, setMemoOpen] = useState(false);
  const [assistantOpen, setAssistantOpen] = useState(false);
  const [messages, setMessages] = useState<string[]>([]);
  const [inputMessage, setInputMessage] = useState<string>("");
  const [isLoading, setIsLoading] = useState(false); // ローディング状態を管理
  const [memo1, setMemo1] = useState(""); // メモ1を管理
  const [memo2, setMemo2] = useState(""); // メモ2を管理
  const [value, setValue] = useState("1");
  const [isMuted, setIsMuted] = useState(false);
  const [countdown, setCountdown] = useState(sessionTime + 3);
  const navigate = useNavigate();
  const { sessionStep, setSessionStep } = useContext(SessionStepContext);

  const [showTopicPopup, setShowTopicPopup] = useState(false);
  const [isPriorityHighClicked, setIsPriorityHighClicked] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => {
      if (!isPriorityHighClicked) {
        setShowTopicPopup(true);
      }
    }, 8000); // 2 minutes and 3 seconds
    return () => clearTimeout(timer);
  }, [isPriorityHighClicked]);

  useEffect(() => {
    if (countdown > 0) {
      const timer = setTimeout(() => setCountdown(countdown - 1), 1000);
      return () => clearTimeout(timer);
    } else {
      navigate("/sessioninterval");
      const nextStep = sessionStep + 1;
      setSessionStep(nextStep);
      if (nextStep >= 3) {
        navigate("/sessionrecord");
      }
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

  const handleChange = (_: React.SyntheticEvent, newValue: string) => {
    setValue(newValue);
  };
  const handleCloseTopicPopup = () => {
    setShowTopicPopup(false);
  };
  const handlePriorityHighClick = () => {
    setShowTopicPopup(true);
    setIsPriorityHighClicked(true);
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

      console.log("Response data:", response.data); // ここでレスポンスデータを確認

      // 応答メッセージを表示
      if (
        response.data &&
        response.data.choices &&
        response.data.choices[0].message.content
      ) {
        const assistantMessage = response.data.choices[0].message.content;
        setMessages((prevMessages) => [
          ...prevMessages,
          `Assistant: ${assistantMessage}`,
        ]);
      } else {
        setMessages((prevMessages) => [
          ...prevMessages,
          "Error: Invalid response format",
        ]);
      }
    } catch (error) {
      setMessages((prevMessages) => [
        ...prevMessages,
        "Error: Failed to get response",
      ]);
      console.log(error);
    } finally {
      setIsLoading(false); // ローディング状態を終了
    }
  };

  // WebRTC関連の処理
  // webbbb vimジャンプ用
  const host = "10.70.174.101";
  const WEBSOCKET_URL = "ws://" + host + ":8081/ws";
  let isOffer = false;

  const peerConnectionRef = useRef<RTCPeerConnection | null>(null);
  const websocketRef = useRef<WebSocket | null>(null);
  const remoteAudioRef = useRef<HTMLAudioElement>(null);
  const localStreamRef = useRef<MediaStream | null>(null);

  const cleanupResources = useCallback(() => {
    if (peerConnectionRef.current) {
      peerConnectionRef.current.close();
      peerConnectionRef.current = null;
    }
    if (localStreamRef.current) {
      localStreamRef.current.getTracks().forEach((track) => track.stop());
      localStreamRef.current = null;
    }
    if (remoteAudioRef.current) {
      remoteAudioRef.current.srcObject = null;
    }
  }, []);

  useEffect(() => {
    connectToSignalingServer(); // コンポーネントがマウントされたときにシグナリングサーバーに接続する
    setTimeout(() => startCall(false), 2000);
    return () => {
      if (websocketRef.current) {
        websocketRef.current.close();
      }
      cleanupResources();
    };
  }, [cleanupResources]);

  const connectToSignalingServer = (): void => {
    const ws = new WebSocket(WEBSOCKET_URL);

    ws.onopen = () => {
      console.log("Connected to signaling server");

      const token = localStorage.getItem("token"); // Assuming you store the token in localStorage
      const authMessage = {
        type: "Authorization",
        token: `Bearer ${token}`,
      };
      ws.send(JSON.stringify(authMessage));
    };

    ws.onmessage = async (event: MessageEvent) => {
      const message: {
        type: string;
        offer?: RTCSessionDescriptionInit;
        answer?: RTCSessionDescriptionInit;
        candidate?: RTCIceCandidateInit;
        isOffer?: boolean;
      } = JSON.parse(event.data);

      if (message.type == "callType") {
        if (message.isOffer != undefined) {
          isOffer = message.isOffer;
        }
      } else if (message.type == "offer" && !isOffer) {
        console.log("Received offer. start call after 1000ms.");
        setTimeout(() => startCall(true), 1000);
      }

      if (!peerConnectionRef.current) {
        console.warn("Received message but peer connection is not established");
        return;
      }
      try {
        if (message.type === "offer" && message.offer) {
          await handleOffer(message.offer);
        } else if (message.type === "answer" && message.answer) {
          await peerConnectionRef.current.setRemoteDescription(
            new RTCSessionDescription(message.answer)
          );
        } else if (message.type === "ice-candidate" && message.candidate) {
          await peerConnectionRef.current.addIceCandidate(
            new RTCIceCandidate(message.candidate)
          );
        } else if (message.type === "callType" && message.isOffer) {
          isOffer = message.isOffer;
        }
      } catch (error) {
        console.error("Error handling WebSocket message:", error);
      }
    };

    ws.onclose = () => {
      console.log("Disconnected from signaling server");
      cleanupResources();
    };

    websocketRef.current = ws;
  };

  const createPeerConnection = (): RTCPeerConnection => {
    const pc = new RTCPeerConnection(STUN_SERVERS);

    pc.onicecandidate = (event: RTCPeerConnectionIceEvent) => {
      if (event.candidate && websocketRef.current) {
        websocketRef.current.send(
          JSON.stringify({
            type: "ice-candidate",
            candidate: event.candidate,
          })
        );
      }
    };

    pc.ontrack = (event: RTCTrackEvent) => {
      if (remoteAudioRef.current && event.streams[0]) {
        remoteAudioRef.current.srcObject = event.streams[0];
        opponentVolumeAnalyzerRef.current = new AudioVolumeAnalyzer(
          setIsOpponentSpeak
        );
        opponentVolumeAnalyzerRef.current.start(event.streams[0]);
      }
    };

    pc.oniceconnectionstatechange = () => {
      console.log("ICE Connection State:", pc.iceConnectionState);
      if (
        pc.iceConnectionState === "disconnected" ||
        pc.iceConnectionState === "failed" ||
        pc.iceConnectionState === "closed"
      ) {
        cleanupResources();
      }
    };

    return pc;
  };

  let startCallCount = 0;
  const startCall = async (forceExcute: boolean): Promise<void> => {
    if (!forceExcute && !isOffer) {
      console.log("not start call", forceExcute, isOffer);
      console.log(
        `not start call forceExcute=${forceExcute}, isOffer=${isOffer}`
      );
      return;
    } else {
      console.log("startCall()");
    }
    if (startCallCount > 0) {
      return;
    }
    startCallCount++;
    if (forceExcute) {
      console.log("forceExcute");
    }
    console.log("start call");

    const pc = createPeerConnection();
    peerConnectionRef.current = pc;

    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      localStreamRef.current = stream;
      stream.getTracks().forEach((track) => pc.addTrack(track, stream));

      const offer = await pc.createOffer();
      await pc.setLocalDescription(offer);

      if (websocketRef.current) {
        websocketRef.current.send(
          JSON.stringify({
            type: "offer",
            offer: pc.localDescription,
          })
        );
      }

      volumeAnalyzerRef.current = new AudioVolumeAnalyzer(setisSpeak);
      volumeAnalyzerRef.current.start(stream);
    } catch (error) {
      console.error("Error starting call:", error);
      cleanupResources();
    }
  };

  const handleOffer = async (
    offer: RTCSessionDescriptionInit
  ): Promise<void> => {
    console.log("handle offer");
    const pc = createPeerConnection();
    peerConnectionRef.current = pc;

    try {
      const stream = await navigator.mediaDevices.getUserMedia({ audio: true });
      localStreamRef.current = stream;
      stream.getTracks().forEach((track) => pc.addTrack(track, stream));

      await pc.setRemoteDescription(new RTCSessionDescription(offer));

      const answer = await pc.createAnswer();
      await pc.setLocalDescription(answer);

      if (websocketRef.current) {
        websocketRef.current.send(
          JSON.stringify({
            type: "answer",
            answer: pc.localDescription,
          })
        );
      }
    } catch (error) {
      console.error("Error handling offer:", error);
      cleanupResources();
    }
  };

  const toggleMute = (): void => {
    setIsMuted((prev) => !prev);
    if (localStreamRef.current) {
      const audioTrack = localStreamRef.current.getAudioTracks()[0];
      if (audioTrack) {
        audioTrack.enabled = !audioTrack.enabled;
        setIsMuted(!audioTrack.enabled);
      }
    }
  };

  //const endCall = (): void => {
  //  cleanupResources();
  //};

  // visualize speaker
  const volumeAnalyzerRef = useRef<AudioVolumeAnalyzer | null>(null);
  const opponentVolumeAnalyzerRef = useRef<AudioVolumeAnalyzer | null>(null);
  const [isSpeak, setisSpeak] = useState(false);
  const [isOpponentSpeak, setIsOpponentSpeak] = useState(false);

  // userInfo
  const initialUserCardInfo = { name: "", icon: <Person /> };
  const [userCardInfo, setUserCardInfo] =
    useState<UserCardData>(initialUserCardInfo);
  const [opponentUserCardInfo, setOpponentUserCardInfo] =
    useState<UserCardData>(initialUserCardInfo);

  const getFullAvatarUrl = (avatarUrl: string) => {
    if (!avatarUrl) return ""; // デフォルトのアバター画像のURLを設定することもできます
    if (avatarUrl.startsWith("http")) return avatarUrl; // すでに完全なURLの場合
    return `http://localhost:8081${avatarUrl}`; // ローカル開発環境の場合
  };

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const userInfoResponse = await api.get<UserData>("/user/info");
        const opponentUserCardDataResponse =
          await api.get<UserData>("/user/info");
        const userInfo: UserCardData = {
          name: userInfoResponse.data.username,
          icon: (
            <Avatar
              src={getFullAvatarUrl(userInfoResponse.data.avatar_url)}
              sx={{ width: 80, height: 80 }}
            />
          ),
        };
        const opponentUserInfo: UserCardData = {
          name: opponentUserCardDataResponse.data.username,
          icon: <Person />,
        };
        setUserCardInfo(userInfo);
        setOpponentUserCardInfo(opponentUserInfo);
      } catch (error) {
        console.error("Failed to fetch user data:", error);
      }
    };

    fetchUserData();
  }, []);

  return (
    <SessionBottomNavigationTemplate
      value="other"
      isMute={isMuted}
      toggleMute={toggleMute}
      setMemoOpen={setMemoOpen}
      setAssistantOpen={setAssistantOpen}
      onPriorityHighClick={handlePriorityHighClick}
    >
      <SessionContainer
        theme={theme}
        users={[userCardInfo, opponentUserCardInfo]}
        isSpeak={isSpeak && !isMuted}
        isOpponentSpeak={isOpponentSpeak}
      />
      <TopicPopup isVisible={showTopicPopup} onClose={handleCloseTopicPopup} />
      <HalfModal open={memoOpen} handleClose={handleMemoClose} title="">
        <TabContext value={value}>
          <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
            <TabList
              onChange={handleChange}
              sx={{ display: "grid", placeContent: "center" }}
            >
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
      <HalfModal
        open={assistantOpen}
        handleClose={handleAssistantClose}
        title="アシスタント"
      >
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
              <ListItem
                key={index}
                sx={{
                  justifyContent: message.startsWith("You:")
                    ? "flex-end"
                    : "flex-start",
                }}
              >
                <Paper
                  sx={{
                    padding: "5px",
                    backgroundColor: message.startsWith("You:")
                      ? "#f0f0f0"
                      : "background.default",
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
          <Button
            variant="contained"
            color="primary"
            onClick={handleSendMessage}
            disabled={isLoading}
          >
            {isLoading ? "送信中..." : "送信"}
          </Button>
        </Box>
      </HalfModal>
      <audio ref={remoteAudioRef} autoPlay />
    </SessionBottomNavigationTemplate>
  );
};

export default Session;
